package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func (s serverImpl) getCreditType(ctx sdk.Context, creditTypeName string) (ecocredit.CreditType, error) {
	creditTypes := s.getAllCreditTypes(ctx)
	creditTypeName = ecocredit.NormalizeCreditTypeName(creditTypeName)
	for _, creditType := range creditTypes {
		// credit type name's stored via params have enforcement on normalization, so we can be sure they will already
		// be normalized here.
		if creditType.Name == creditTypeName {
			return *creditType, nil
		}
	}
	return ecocredit.CreditType{}, sdkerrors.ErrInvalidType.Wrapf("%s is not a valid credit type", creditTypeName)
}

func (s serverImpl) getAllCreditTypes(ctx sdk.Context) []*ecocredit.CreditType {
	var params ecocredit.Params
	s.paramSpace.GetParamSet(ctx, &params)
	return params.CreditTypes
}

// getCreditTypeSeqNextVal looks up the CreditTypeSeq for the given CreditType,
// returns its next value, and persists that value in the store. If there is no
// CreditTypeSeq for the given CreditType, then a new CreditTypeSeq is created
// starting at 1.
func (s serverImpl) getCreditTypeSeqNextVal(ctx sdk.Context, creditType ecocredit.CreditType) (uint64, error) {
	// Lookup the sequence number for the given credit type
	var creditTypeSeq ecocredit.CreditTypeSeq
	err := s.creditTypeSeqTable.GetOne(ctx, orm.RowID(creditType.Abbreviation), &creditTypeSeq)

	switch err {

	// There is an existing CreditTypeSeq, so increment it
	case nil:
		// Increment the sequence number
		creditTypeSeq.SeqNumber++
		err = s.creditTypeSeqTable.Update(ctx, &creditTypeSeq)
		if err != nil {
			return 0, err
		}

		// Return the new value
		return creditTypeSeq.SeqNumber, nil

	// There isn't a CreditTypeSeq for this CreditType, so create one
	case orm.ErrNotFound:
		// Create a new CreditTypeSeq starting at 1
		creditTypeSeq.Abbreviation = creditType.Abbreviation
		creditTypeSeq.SeqNumber = 1
		err = s.creditTypeSeqTable.Create(ctx, &creditTypeSeq)
		if err != nil {
			return 0, err
		}

		// Return the new value
		return creditTypeSeq.SeqNumber, nil

	// We got an unexpected err, so return it
	default:
		return 0, err
	}
}
