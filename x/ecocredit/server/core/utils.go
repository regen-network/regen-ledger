package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// assertClassIssuer makes sure that the issuer is part of issuers of given classID.
// Returns ErrUnauthorized otherwise.
func (k Keeper) assertClassIssuer(goCtx context.Context, classID uint64, addr sdk.AccAddress) error {
	found, err := k.stateStore.ClassIssuerTable().Has(goCtx, classID, addr)
	if err != nil {
		return err
	}
	if !found {
		return sdkerrors.ErrUnauthorized.Wrapf("%s is not an issuer for the class", addr)
	}
	return nil
}
