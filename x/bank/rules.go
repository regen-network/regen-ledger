package bank

import (
	"github.com/cockroachdb/apd/v2"
)

type MintRule interface {
	CanMint(minter, recipient string, amount *apd.Decimal) bool
}

type SendRule interface {
	CanSend(from, to string, amount *apd.Decimal) bool
}

type MoveRule interface {
	CanMove(mover, from, to string, amount *apd.Decimal) bool
}

type BurnRule interface {
	CanBurn(burner string, amount *apd.Decimal) bool
}

var _ MintRule = &ACLRule{}
var _ MoveRule = &ACLRule{}
var _ BurnRule = &ACLRule{}
var _ SendRule = &BooleanRule{}

func (m *ACLRule) isAllowedAddress(addr string) bool {
	for _, allowed := range m.AllowedAddresses {
		if allowed == addr {
			return true
		}
	}
	return false
}

func (m *ACLRule) CanMint(minter, _ string, _ *apd.Decimal) bool {
	return m.isAllowedAddress(minter)
}

func (m *ACLRule) CanMove(mover, _, _ string, _ *apd.Decimal) bool {
	return m.isAllowedAddress(mover)
}

func (m *ACLRule) CanBurn(burner string, _ *apd.Decimal) bool {
	return m.isAllowedAddress(burner)
}

func (m *BooleanRule) CanSend(string, string, *apd.Decimal) bool {
	return m.Enabled
}
