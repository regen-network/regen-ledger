package basket

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

// Create is an RPC to handle basket.MsgCreate
func (k Keeper) Create(ctx context.Context, msg *basket.MsgCreate) (*basket.MsgCreateResponse, error) {
	rgCtx := types.UnwrapSDKContext(ctx)
	fee := k.ecocreditKeeper.GetCreateBasketFee(ctx)
	if err := basket.ValidateMsgCreate(msg, fee); err != nil {
		return nil, err

	}
	sender, err := sdk.AccAddressFromBech32(msg.Curator)
	if err != nil {
		return nil, err
	}

	err = k.distKeeper.FundCommunityPool(rgCtx.Context, fee, sender)
	if err != nil {
		return nil, err
	}
	if err = validateCreditType(ctx, k.ecocreditKeeper, msg.CreditTypeAbbrev, msg.Exponent); err != nil {
		return nil, err
	}

	denom := "eco/" + msg.Prefix + msg.Name
	display := "eco/" + msg.Name

	id, err := k.stateStore.BasketStore().InsertReturningID(ctx, &basketv1.Basket{
		BasketDenom:       denom,
		DisableAutoRetire: msg.DisableAutoRetire,
		// TODO: need to release new api
		CreditTypeName: msg.CreditTypeAbbrev,
		DateCriteria:   msg.DateCriteria.ToApi(),
		Exponent:       msg.Exponent,
	})
	if err != nil {
		return nil, err
	}
	if err = k.indexAllowedClasses(rgCtx, id, msg.AllowedClasses); err != nil {
		return nil, err
	}

	k.bankKeeper.SetDenomMetaData(rgCtx.Context, banktypes.Metadata{
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    denom,
				Exponent: msg.Exponent,
				Aliases:  nil,
			},
		},
		Description: msg.Description,
		Base:        denom,
		Display:     display,
		Name:        msg.Name,
		Symbol:      msg.Name,
	})

	err = rgCtx.Context.EventManager().EmitTypedEvent(&basket.EventCreate{
		BasketDenom: denom,
		Curator:     msg.Curator,
	})

	return &basket.MsgCreateResponse{BasketDenom: denom}, err
}

// validateCreditType returns error if a given credit type abbreviation doesn't exist or
// it's precision is bigger then the requested exponent.
func validateCreditType(ctx context.Context, k EcocreditKeeper, creditTypeAbbr string, exponent uint32) error {
	res, err := k.CreditTypes(ctx, &ecocredit.QueryCreditTypesRequest{})
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

func (k Keeper) indexAllowedClasses(ctx types.Context, basketID uint64, allowedClasses []string) error {
	for _, class := range allowedClasses {
		if !k.ecocreditKeeper.HasClassInfo(ctx, class) {
			return sdkerrors.ErrInvalidRequest.Wrapf("credit class %q doesn't exist", class)
		}

		err := k.stateStore.BasketClassStore().Insert(ctx,
			&basketv1.BasketClass{
				BasketId: basketID,
				ClassId:  class,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}
