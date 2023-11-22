package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	regentypes "github.com/regen-network/regen-ledger/types/v2"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/basket"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/basket/types/v1"
)

// Create is an RPC to handle basket.MsgCreate
func (k Keeper) Create(ctx context.Context, msg *types.MsgCreate) (*types.MsgCreateResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	basketFee, err := k.stateStore.BasketFeeTable().Get(ctx)
	if err != nil {
		return nil, err
	}

	curator, err := sdk.AccAddressFromBech32(msg.Curator)
	if err != nil {
		return nil, err
	}

	// only check and charge fee if required fee is set
	if basketFee.Fee != nil {

		requiredFee, ok := regentypes.ProtoCoinToCoin(basketFee.Fee)
		if !ok {
			return nil, sdkerrors.ErrInvalidType.Wrapf("basket fee")
		}

		// check if fee is empty
		if msg.Fee == nil {
			return nil, sdkerrors.ErrInsufficientFee.Wrapf(
				"fee cannot be empty: must be %s", requiredFee,
			)
		}

		// In the next version of the basket package, the fee provided will be
		// updated to a single Coin rather than a list of Coins. In the meantime,
		// the message will fail basic validation if more than one Coin is provided.
		msgFee := msg.Fee[0]

		// check if fee is the correct denom
		if msgFee.Denom != requiredFee.Denom {
			return nil, sdkerrors.ErrInsufficientFee.Wrapf(
				"fee must be %s, got %s", requiredFee, msgFee,
			)
		}

		// check if fee is greater than or equal to required fee
		if !msgFee.IsGTE(requiredFee) {
			return nil, sdkerrors.ErrInsufficientFee.Wrapf(
				"fee must be %s, got %s", requiredFee, msgFee,
			)
		}

		// check curator balance against required fee
		curatorBalance := k.bankKeeper.GetBalance(sdkCtx, curator, requiredFee.Denom)
		if curatorBalance.IsNil() || curatorBalance.IsLT(requiredFee) {
			return nil, sdkerrors.ErrInsufficientFunds.Wrapf(
				"insufficient balance %s for bank denom %s", curatorBalance.Amount, requiredFee.Denom,
			)
		}

		// convert required fee to multiple coins for processing
		requiredFees := sdk.Coins{requiredFee}

		err = k.bankKeeper.SendCoinsFromAccountToModule(sdkCtx, curator, basket.BasketSubModuleName, requiredFees)
		if err != nil {
			return nil, err
		}

		err = k.bankKeeper.BurnCoins(sdkCtx, basket.BasketSubModuleName, requiredFees)
		if err != nil {
			return nil, err
		}
	}

	creditType, err := k.baseStore.CreditTypeTable().Get(ctx, msg.CreditTypeAbbrev)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf(
			"could not get credit type with abbreviation %s: %s", msg.CreditTypeAbbrev, err,
		)
	}

	denom, displayDenom, err := basket.FormatBasketDenom(msg.Name, msg.CreditTypeAbbrev, creditType.Precision)
	if err != nil {
		return nil, err
	}

	id, err := k.stateStore.BasketTable().InsertReturningID(ctx, &api.Basket{
		Curator:           curator,
		BasketDenom:       denom,
		DisableAutoRetire: msg.DisableAutoRetire,
		CreditTypeAbbrev:  msg.CreditTypeAbbrev,
		DateCriteria:      msg.DateCriteria.ToAPI(),
		Exponent:          creditType.Precision, // exponent is no longer used but set until removed
		Name:              msg.Name,
	})
	if err != nil {
		return nil, ormerrors.UniqueKeyViolation.Wrapf("basket with name %s already exists", msg.Name)
	}
	if err = k.indexAllowedClasses(ctx, id, msg.AllowedClasses, msg.CreditTypeAbbrev); err != nil {
		return nil, err
	}

	denomUnits := make([]*banktypes.DenomUnit, 0)

	// Set denomination units in ascending order and
	// the first denomination unit must be the base
	if creditType.Precision == 0 {
		denomUnits = append(denomUnits, &banktypes.DenomUnit{
			Denom:    displayDenom,
			Exponent: creditType.Precision,
		})
	} else {
		denomUnits = append(denomUnits, &banktypes.DenomUnit{
			Denom: denom,
		})
		denomUnits = append(denomUnits, &banktypes.DenomUnit{
			Denom:    displayDenom,
			Exponent: creditType.Precision,
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

	err = sdkCtx.EventManager().EmitTypedEvent(&types.EventCreate{
		BasketDenom: denom,
		Curator:     msg.Curator,
	})

	return &types.MsgCreateResponse{BasketDenom: denom}, err
}

// indexAllowedClasses checks that all `allowedClasses` both exist, and are of the specified credit type, then inserts
// the class into the BasketClass table.
func (k Keeper) indexAllowedClasses(ctx context.Context, basketID uint64, allowedClasses []string, creditTypeAbbrev string) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	for _, class := range allowedClasses {
		classInfo, err := k.baseStore.ClassTable().GetById(ctx, class)
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
