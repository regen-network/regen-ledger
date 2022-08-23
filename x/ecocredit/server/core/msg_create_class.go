package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// CreateClass creates a new class of ecocredit
//
// The admin is charged a fee for creating the class. This is controlled by
// the global parameter CreditClassFee, which can be updated through the
// governance process.
func (k Keeper) CreateClass(goCtx context.Context, req *core.MsgCreateClass) (*core.MsgCreateClassResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(goCtx)
	adminAddress, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}

	if err := k.assertCanCreateClass(goCtx, adminAddress); err != nil {
		return nil, err
	}

	var fees *api.ClassFees
	fees, err = k.stateStore.ClassFeesTable().Get(goCtx)
	if err != nil {
		if !ormerrors.NotFound.Is(err) {
			return nil, err
		}
		fees = &api.ClassFees{}
	}

	allowedFees, ok := types.ProtoCoinsToCoins(fees.Fees)
	if !ok {
		return nil, sdkerrors.ErrInvalidType.Wrap("credit class fee")
	}

	// only check and charge fee if allowed fees is not empty
	if allowedFees.Len() > 0 {

		// check if fee is empty
		if req.Fee == nil {
			if len(allowedFees) > 1 {
				return nil, sdkerrors.ErrInsufficientFee.Wrapf(
					"fee cannot be empty: must be one of %s", allowedFees,
				)
			}
			return nil, sdkerrors.ErrInsufficientFee.Wrapf(
				"fee cannot be empty: must be %s", allowedFees,
			)
		}

		// convert fee to multiple coins for verification
		coins := sdk.Coins{*req.Fee}

		// check if fee is greater than or equal to any coin in allowedFees
		if !coins.IsAnyGTE(allowedFees) {
			if len(allowedFees) > 1 {
				return nil, sdkerrors.ErrInsufficientFee.Wrapf(
					"fee must be one of %s, got %s", allowedFees, req.Fee,
				)
			}
			return nil, sdkerrors.ErrInsufficientFee.Wrapf(
				"fee must be %s, got %s", allowedFees, req.Fee,
			)
		}

		// only check and charge the minimum fee amount
		minimumFee := sdk.Coin{
			Denom:  req.Fee.Denom,
			Amount: allowedFees.AmountOf(req.Fee.Denom),
		}

		// check admin balance against minimum fee
		adminBalance := k.bankKeeper.GetBalance(sdkCtx, adminAddress, minimumFee.Denom)
		if adminBalance.IsNil() || adminBalance.IsLT(minimumFee) {
			return nil, sdkerrors.ErrInsufficientFunds.Wrapf(
				"insufficient balance %s for bank denom %s", adminBalance.Amount, minimumFee.Denom,
			)
		}

		// send coins from account to module and then burn the coins
		err = k.chargeCreditClassFee(sdkCtx, adminAddress, sdk.Coins{minimumFee})
		if err != nil {
			return nil, err
		}
	}

	creditType, err := k.stateStore.CreditTypeTable().Get(goCtx, req.CreditTypeAbbrev)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get credit type with abbreviation %s: %s", req.CreditTypeAbbrev, err.Error())
	}

	// default the sequence to 1 for the `not found` case.
	// will get overwritten by the actual sequence if it exists.
	var seq uint64 = 1
	classSeq, err := k.stateStore.ClassSequenceTable().Get(goCtx, creditType.Abbreviation)
	if err != nil {
		if !ormerrors.IsNotFound(err) {
			return nil, err
		}
	} else {
		seq = classSeq.NextSequence
	}
	if err = k.stateStore.ClassSequenceTable().Save(goCtx, &api.ClassSequence{
		CreditTypeAbbrev: creditType.Abbreviation,
		NextSequence:     seq + 1,
	}); err != nil {
		return nil, err
	}

	classID := core.FormatClassID(creditType.Abbreviation, seq)

	key, err := k.stateStore.ClassTable().InsertReturningID(goCtx, &api.Class{
		Id:               classID,
		Admin:            adminAddress,
		Metadata:         req.Metadata,
		CreditTypeAbbrev: creditType.Abbreviation,
	})
	if err != nil {
		return nil, err
	}

	for _, issuer := range req.Issuers {
		issuer, err := sdk.AccAddressFromBech32(issuer)
		if err != nil {
			return nil, err
		}
		if err = k.stateStore.ClassIssuerTable().Insert(goCtx, &api.ClassIssuer{
			ClassKey: key,
			Issuer:   issuer,
		}); err != nil {
			return nil, err
		}

		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/core/MsgCreateClass issuer iteration")
	}

	err = sdkCtx.EventManager().EmitTypedEvent(&core.EventCreateClass{
		ClassId: classID,
	})
	if err != nil {
		return nil, err
	}

	return &core.MsgCreateClassResponse{ClassId: classID}, nil
}

func (k Keeper) assertCanCreateClass(ctx context.Context, adminAddress sdk.AccAddress) error {
	allowListEnabled, err := k.stateStore.AllowListEnabledTable().Get(ctx)
	if err != nil {
		return err
	}

	if allowListEnabled.Enabled {
		_, err := k.stateStore.AllowedClassCreatorTable().Get(ctx, adminAddress)
		if err != nil {
			if ormerrors.NotFound.Is(err) {
				return sdkerrors.ErrUnauthorized.Wrapf("%s is not allowed to create credit classes", adminAddress.String())
			}
			return err
		}
	}
	return nil
}

func (k Keeper) chargeCreditClassFee(ctx sdk.Context, creatorAddr sdk.AccAddress, creditClassFee sdk.Coins) error {
	// Move the fee to the ecocredit module's account
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddr, ecocredit.ModuleName, creditClassFee)
	if err != nil {
		return err
	}

	err = k.bankKeeper.BurnCoins(ctx, ecocredit.ModuleName, creditClassFee)
	if err != nil {
		return err
	}

	return nil
}
