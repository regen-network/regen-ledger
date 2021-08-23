package ecocredit

import (
	"strings"
	"unicode"

	"github.com/regen-network/regen-ledger/orm"
)

var _, _, _ orm.PrimaryKeyed = &ClassInfo{}, &BatchInfo{}, &CreditTypeSeq{}

func (m *ClassInfo) PrimaryKey() []byte {
	return []byte(m.ClassId)
}

func (m *BatchInfo) PrimaryKey() []byte {
	return []byte(m.BatchDenom)
}

func (m *CreditTypeSeq) PrimaryKey() []byte {
	return []byte(m.Abbreviation)
}

// Normalize credit type name by removing whitespace and converting to lowercase
func NormalizeCreditTypeName(name string) string {
	return FastRemoveWhitespace(strings.ToLower(name))
}

func FastRemoveWhitespace(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
