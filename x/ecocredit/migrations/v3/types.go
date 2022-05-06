package v3

import (
	"strings"
	"unicode"

	orm "github.com/regen-network/regen-ledger/orm"
)

type BatchDenomT string

var _, _, _ orm.PrimaryKeyed = &ClassInfo{}, &BatchInfo{}, &CreditTypeSeq{}

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

// NormalizeCreditTypeName credit type name by removing whitespace and converting to lowercase.
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
