package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
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

	// TODO: remove params https://github.com/regen-network/regen-ledger/issues/729
	var params ecocredit.Params
	k.params.GetParamSet(sdkCtx, &params)
	if params.AllowlistEnabled && !k.isCreatorAllowListed(sdkCtx, params.AllowedClassCreators, adminAddress) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("%s is not allowed to create credit classes", adminAddress.String())
	}

	// Charge the admin a fee to create the credit class
	err = k.chargeCreditClassFee(sdkCtx, adminAddress, params.CreditClassFee)
	if err != nil {
		return nil, err
	}

	creditType, err := server.GetCreditType(req.CreditTypeAbbrev, params.CreditTypes)
	if err != nil {
		return nil, err
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
		seq = classSeq.NextClassId
	}
	if err = k.stateStore.ClassSequenceTable().Save(goCtx, &api.ClassSequence{
		CreditType:  creditType.Abbreviation,
		NextClassId: seq + 1,
	}); err != nil {
		return nil, err
	}

	classID := ecocredit.FormatClassID(creditType.Abbreviation, seq)

	rowId, err := k.stateStore.ClassInfoTable().InsertReturningID(goCtx, &api.ClassInfo{
		Name:       classID,
		Admin:      adminAddress,
		Metadata:   req.Metadata,
		CreditType: creditType.Abbreviation,
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
			ClassId: rowId,
			Issuer:  issuer,
		}); err != nil {
			return nil, err
		}
	}

	err = sdkCtx.EventManager().EmitTypedEvent(&core.EventCreateClass{
		ClassId: classID,
		Admin:   req.Admin,
	})
	if err != nil {
		return nil, err
	}

	return &core.MsgCreateClassResponse{ClassId: classID}, nil
}

func (k Keeper) isCreatorAllowListed(ctx sdk.Context, allowlist []string, designer sdk.AccAddress) bool {
	for _, addr := range allowlist {
		ctx.GasMeter().ConsumeGas(gasCostPerIteration, "credit class creators allowlist")
		allowListedAddr, _ := sdk.AccAddressFromBech32(addr)
		if designer.Equals(allowListedAddr) {
			return true
		}
	}
	return false
}

func (k Keeper) chargeCreditClassFee(ctx sdk.Context, creatorAddr sdk.AccAddress, creditClassFee sdk.Coins) error {
	// Move the fee to the ecocredit module's account
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddr, ecocredit.ModuleName, creditClassFee)
	if err != nil {
		return err
	}

	// Burn the coins
	// TODO: Update this implementation based on the discussion at
	// https://github.com/regen-network/regen-ledger/issues/351
	err = k.bankKeeper.BurnCoins(ctx, ecocredit.ModuleName, creditClassFee)
	if err != nil {
		return err
	}

	return nil
}
