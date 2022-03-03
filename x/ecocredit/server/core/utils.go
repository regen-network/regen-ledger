package core

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// getCreditType searches for a credit type that matches the given abbreviation within a credit type slice.
func (k Keeper) getCreditType(ctAbbrev string, creditTypes []*ecocredit.CreditType) (ecocredit.CreditType, error) {
	//creditTypeName = ecocredit.NormalizeCreditTypeName(creditTypeName)
	for _, creditType := range creditTypes {
		// credit type name's stored via params have enforcement on normalization, so we can be sure they will already
		// be normalized here.
		if creditType.Abbreviation == ctAbbrev {
			return *creditType, nil
		}
	}
	return ecocredit.CreditType{}, sdkerrors.ErrInvalidType.Wrapf("%s is not a valid credit type", ctAbbrev)
}
