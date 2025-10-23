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
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/marketplace/types/v1"
)

func WeightedOperations(
	appParams simtypes.AppParams,
	txCfg client.TxConfig,
	ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient basetypes.QueryServer, mktQryClient types.QueryServer,
	govk govkeeper.Keeper, authority sdk.AccAddress,
) simulation.WeightedOperations {
	var (
		weightMsgBuyDirect          int
		weightMsgSell               int
		weightMsgUpdateSellOrder    int
		weightMsgCancelSellOrder    int
		weightMsgAddAllowedDenom    int
		weightMsgRemoveAllowedDenom int
	)

	appParams.GetOrGenerate(OpWeightMsgBuy, &weightMsgBuyDirect, nil,
		func(_ *rand.Rand) {
			weightMsgBuyDirect = WeightBuyDirect
		},
	)

	appParams.GetOrGenerate(OpWeightMsgSell, &weightMsgSell, nil,
		func(_ *rand.Rand) {
			weightMsgSell = WeightSell
		},
	)

	appParams.GetOrGenerate(OpWeightMsgUpdateSellOrder, &weightMsgUpdateSellOrder, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSellOrder = WeightUpdateSellOrder
		},
	)

	appParams.GetOrGenerate(OpWeightMsgCancelSellOrder, &weightMsgCancelSellOrder, nil,
		func(_ *rand.Rand) {
			weightMsgCancelSellOrder = WeightCancelSellOrder
		},
	)

	appParams.GetOrGenerate(OpWeightMsgAddAllowedDenom, &weightMsgAddAllowedDenom, nil,
		func(_ *rand.Rand) {
			weightMsgAddAllowedDenom = WeightAddAllowedDenom
		},
	)

	appParams.GetOrGenerate(OpWeightMsgRemoveAllowedDenom, &weightMsgRemoveAllowedDenom, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveAllowedDenom = WeightRemoveAllowedDenom
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgBuyDirect,
			SimulateMsgBuyDirect(txCfg, ak, bk, qryClient, mktQryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgSell,
			SimulateMsgSell(txCfg, ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateSellOrder,
			SimulateMsgUpdateSellOrder(txCfg, ak, bk, qryClient, mktQryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCancelSellOrder,
			SimulateMsgCancelSellOrder(txCfg, ak, bk, qryClient, mktQryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgAddAllowedDenom,
			SimulateMsgAddAllowedDenom(txCfg, ak, bk, mktQryClient, govk, authority),
		),
		simulation.NewWeightedOperation(
			weightMsgRemoveAllowedDenom,
			SimulateMsgRemoveAllowedDenom(txCfg, ak, bk, mktQryClient, govk, authority),
		),
	}
}
