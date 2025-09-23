package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/x/ecocredit/v4"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/basket/types/v1"
)

func WeightedOperations(
	appParams simtypes.AppParams,
	txCfg client.TxConfig,
	ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	govk govkeeper.Keeper,
	qryClient basetypes.QueryServer, basketQryClient types.QueryServer,
	authority sdk.AccAddress,
) simulation.WeightedOperations {
	var (
		weightMsgCreate           int
		weightMsgPut              int
		weightMsgTake             int
		weightMsgUpdateBasketFees int
	)

	appParams.GetOrGenerate(OpWeightMsgCreate, &weightMsgCreate, nil,
		func(_ *rand.Rand) {
			weightMsgCreate = WeightCreate
		},
	)

	appParams.GetOrGenerate(OpWeightMsgPut, &weightMsgPut, nil,
		func(_ *rand.Rand) {
			weightMsgPut = WeightPut
		},
	)

	appParams.GetOrGenerate(OpWeightMsgTake, &weightMsgCreate, nil,
		func(_ *rand.Rand) {
			weightMsgTake = WeightTake
		},
	)

	appParams.GetOrGenerate(OpWeightMsgUpdateBasketFee, &weightMsgUpdateBasketFees, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateBasketFees = WeightUpdateBasketFees
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgCreate,
			SimulateMsgCreate(txCfg, ak, bk, qryClient, basketQryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgPut,
			SimulateMsgPut(txCfg, ak, bk, qryClient, basketQryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgTake,
			SimulateMsgTake(txCfg, ak, bk, qryClient, basketQryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateBasketFees,
			SimulateMsgUpdateBasketFee(txCfg, ak, bk, qryClient, basketQryClient, govk, authority),
		),
	}
}
