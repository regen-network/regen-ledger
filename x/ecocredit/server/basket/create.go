package basket

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

// Create is an RPC to handle basket.MsgCreate
func (k Keeper) Create(ctx context.Context, msg *basket.MsgCreate) (*basket.MsgCreateResponse, error) {
	sdkCtx := types.UnwrapSDKContext(ctx)
	fee := k.ecocreditKeeper.GetCreateBasketFee(ctx)
	if err := basket.ValidateMsgCreate(msg, fee); err != nil {
		return nil, err
	}
	sender, err := sdk.AccAddressFromBech32(msg.Curator)
	if err != nil {
		return nil, err
	}

	err = k.distKeeper.FundCommunityPool(sdkCtx.Context, fee, sender)
	if err != nil {
		return nil, err
	}
	if err = validateCreditType(ctx, k.ecocreditKeeper, msg.CreditTypeAbbrev, msg.Exponent); err != nil {
		return nil, err
	}
	denom, displayDenom, err := basket.BasketDenom(msg.Name, msg.CreditTypeAbbrev, msg.Exponent)
	if err != nil {
		return nil, err
	}

	id, err := k.stateStore.BasketTable().InsertReturningID(ctx, &api.Basket{
		Curator:           msg.Curator,
		BasketDenom:       denom,
		DisableAutoRetire: msg.DisableAutoRetire,
		CreditTypeAbbrev:  msg.CreditTypeAbbrev,
		DateCriteria:      msg.DateCriteria.ToApi(),
		Exponent:          msg.Exponent,
		Name:              msg.Name,
	})
	if err != nil {
		return nil, err
	}
	if err = k.indexAllowedClasses(sdkCtx, id, msg.AllowedClasses); err != nil {
		return nil, err
	}

	denomUnits := []*banktypes.DenomUnit{{
		Denom:    displayDenom,
		Exponent: msg.Exponent,
		Aliases:  nil,
	}}
	if msg.Exponent != 0 {
		denomUnits = append(denomUnits, &banktypes.DenomUnit{
			Denom:    denom,
			Exponent: 0, // conversion from base denom to this denom
			Aliases:  nil,
		})
	}

	k.bankKeeper.SetDenomMetaData(sdkCtx.Context, banktypes.Metadata{
		DenomUnits:  denomUnits,
		Description: msg.Description,
		Base:        denom,
		Display:     displayDenom,
		Name:        msg.Name,
		Symbol:      msg.Name,
	})

	err = sdkCtx.Context.EventManager().EmitTypedEvent(&basket.EventCreate{
		BasketDenom: denom,
		Curator:     msg.Curator,
	})

	return &basket.MsgCreateResponse{BasketDenom: denom}, err
}

// validateCreditType returns error if a given credit type abbreviation doesn't exist or
// it's precision is bigger then the requested exponent.
func validateCreditType(sdkCtx context.Context, k EcocreditKeeper, creditTypeAbbr string, exponent uint32) error {
	res, err := k.CreditTypes(sdkCtx, &ecocredit.QueryCreditTypesRequest{})
	if err != nil {
		return err
	}

	for _, c := range res.CreditTypes {
		if c.Abbreviation == creditTypeAbbr {
			if c.Precision > exponent {
				return sdkerrors.ErrInvalidRequest.Wrapf(
					"exponent %d must be >= credit type precision %d",
					exponent,
					c.Precision,
				)
			}
			return nil
		}
	}
	return sdkerrors.ErrInvalidRequest.Wrapf("credit type abbreviation %q doesn't exist", creditTypeAbbr)
}

func (k Keeper) indexAllowedClasses(sdkCtx types.Context, basketID uint64, allowedClasses []string) error {
	for _, class := range allowedClasses {
		if !k.ecocreditKeeper.HasClassInfo(sdkCtx, class) {
			return sdkerrors.ErrInvalidRequest.Wrapf("credit class %q doesn't exist", class)
		}

		ctx := sdk.WrapSDKContext(sdkCtx.Context)
		err := k.stateStore.BasketClassTable().Insert(ctx,
			&api.BasketClass{
				BasketId: basketID,
				ClassId:  class,
			},
		)
		if err != nil {
			return err
		}

		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/basket/MsgCreate class iteration")
	}
	return nil
}
