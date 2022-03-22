package marketplace

import (
	marketApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// TODO: Revisit this once we have proper gas fee framework.
// Tracking issues https://github.com/cosmos/cosmos-sdk/issues/9054, https://github.com/cosmos/cosmos-sdk/discussions/9072
const gasCostPerIteration = uint64(10)

type Keeper struct {
	stateStore    marketApi.StateStore
	coreStore     ecoApi.StateStore
	bankKeeper    ecocredit.BankKeeper
	paramsKeeper  ecocredit.ParamKeeper
	accountKeeper ecocredit.AccountKeeper
}

func NewKeeper(ss marketApi.StateStore, cs ecoApi.StateStore, bk ecocredit.BankKeeper, pk ecocredit.ParamKeeper, ak ecocredit.AccountKeeper) Keeper {
	return Keeper{
		coreStore:     cs,
		stateStore:    ss,
		bankKeeper:    bk,
		paramsKeeper:  pk,
		accountKeeper: ak,
	}
}

// TODO: uncomment when impl
// var _ v1.MsgServer = Keeper{}
// var _ v1.QueryServer = Keeper{}
