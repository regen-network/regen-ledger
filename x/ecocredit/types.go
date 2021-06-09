package ecocredit

import (
	"github.com/regen-network/regen-ledger/orm"
)

var _, _ orm.PrimaryKeyed = &ClassInfo{}, &BatchInfo{}

func (m *ClassInfo) PrimaryKey() []byte {
	return []byte(m.ClassId)
}

func (m *BatchInfo) PrimaryKey() []byte {
	return []byte(m.BatchDenom)
}

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}
