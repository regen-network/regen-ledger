package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/basket/types/v1"
)

func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	govk ecocredit.GovKeeper,
	qryClient basetypes.QueryServer, basketQryClient types.QueryServer,
	authority sdk.AccAddress) simulation.WeightedOperations {

	var (
		weightMsgCreate           int
		weightMsgPut              int
		weightMsgTake             int
		weightMsgUpdateBasketFees int
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCreate, &weightMsgCreate, nil,
		func(_ *rand.Rand) {
			weightMsgCreate = WeightCreate
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgPut, &weightMsgPut, nil,
		func(_ *rand.Rand) {
			weightMsgPut = WeightPut
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgTake, &weightMsgCreate, nil,
		func(_ *rand.Rand) {
			weightMsgTake = WeightTake
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateBasketFee, &weightMsgUpdateBasketFees, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateBasketFees = WeightUpdateBasketFees
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgCreate,
			SimulateMsgCreate(ak, bk, qryClient, basketQryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgPut,
			SimulateMsgPut(ak, bk, qryClient, basketQryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgTake,
			SimulateMsgTake(ak, bk, qryClient, basketQryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateBasketFees,
			SimulateMsgUpdateBasketFee(ak, bk, qryClient, basketQryClient, govk, authority),
		),
	}
}
