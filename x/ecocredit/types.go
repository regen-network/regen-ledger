package ecocredit

import "github.com/cosmos/modules/incubator/orm"

const ModuleName = "ecocredit"

var _, _ orm.NaturalKeyed = &ClassInfo{}, &BatchInfo{}

func (m *ClassInfo) NaturalKey() []byte {
	return []byte(m.ClassId)
}

func (m *BatchInfo) NaturalKey() []byte {
	return []byte(m.BatchDenom)
}
