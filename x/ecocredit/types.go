package ecocredit

import (
	"github.com/regen-network/regen-ledger/orm"
)

var _, _ orm.PrimaryKeyed = &ClassInfo{}, &BatchInfo{}

func (m *ClassInfo) PrimaryKeyFields() []interface{} {
	return []interface{}{m.ClassId}
}

func (m *BatchInfo) PrimaryKeyFields() []interface{} {
	return []interface{}{m.BatchDenom}
}

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}
