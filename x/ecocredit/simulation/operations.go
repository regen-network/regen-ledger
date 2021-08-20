package simulation

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/exported"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	regentypes "github.com/regen-network/regen-ledger/types"
)

// Simulation operation weights constants
const (
	OpWeightMsgCreateClass = "op_weight_msg_create_class"
	OpWeightMsgCreateBatch = "op_weight_msg_create_batch"
	OpWeightMsgSend        = "op_weight_msg_send"
	OpWeightMsgRetire      = "op_weight_msg_retire"
	OpWeightMsgCancel      = "op_weight_msg_cancel"
)

// ecocredit operations weights
const (
	WeightCreateClass = 100
	WeightCreateBatch = 80
	WeightSend        = 60
	WeightRetire      = 60
	WeightCancel      = 60
)

// ecocredit message types
var (
	TypeMsgCreateClass = sdk.MsgTypeURL(&ecocredit.MsgCreateClass{})
	TypeMsgCreateBatch = sdk.MsgTypeURL(&ecocredit.MsgCreateBatch{})
	TypeMsgSend        = sdk.MsgTypeURL(&ecocredit.MsgSend{})
	TypeMsgRetire      = sdk.MsgTypeURL(&ecocredit.MsgRetire{})
	TypeMsgCancel      = sdk.MsgTypeURL(&ecocredit.MsgCancel{})
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	ak exported.AccountKeeper, bk exported.BankKeeper, qryClient ecocredit.QueryClient) simulation.WeightedOperations {

	var (
		weightMsgCreateClass int
		weightMsgCreateBatch int
		weightMsgSend        int
		weightMsgRetire      int
		weightMsgCancel      int
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCreateClass, &weightMsgCreateClass, nil,
		func(_ *rand.Rand) {
			weightMsgCreateClass = WeightCreateClass
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCreateBatch, &weightMsgCreateBatch, nil,
		func(_ *rand.Rand) {
			weightMsgCreateBatch = WeightCreateBatch
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgSend, &weightMsgSend, nil,
		func(_ *rand.Rand) {
			weightMsgSend = WeightSend
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgRetire, &weightMsgRetire, nil,
		func(_ *rand.Rand) {
			weightMsgRetire = WeightRetire
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCancel, &weightMsgCancel, nil,
		func(_ *rand.Rand) {
			weightMsgCancel = WeightCancel
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgCreateClass,
			SimulateMsgCreateClass(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCreateBatch,
			SimulateMsgCreateBatch(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgSend,
			SimulateMsgSend(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgRetire,
			SimulateMsgRetire(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCancel,
			SimulateMsgCancel(ak, bk, qryClient),
		),
	}
}

// SimulateMsgCreateClass generates a MsgCreateClass with random values.
func SimulateMsgCreateClass(ak exported.AccountKeeper, bk exported.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		designer := accs[0]
		issuers := randomIssuers(r, accs)

		ctx := regentypes.Context{Context: sdkCtx}
		res, err := qryClient.Params(ctx, &ecocredit.QueryParamsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, err.Error()), nil, err
		}

		params := res.Params
		if params.AllowlistEnabled && !contains(params.AllowedClassDesigners, designer.Address.String()) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, "not allowed to create credit class"), nil, nil // skip
		}

		spendable := bk.SpendableCoins(sdkCtx, designer.Address)
		if spendable.IsAllLTE(params.CreditClassFee) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, "not enough balance"), nil, nil
		}

		msg := &ecocredit.MsgCreateClass{
			Designer:   designer.Address.String(),
			Issuers:    issuers,
			Metadata:   []byte(simtypes.RandStringOfLength(r, 10)),
			CreditType: "carbon",
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      designer,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgCreateBatch(ak exported.AccountKeeper, bk exported.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		issuer := accs[0]

		ctx := regentypes.Context{Context: sdkCtx}
		res, err := qryClient.Classes(ctx, &ecocredit.QueryClassesRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateBatch, err.Error()), nil, err
		}

		classes := res.Classes
		if len(classes) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateBatch, "no credit classes"), nil, nil
		}

		var classID string
		for _, class := range classes {
			if contains(class.Issuers, issuer.Address.String()) {
				classID = class.ClassId
				break
			}
		}

		if classID == "" {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateBatch, "don't have permission to create credit batch"), nil, nil
		}

		issuerAcc := ak.GetAccount(sdkCtx, issuer.Address)
		spendable := bk.SpendableCoins(sdkCtx, issuerAcc.GetAddress())

		now := ctx.BlockTime()
		tenHours := now.Add(10 * time.Hour)
		msg := &ecocredit.MsgCreateBatch{
			Issuer:          issuer.Address.String(),
			ClassId:         classID,
			Issuance:        generateBatchIssuence(r, accs),
			StartDate:       &now,
			EndDate:         &tenHours,
			Metadata:        []byte(simtypes.RandStringOfLength(r, 10)),
			ProjectLocation: "AB-CDE FG1 345",
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      issuer,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgSend(ak exported.AccountKeeper, bk exported.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		ctx := regentypes.Context{Context: sdkCtx}
		res, err := qryClient.Classes(ctx, &ecocredit.QueryClassesRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, err
		}

		classes := res.Classes
		if len(classes) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, "no credit classes"), nil, nil
		}
		index := r.Intn(len(classes))
		class := classes[index]

		res1, err := qryClient.Batches(ctx, &ecocredit.QueryBatchesRequest{
			ClassId: class.ClassId,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, err
		}

		batches := res1.Batches
		if len(batches) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, "no credit batches"), nil, nil
		}

		index = r.Intn(len(batches))
		batch := batches[index]

		balres, err := qryClient.Balance(ctx, &ecocredit.QueryBalanceRequest{
			Account:    batch.Issuer,
			BatchDenom: batch.BatchDenom,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, err
		}

		tradableBalance, err := math.NewNonNegativeDecFromString(balres.TradableAmount)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, err
		}

		retiredBalance, err := math.NewNonNegativeDecFromString(balres.RetiredAmount)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, err
		}

		if tradableBalance.IsZero() || retiredBalance.IsZero() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, "balance is zero"), nil, nil
		}

		recipient, _ := simtypes.RandomAcc(r, accs)
		if batch.Issuer == recipient.Address.String() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, "sender & reciever are same"), nil, nil
		}

		addr, err := sdk.AccAddressFromBech32(batch.Issuer)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, err
		}

		acc, found := simtypes.FindAccount(accs, addr)
		if !found {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, "account not found"), nil, nil
		}

		randSub := simtypes.RandIntBetween(r, 1, 100)
		issuer := ak.GetAccount(sdkCtx, acc.Address)
		spendable := bk.SpendableCoins(sdkCtx, issuer.GetAddress())

		msg := &ecocredit.MsgSend{
			Sender:    batch.Issuer,
			Recipient: recipient.Address.String(),
			Credits: []*ecocredit.MsgSend_SendCredits{
				{
					BatchDenom:         batch.BatchDenom,
					TradableAmount:     math.NewDecFromInt64(int64(randSub)).String(),
					RetiredAmount:      math.NewDecFromInt64(int64(randSub)).String(),
					RetirementLocation: "ST-UVW XY Z12",
				},
			},
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      acc,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgRetire(ak exported.AccountKeeper, bk exported.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		ctx := regentypes.Context{Context: sdkCtx}
		res, err := qryClient.Classes(ctx, &ecocredit.QueryClassesRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, err.Error()), nil, err
		}

		classes := res.Classes
		if len(classes) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, "no credit classes"), nil, nil
		}
		index := r.Intn(len(classes))
		class := classes[index]

		res1, err := qryClient.Batches(ctx, &ecocredit.QueryBatchesRequest{
			ClassId: class.ClassId,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, err.Error()), nil, err
		}

		batches := res1.Batches
		if len(batches) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, "no credit batches"), nil, nil
		}
		index = r.Intn(len(batches))
		batch := batches[index]

		randSub := simtypes.RandIntBetween(r, 1, 100)
		balanceRes, err := qryClient.Balance(ctx, &ecocredit.QueryBalanceRequest{
			Account:    batch.Issuer,
			BatchDenom: batch.BatchDenom,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, err
		}

		retiredBalance, err := math.NewNonNegativeDecFromString(balanceRes.RetiredAmount)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, err
		}

		if retiredBalance.IsZero() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, "balance is zero"), nil, nil
		}

		addr, err := sdk.AccAddressFromBech32(batch.Issuer)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, err.Error()), nil, err
		}

		holder, found := simtypes.FindAccount(accs, addr)
		if !found {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, "account not found"), nil, nil
		}

		holderAcc := ak.GetAccount(sdkCtx, holder.Address)
		spendable := bk.SpendableCoins(sdkCtx, holderAcc.GetAddress())

		msg := &ecocredit.MsgRetire{
			Holder: holder.Address.String(),
			Credits: []*ecocredit.MsgRetire_RetireCredits{
				{
					BatchDenom: batch.BatchDenom,
					Amount:     math.NewDecFromInt64(int64(randSub)).String(),
				},
			},
			Location: "ST-UVW XY Z12",
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      holder,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgCancel(ak exported.AccountKeeper, bk exported.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		ctx := regentypes.Context{Context: sdkCtx}
		res, err := qryClient.Classes(ctx, &ecocredit.QueryClassesRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancel, err.Error()), nil, err
		}

		classes := res.Classes
		if len(classes) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancel, "no credit classes"), nil, nil
		}
		index := r.Intn(len(classes))
		class := classes[index]

		res1, err := qryClient.Batches(ctx, &ecocredit.QueryBatchesRequest{
			ClassId: class.ClassId,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancel, err.Error()), nil, err
		}

		batches := res1.Batches
		if len(batches) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancel, "no credit batches"), nil, nil
		}
		index = r.Intn(len(batches))
		batch := batches[index]

		addr, err := sdk.AccAddressFromBech32(batch.Issuer)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancel, err.Error()), nil, err
		}

		acc, found := simtypes.FindAccount(accs, addr)
		if !found {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancel, "account not found"), nil, nil
		}

		balanceRes, err := qryClient.Balance(ctx, &ecocredit.QueryBalanceRequest{
			Account:    batch.Issuer,
			BatchDenom: batch.BatchDenom,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancel, err.Error()), nil, err
		}

		tradableBalance, err := math.NewNonNegativeDecFromString(balanceRes.TradableAmount)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancel, err.Error()), nil, err
		}

		if tradableBalance.IsZero() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancel, "balance is zero"), nil, nil
		}

		msg := &ecocredit.MsgCancel{
			Holder: batch.Issuer,
			Credits: []*ecocredit.MsgCancel_CancelCredits{
				{
					BatchDenom: batch.BatchDenom,
					Amount:     balanceRes.TradableAmount,
				},
			},
		}

		spendable := bk.SpendableCoins(sdkCtx, acc.Address)
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      acc,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func randomIssuers(r *rand.Rand, accounts []simtypes.Account) []string {
	n := simtypes.RandIntBetween(r, 3, 15)
	issuers := make([]string, n)
	for i := 0; i < n; i++ {
		acc, _ := simtypes.RandomAcc(r, accounts)
		issuers[i] = acc.Address.String()
	}

	return issuers
}

