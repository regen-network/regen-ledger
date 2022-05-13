package basket

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Create is an RPC to handle basket.MsgCreate
func (k Keeper) Create(ctx context.Context, msg *basket.MsgCreate) (*basket.MsgCreateResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	var fee sdk.Coins
	k.paramsKeeper.Get(sdkCtx, core.KeyBasketCreationFee, &fee)
	if err := basket.ValidateMsgCreate(msg, fee); err != nil {
		return nil, err
	}

	curator, err := sdk.AccAddressFromBech32(msg.Curator)
	if err != nil {
		return nil, err
	}

	err = k.distKeeper.FundCommunityPool(sdkCtx, fee, curator)
	if err != nil {
		return nil, err
	}
	if err = k.validateCreditType(ctx, msg.CreditTypeAbbrev, msg.Exponent); err != nil {
		return nil, err
	}
	denom, displayDenom, err := basket.BasketDenom(msg.Name, msg.CreditTypeAbbrev, msg.Exponent)
	if err != nil {
		return nil, err
	}

	id, err := k.stateStore.BasketTable().InsertReturningID(ctx, &api.Basket{
		Curator:           curator,
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
	if err = k.indexAllowedClasses(ctx, id, msg.AllowedClasses, msg.CreditTypeAbbrev); err != nil {
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

	k.bankKeeper.SetDenomMetaData(sdkCtx, banktypes.Metadata{
		DenomUnits:  denomUnits,
		Description: msg.Description,
		Base:        denom,
		Display:     displayDenom,
		Name:        msg.Name,
		Symbol:      msg.Name,
	})

	err = sdkCtx.EventManager().EmitTypedEvent(&basket.EventCreate{
		BasketDenom: denom,
		Curator:     msg.Curator,
	})

	return &basket.MsgCreateResponse{BasketDenom: denom}, err
}

// validateCreditType returns error if a given credit type abbreviation doesn't exist or
// it's precision is bigger then the requested exponent.
func (k Keeper) validateCreditType(ctx context.Context, abbreviation string, exponent uint32) error {
	creditType, err := k.coreStore.CreditTypeTable().Get(ctx, abbreviation)
	if err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("could not get credit type with abbreviation %s: %s", abbreviation, err.Error())
	}
	if creditType.Precision > exponent {
		return sdkerrors.ErrInvalidRequest.Wrapf(
			"exponent %d must be >= credit type precision %d",
			exponent,
			creditType.Precision,
		)
	}
	return nil
}

// indexAllowedClasses checks that all `allowedClasses` both exist, and are of the specified credit type, then inserts
// the class into the BasketClass table.
func (k Keeper) indexAllowedClasses(ctx context.Context, basketID uint64, allowedClasses []string, creditTypeAbbrev string) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	for _, class := range allowedClasses {
		classInfo, err := k.coreStore.ClassTable().GetById(ctx, class)
		if err != nil {
			return sdkerrors.ErrInvalidRequest.Wrapf("could not get credit class %s: %s", class, err.Error())
		}

		if classInfo.CreditTypeAbbrev != creditTypeAbbrev {
			return sdkerrors.ErrInvalidRequest.Wrapf("basket specified credit type %s, but class %s is of type %s",
				creditTypeAbbrev, class, classInfo.CreditTypeAbbrev)
		}

		if err := k.stateStore.BasketClassTable().Insert(ctx,
			&api.BasketClass{
				BasketId: basketID,
				ClassId:  class,
			},
		); err != nil {
			return err
		}

		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/basket/MsgCreate class iteration")
	}
	return nil
}
