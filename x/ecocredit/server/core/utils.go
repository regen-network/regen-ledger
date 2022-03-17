package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// assertClassIssuer makes sure that the issuer is part of issuers of given classID.
// Returns ErrUnauthorized otherwise.
func (k Keeper) assertClassIssuer(goCtx context.Context, classID uint64, issuer string) error {
	addr, err := sdk.AccAddressFromBech32(issuer)
	if err != nil {
		return err
	}
	found, err := k.stateStore.ClassIssuerTable().Has(goCtx, classID, addr)
	if err != nil {
		return err
	}
	if !found {
		return sdkerrors.ErrUnauthorized.Wrapf("%s is not an issuer for the class", issuer)
	}
	return nil
}

// GetCreditTypeFromBatchDenom extracts the classId from a batch denom string, then retrieves it from the params.
func GetCreditTypeFromBatchDenom(ctx context.Context, store ecocreditv1.StateStore, k ParamKeeper, denom string) (core.CreditType, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	classId := core.GetClassIdFromBatchDenom(denom)
	classInfo, err := store.ClassInfoTable().GetByName(ctx, classId)
	if err != nil {
		return core.CreditType{}, err
	}
	p := &core.Params{}
	k.GetParamSet(sdkCtx, p)
	return getCreditType(classInfo.CreditType, p.CreditTypes)
}
