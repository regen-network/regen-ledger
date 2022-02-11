package basket

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var feeBasketAcc []byte

func init() {
	feeBasketAcc = authtypes.NewModuleAddress(ecocredit.ModuleName + "basket-fees")
}
