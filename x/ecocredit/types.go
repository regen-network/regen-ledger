package ecocredit

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/orm"
)

var _, _, _, _, _, _, _ orm.PrimaryKeyed = &ClassInfo{}, &BatchInfo{}, &CreditTypeSeq{}, &SellOrder{}, &BuyOrder{}, &AskDenom{}, &ProjectInfo{}

// PrimaryKeyFields returns the fields of the object that will make up the
// primary key for ClassInfo.
func (m *ClassInfo) PrimaryKeyFields() []interface{} {
	return []interface{}{m.ClassId}
}

// PrimaryKeyFields returns the fields of the object that will make up the
// primary key for ProjectInfo.
func (m *ProjectInfo) PrimaryKeyFields() []interface{} {
	return []interface{}{m.ProjectId}
}

// PrimaryKeyFields returns the fields of the object that will make up the
// primary key for BatchInfo.
func (m *BatchInfo) PrimaryKeyFields() []interface{} {
	return []interface{}{m.BatchDenom}
}

// PrimaryKeyFields returns the fields of the object that will make up the
// primary key for CreditTypeSeq.
func (m *CreditTypeSeq) PrimaryKeyFields() []interface{} {
	return []interface{}{m.Abbreviation}
}

// PrimaryKeyFields returns the fields of the object that will make up the
// primary key for SellOrder.
func (m *SellOrder) PrimaryKeyFields() []interface{} {
	return []interface{}{m.OrderId}
}

// PrimaryKeyFields returns the fields of the object that will make up the
// primary key for BuyOrder.
func (m *BuyOrder) PrimaryKeyFields() []interface{} {
	return []interface{}{m.BuyOrderId}
}

// PrimaryKeyFields returns the fields of the object that will make up the
// primary key for AskDenom.
func (m *AskDenom) PrimaryKeyFields() []interface{} {
	return []interface{}{m.Denom}
}

// AssertClassIssuer makes sure that the issuer is part of issuers of given classID.
// Returns ErrUnauthorized otherwise.
func (m *ClassInfo) AssertClassIssuer(issuer string) error {
	for _, i := range m.Issuers {
		if issuer == i {
			return nil
		}
	}
	return sdkerrors.ErrUnauthorized
}

// AssertProjectIssuer makes sure that the issuer is equals to the issuer of the credit batches for this project.
// Returns ErrUnauthorized otherwise.
func (m *ProjectInfo) AssertProjectIssuer(issuer string) error {
	if m.Issuer == issuer {
		return nil
	}

	return sdkerrors.ErrUnauthorized
}
