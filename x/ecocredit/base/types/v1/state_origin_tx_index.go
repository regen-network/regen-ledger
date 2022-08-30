package v1

import (
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Validate performs basic validation of the OriginTxIndex state type
func (m *OriginTxIndex) Validate() error {
	if m.ClassKey == 0 {
		return ecocredit.ErrParseFailure.Wrap("class key cannot be zero")
	}

	if m.Id == "" {
		return ecocredit.ErrParseFailure.Wrap("id cannot be empty")
	}

	if !reOriginTxID.MatchString(m.Id) {
		return ecocredit.ErrParseFailure.Wrap("id must be at most 128 characters long, valid characters: alpha-numberic, space, '-' or '_'")
	}

	if m.Source == "" {
		return ecocredit.ErrParseFailure.Wrap("source cannot be empty")
	}

	if !reOriginTxSource.MatchString(m.Source) {
		return ecocredit.ErrParseFailure.Wrap("source must be at most 32 characters long, valid characters: alpha-numberic, space, '-' or '_'")
	}

	return nil
}
