package ecocredit

import (
	"strings"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/ecocredit/util"
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

// Normalize credit type name by removing whitespace and converting to lowercase
func NormalizeCreditTypeName(name string) string {
	return util.FastRemoveWhitespace(strings.ToLower(name))
}
