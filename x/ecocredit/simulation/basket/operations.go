package basketsims

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	gogotypes "github.com/gogo/protobuf/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	regentypes "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
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

		precision := int(creditType.Precision)
		exponent := uint32(simtypes.RandIntBetween(r, precision, precision+5))

		dateCriteria := randomDateCriteria(r)
		dateCriteria = nil
		msg := &basket.MsgCreate{
			Name:              simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 3, 32)),
			DisplayName:       simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 3, 32)),
			Fee:               params.BasketCreationFee,
			DisableAutoRetire: r.Float32() < 0.5,
			Curator:           curator.Address.String(),
			Exponent:          exponent,
			AllowedClasses:    classIds,
			CreditTypeName:    creditType.Abbreviation,
			DateCriteria:      dateCriteria,
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
		seconds := time.Hour * 24
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
		return nil
	}
}

// SimulateMsgPut generates a Basket/MsgPut with random values.
func SimulateMsgPut(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient, bsktQryClient basket.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		owner, _ := simtypes.RandomAcc(r, accs)

		ctx := regentypes.Context{Context: sdkCtx}
		res, err := bsktQryClient.Baskets(ctx, &basket.QueryBasketsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, err.Error()), nil, err
		}

		baskets := res.Baskets
		if len(baskets) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, "no baskets"), nil, nil
		}

		classesRes, err := qryClient.Classes(ctx, &ecocredit.QueryClassesRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, err.Error()), nil, err
		}

		classes := classesRes.Classes
		if len(classes) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, "no classes"), nil, nil
		}

		rBasket := baskets[r.Intn(len(baskets))]
		var classInfoList []ecocredit.ClassInfo
		max := 0

		r.Shuffle(len(classes), func(i, j int) { classes[i], classes[j] = classes[j], classes[i] })
		for _, class := range classes {
			if class.CreditType.Abbreviation == rBasket.CreditTypeName {
				classInfoList = append(classInfoList, *class)
				if max == 3 {
					break
				}
				max++
			}
		}

		if len(classInfoList) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, "different credit type"), nil, nil
		}

		var credits []*basket.BasketCredit
		for _, info := range classInfoList {
			batchesRes, err := qryClient.Batches(ctx, &ecocredit.QueryBatchesRequest{ClassId: info.ClassId})
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, err.Error()), nil, err
			}

			batches := batchesRes.Batches
			if len(batches) != 0 {
				for _, item := range batches {
					balanceRes, err := qryClient.Balance(ctx, &ecocredit.QueryBalanceRequest{
						Account: owner.Address.String(), BatchDenom: item.BatchDenom,
					})
					if err != nil {
						return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, err.Error()), nil, err
					}

					tradableAmount := balanceRes.TradableAmount
					if tradableAmount != "0" {
						d, err := math.NewPositiveDecFromString(tradableAmount)
						if err != nil {
							return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, err.Error()), nil, err
						}

						dInt, err := d.Int64()
						if err != nil {
							return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, err.Error()), nil, err
						}

						if dInt == 1 {
							credits = append(credits, &basket.BasketCredit{
								BatchDenom: item.BatchDenom,
								Amount:     "1",
							})
						} else {
							amt := simtypes.RandIntBetween(r, 1, int(dInt))
							credits = append(credits, &basket.BasketCredit{
								BatchDenom: item.BatchDenom,
								Amount:     fmt.Sprintf("%d", amt),
							})

						}
					}
				}
			}
		}
		if len(credits) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, "no basket credits"), nil, nil
		}

		msg := &basket.MsgPut{
			Owner:       owner.Address.String(),
			BasketDenom: rBasket.BasketDenom,
			Credits:     credits,
		}
		spendable := bk.SpendableCoins(sdkCtx, owner.Address)
		fees, err := simtypes.RandomFees(r, sdkCtx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, "fee error"), nil, err
		}

		account := ak.GetAccount(sdkCtx, owner.Address)
		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			owner.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			if strings.Contains(err.Error(), "is not allowed in this basket") {
				return simtypes.NoOpMsg(ecocredit.ModuleName, msg.Type(), "class is not allowed"), nil, nil
			}

			return simtypes.NoOpMsg(ecocredit.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}

