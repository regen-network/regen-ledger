package basketsims

import (
	"math/rand"
	"time"

	gogotypes "github.com/gogo/protobuf/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	regentypes "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

// Simulation operation weights constants
const (
	OpWeightMsgCreate = "op_weight_msg_create_basket"
	OpWeightMsgPut    = "op_weight_msg_put_into_basket"
	OpWeightMsgTake   = "op_weight_take_from_basket"
)

// basket operations weights
const (
	WeightCreate = 100
	WeightPut    = 100
	WeightTake   = 100
)

// ecocredit message types
var (
	TypeMsgCreate = basket.MsgCreate{}.Route()
	TypeMsgPut    = basket.MsgPut{}.Route()
	TypeMsgTake   = basket.MsgTake{}.Route()
)

func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient, basketQryClient basket.QueryClient) simulation.WeightedOperations {

	var (
		weightMsgCreate int
		weightMsgPut    int
		weightMsgTake   int
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

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgCreate,
			SimulateMsgCreate(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgPut,
			SimulateMsgPut(ak, bk, qryClient, basketQryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgTake,
			SimulateMsgTake(ak, bk, qryClient, basketQryClient),
		),
	}
}

// SimulateMsgCreate generates a Basket/MsgCreate with random values.
func SimulateMsgCreate(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		curator, _ := simtypes.RandomAcc(r, accs)

		ctx := regentypes.Context{Context: sdkCtx}
		res, err := qryClient.Params(ctx, &ecocredit.QueryParamsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, err.Error()), nil, err
		}

		params := res.Params
		spendable := bk.SpendableCoins(sdkCtx, curator.Address)
		if spendable.IsAllLTE(params.BasketCreationFee) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, "not enough balance"), nil, nil
		}

		classIds, err := randomClasses(r, ctx, qryClient)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, err.Error()), nil, err
		}

		if len(classIds) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, "no classes"), nil, nil
		}

		creditType, err := randomCreditType(r, ctx, qryClient)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, err.Error()), nil, err
		}

		if creditType == nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, "credit type not found"), nil, nil
		}

		var exponent uint32
		if creditType.Precision <= 1 {
			exponent = creditType.Precision
		} else {
			exponent = uint32(simtypes.RandIntBetween(r, 1, int(creditType.Precision)))
		}

		msg := &basket.MsgCreate{
			Name:              simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 2, 10)),
			DisplayName:       simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 2, 10)),
			Fee:               sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, int64(simtypes.RandIntBetween(r, 0, 5)))),
			DisableAutoRetire: r.Float32() < 0.5,
			Curator:           curator.Address.String(),
			Exponent:          exponent,
			AllowedClasses:    classIds,
			CreditTypeName:    creditType.Name,
			DateCriteria:      randomDateCriteria(r),
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      curator,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func randomDateCriteria(r *rand.Rand) *basket.DateCriteria {
	includeCriteria := r.Int63n(101) <= 50
	if includeCriteria {
		seconds := time.Hour * 24 * 356 * 5
		if includeCriteria {
			return &basket.DateCriteria{
				Sum: &basket.DateCriteria_MinStartDate{
					MinStartDate: &gogotypes.Timestamp{
						Seconds: int64(seconds),
					},
				},
			}
		} else {
			return &basket.DateCriteria{
				Sum: &basket.DateCriteria_StartDateWindow{
					StartDateWindow: gogotypes.DurationProto(seconds),
				},
			}
		}
	} else {
		return &basket.DateCriteria{}
	}
}

// SimulateMsgPut generates a Basket/MsgPut with random values.
func SimulateMsgPut(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient, bsktQryClient basket.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, ""), nil, nil
	}
}

// SimulateMsgTake generates a Basket/MsgTake with random values.
func SimulateMsgTake(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient, bsktQryClient basket.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, ""), nil, nil
	}
}

func randomClasses(r *rand.Rand, ctx regentypes.Context, qryClient ecocredit.QueryClient) ([]string, error) {
	res, err := qryClient.Classes(ctx, &ecocredit.QueryClassesRequest{})
	if err != nil {
		return nil, err
	}

	classes := res.GetClasses()
	if len(classes) == 0 {
		return []string{}, nil
	}

	max := simtypes.RandIntBetween(r, 1, len(classes))

	classIds := make([]string, max)
	for i := 0; i < max; i++ {
		classIds[i] = classes[i].ClassId
	}

	return classIds, nil
}

func randomCreditType(r *rand.Rand, ctx regentypes.Context, qryClient ecocredit.QueryClient) (*ecocredit.CreditType, error) {
	res, err := qryClient.CreditTypes(ctx, &ecocredit.QueryCreditTypesRequest{})
	if err != nil {
		return nil, err
	}

	creditTypes := res.CreditTypes
	if len(creditTypes) == 0 {
		return nil, nil
	}

	return creditTypes[r.Intn(len(creditTypes))], nil
}
