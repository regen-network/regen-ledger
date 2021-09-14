package ecocredit

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
	"unicode"

	"github.com/regen-network/regen-ledger/orm"
)

var _, _, _ orm.PrimaryKeyed = &ClassInfo{}, &BatchInfo{}, &CreditTypeSeq{}

func (m *ClassInfo) PrimaryKeyFields() []interface{} {
	return []interface{}{m.ClassId}
}

func (m *BatchInfo) PrimaryKeyFields() []interface{} {
	return []interface{}{m.BatchDenom}
}

func (m *CreditTypeSeq) PrimaryKeyFields() []interface{} {
	return []interface{}{m.Abbreviation}
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

// Normalize credit type name by removing whitespace and converting to lowercase
func NormalizeCreditTypeName(name string) string {
	return fastRemoveWhitespace(strings.ToLower(name))
}

func fastRemoveWhitespace(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
