package core

import (
	"fmt"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
)

// FormatClassID formats the ID to use for a new credit class, based on the credit type and
// sequence number. This format may evolve over time, but will maintain
// backwards compatibility.
//
// The initial version has format:
// <credit type abbreviation><class seq no>
func FormatClassID(creditType v1beta1.CreditType, classSeqNo uint64) string {
	return fmt.Sprintf("%s%02d", creditType.Abbreviation, classSeqNo)
}
