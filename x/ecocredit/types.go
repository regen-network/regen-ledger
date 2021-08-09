package ecocredit

import (
	"strings"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/ecocredit/util"
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

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

// Normalize credit type name by removing whitespace and converting to lowercase
func NormalizeCreditTypeName(name string) string {
	return util.FastRemoveWhitespace(strings.ToLower(name))
}
