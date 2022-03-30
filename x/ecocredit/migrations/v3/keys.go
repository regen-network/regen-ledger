package v3

import orm "github.com/regen-network/regen-ledger/orm"

const (
	TradableBalancePrefix byte = 0x0
	TradableSupplyPrefix  byte = 0x1
	RetiredBalancePrefix  byte = 0x2
	RetiredSupplyPrefix   byte = 0x3
	ClassInfoTablePrefix  byte = 0x5
	BatchInfoTablePrefix  byte = 0x6
)

var _, _, _, _, _, _, _ orm.PrimaryKeyed = &ClassInfo{}, &BatchInfo{}

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
