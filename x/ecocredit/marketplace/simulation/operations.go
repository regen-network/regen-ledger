package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/types/v1"
)

func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient basetypes.QueryServer, mktQryClient types.QueryServer) simulation.WeightedOperations {

	var (
		weightMsgBuyDirect       int
		weightMsgSell            int
		weightMsgUpdateSellOrder int
		weightMsgCancelSellOrder int
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgBuy, &weightMsgBuyDirect, nil,
		func(_ *rand.Rand) {
			weightMsgBuyDirect = WeightBuyDirect
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgSell, &weightMsgSell, nil,
		func(_ *rand.Rand) {
			weightMsgSell = WeightSell
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateSellOrder, &weightMsgUpdateSellOrder, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSellOrder = WeightUpdateSellOrder
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCancelSellOrder, &weightMsgCancelSellOrder, nil,
		func(_ *rand.Rand) {
			weightMsgCancelSellOrder = WeightCancelSellOrder
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgBuyDirect,
			SimulateMsgBuyDirect(ak, bk, qryClient, mktQryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgSell,
			SimulateMsgSell(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateSellOrder,
			SimulateMsgUpdateSellOrder(ak, bk, qryClient, mktQryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCancelSellOrder,
			SimulateMsgCancelSellOrder(ak, bk, qryClient, mktQryClient),
		),
	}
}
