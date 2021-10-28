package testsuite

import (
	"testing"

	ecocredit "github.com/regen-network/regen-ledger/x/ecocredit/client/testsuite"
)

func TestEcocreditIntegration(t *testing.T) {
	cfg := DefaultConfig()

	ecocredit.RunCLITests(t, cfg)
}
