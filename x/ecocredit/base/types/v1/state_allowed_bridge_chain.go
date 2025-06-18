package v1

import (
	"github.com/regen-network/regen-ledger/x/ecocredit/v4"
)

// Validate performs basic validation of the AllowedBridgeChain state type
func (m *AllowedBridgeChain) Validate() error {
	if m.ChainName == "" {
		return ecocredit.ErrParseFailure.Wrap("name cannot be empty")
	}

	return nil
}
