package basket

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// var feeBasketAcc []byte
var feeModuleAccName = ecocredit.ModuleName + "basket-fees"

func init() {
	// feeBasketAcc = authtypes.NewModuleAddress(ecocredit.ModuleName + "basket-fees")
}
