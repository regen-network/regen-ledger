package ecocredit

import (
	"strings"
	"unicode"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/protobuf/reflect/protoreflect"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/orm"
)

var _, _, _ orm.PrimaryKeyed = &ClassInfo{}, &BatchInfo{}, &CreditTypeSeq{}

var ModuleSchema = ormdb.ModuleSchema{
	FileDescriptors: map[uint32]protoreflect.FileDescriptor{
		1: basketv1.File_regen_ecocredit_basket_v1_state_proto,
	},
	Prefix: []byte{ORMPrefix},
}

// PrimaryKeyFields returns the fields of the object that will make up the
// primary key for ClassInfo.
func (m *ClassInfo) PrimaryKeyFields() []interface{} {
	return []interface{}{m.ClassId}
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