func generateBatchIssuence(r *rand.Rand, accs []simtypes.Account) []*ecocredit.MsgCreateBatch_BatchIssuance {
	numIssuences := simtypes.RandIntBetween(r, 3, 12)
	res := make([]*ecocredit.MsgCreateBatch_BatchIssuance, numIssuences)

	for i := 0; i < numIssuences; i++ {
		recipient := accs[i]
		res[i] = &ecocredit.MsgCreateBatch_BatchIssuance{
			Recipient:          recipient.Address.String(),
			TradableAmount:     "12345.123123",
			RetiredAmount:      "1245.123123",
			RetirementLocation: "RD",
		}
	}

	return res
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
		2000000000,
		txCtx.Context.ChainID(),
		[]uint64{account.GetAccountNumber()},
		[]uint64{account.GetSequence()},
		txCtx.SimAccount.PrivKey,
	)

	if err != nil {
		return simtypes.NoOpMsg(txCtx.ModuleName, txCtx.MsgType, "unable to generate mock tx"), nil, err
	}

	_, _, err = txCtx.App.Deliver(txCtx.TxGen.TxEncoder(), tx)
	if err != nil {
		fmt.Println(txCtx.Msg)
		return simtypes.NoOpMsg(txCtx.ModuleName, txCtx.MsgType, "unable to deliver tx"), nil, err
	}

	return simtypes.NewOperationMsg(txCtx.Msg, true, "", txCtx.Cdc), nil, nil
}