// SimulateMsgTake generates a Basket/MsgTake with random values.
func SimulateMsgTake(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient, bsktQryClient basket.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		owner, _ := simtypes.RandomAcc(r, accs)
		ownerAddr := owner.Address.String()

		ctx := regentypes.Context{Context: sdkCtx}
		res, err := bsktQryClient.Baskets(ctx, &basket.QueryBasketsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, err.Error()), nil, err
		}

		baskets := res.Baskets
		if len(baskets) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, "no baskets"), nil, nil
		}

		rBasket := baskets[r.Intn(len(baskets))]
		balancesRes, err := bsktQryClient.BasketBalances(ctx, &basket.QueryBasketBalancesRequest{
			BasketDenom: rBasket.BasketDenom,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, err.Error()), nil, err
		}
		balances := balancesRes.Balances

		if len(balances) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, "no balances"), nil, nil
		}

		balance := balances[r.Intn(len(balances))]

		balanceRes, err := qryClient.Balance(ctx, &ecocredit.QueryBalanceRequest{Account: ownerAddr, BatchDenom: balance.BatchDenom})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, err.Error()), nil, err
		}

		var amt string
		if rBasket.DisableAutoRetire {
			amt = balanceRes.TradableAmount
		} else {
			amt = balanceRes.RetiredAmount
		}

		if amt == "0" {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, "no balances"), nil, nil
		}

		msg := &basket.MsgTake{
			Owner:              ownerAddr,
			BasketDenom:        rBasket.BasketDenom,
			Amount:             amt,
			RetirementLocation: "AQ",
			RetireOnTake:       !rBasket.DisableAutoRetire,
		}

		spendable := bk.SpendableCoins(sdkCtx, owner.Address)
		fees, err := simtypes.RandomFees(r, sdkCtx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, "fee error"), nil, err
		}

		account := ak.GetAccount(sdkCtx, owner.Address)
		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			owner.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, "unable to generate mock tx"), nil, err
		}
		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			if strings.Contains(err.Error(), "insufficient funds") {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, "not enough balance"), nil, nil
			}
			return simtypes.NoOpMsg(ecocredit.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
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
	} else if len(classes) == 1 {
		return []string{classes[0].ClassId}, nil
	}

	max := simtypes.RandIntBetween(r, 1, len(classes))

	r.Shuffle(len(classes), func(i, j int) { classes[i], classes[j] = classes[j], classes[i] })
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

// GenAndDeliverTxWithRandFees generates a transaction with a random fee and delivers it.
func GenAndDeliverTxWithRandFees(txCtx simulation.OperationInput) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
	account := txCtx.AccountKeeper.GetAccount(txCtx.Context, txCtx.SimAccount.Address)
	spendable := txCtx.Bankkeeper.SpendableCoins(txCtx.Context, account.GetAddress())

	var fees sdk.Coins
	var err error

	coins, hasNeg := spendable.SafeSub(txCtx.CoinsSpentInMsg)
	if hasNeg {
		return simtypes.NoOpMsg(txCtx.ModuleName, txCtx.MsgType, "message doesn't leave room for fees"), nil, err
	}

	fees, err = simtypes.RandomFees(txCtx.R, txCtx.Context, coins)
	if err != nil {
		return simtypes.NoOpMsg(txCtx.ModuleName, txCtx.MsgType, "unable to generate fees"), nil, err
	}
	return GenAndDeliverTx(txCtx, fees)
}

// GenAndDeliverTx generates a transactions and delivers it.
func GenAndDeliverTx(txCtx simulation.OperationInput, fees sdk.Coins) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
	account := txCtx.AccountKeeper.GetAccount(txCtx.Context, txCtx.SimAccount.Address)
	tx, err := helpers.GenTx(
		txCtx.TxGen,
		[]sdk.Msg{txCtx.Msg},
		fees,
		helpers.DefaultGenTxGas,
		txCtx.Context.ChainID(),
		[]uint64{account.GetAccountNumber()},
		[]uint64{account.GetSequence()},
		txCtx.SimAccount.PrivKey,
	)

	if err != nil {
		return simtypes.NoOpMsg(txCtx.ModuleName, txCtx.MsgType, "unable to generate mock tx"), nil, err
	}

	_, r1, err := txCtx.App.Deliver(txCtx.TxGen.TxEncoder(), tx)
	fmt.Println(r1)
	if err != nil {
		return simtypes.NoOpMsg(txCtx.ModuleName, txCtx.MsgType, "unable to deliver tx"), nil, err
	}

	return simtypes.NewOperationMsg(txCtx.Msg, true, "", txCtx.Cdc), nil, nil

}
