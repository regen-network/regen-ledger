package bond

import "github.com/regen-network/regen-ledger/orm"

var _ orm.PrimaryKeyed = &BondInfo{}

// PrimaryKeyFields returns the fields of the object that will make up the
// primary key for ClassInfo.
func (m *BondInfo) PrimaryKeyFields() []interface{} {
	return []interface{}{m.Id}
}
