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

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	regentypes "github.com/regen-network/regen-ledger/types"
)

// Simulation operation weights constants
const (
	OpWeightMsgCreateClass         = "op_weight_msg_create_class"
	OpWeightMsgCreateBatch         = "op_weight_msg_create_batch"
	OpWeightMsgSend                = "op_weight_msg_send"
	OpWeightMsgRetire              = "op_weight_msg_retire"
	OpWeightMsgCancel              = "op_weight_msg_cancel"
	OpWeightMsgUpdateClassAdmin    = "op_weight_msg_update_class_admin"
	OpWeightMsgUpdateClassMetadata = "op_weight_msg_update_class_metadata"
	OpWeightMsgUpdateClassIssuers  = "op_weight_msg_update_class_issuers"
	OpWeightMsgCreateProject       = "op_weight_msg_create_project"
)

// ecocredit operations weights
const (
	WeightCreateClass   = 10
	WeightCreateProject = 20
	WeightCreateBatch   = 50
	WeightSend          = 100
	WeightRetire        = 80
	WeightCancel        = 30
	WeightUpdateClass   = 30
)

// ecocredit message types
var (
	TypeMsgCreateClass         = sdk.MsgTypeURL(&ecocredit.MsgCreateClass{})
	TypeMsgCreateProject       = sdk.MsgTypeURL(&ecocredit.MsgCreateProject{})
	TypeMsgCreateBatch         = sdk.MsgTypeURL(&ecocredit.MsgCreateBatch{})
	TypeMsgSend                = sdk.MsgTypeURL(&ecocredit.MsgSend{})
	TypeMsgRetire              = sdk.MsgTypeURL(&ecocredit.MsgRetire{})
	TypeMsgCancel              = sdk.MsgTypeURL(&ecocredit.MsgCancel{})
	TypeMsgUpdateClassAdmin    = sdk.MsgTypeURL(&ecocredit.MsgUpdateClassAdmin{})
	TypeMsgUpdateClassIssuers  = sdk.MsgTypeURL(&ecocredit.MsgUpdateClassIssuers{})
	TypeMsgUpdateClassMetadata = sdk.MsgTypeURL(&ecocredit.MsgUpdateClassMetadata{})
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, qryClient ecocredit.QueryClient) simulation.WeightedOperations {

	var (
		weightMsgCreateClass         int
		weightMsgCreateBatch         int
		weightMsgSend                int
		weightMsgRetire              int
		weightMsgCancel              int
		weightMsgUpdateClassAdmin    int
		weightMsgUpdateClassIssuers  int
		weightMsgUpdateClassMetadata int
		weightMsgCreateProject       int
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

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateClassAdmin, &weightMsgUpdateClassAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateClassAdmin = WeightUpdateClass
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateClassIssuers, &weightMsgUpdateClassIssuers, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateClassIssuers = WeightUpdateClass
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateClassMetadata, &weightMsgUpdateClassMetadata, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateClassMetadata = WeightUpdateClass
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCreateProject, &weightMsgCreateProject, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProject = WeightCreateProject
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
		simulation.NewWeightedOperation(
			weightMsgUpdateClassAdmin,
			SimulateMsgUpdateClassAdmin(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateClassIssuers,
			SimulateMsgUpdateClassIssuers(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateClassMetadata,
			SimulateMsgUpdateClassMetadata(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCreateProject,
			SimulateMsgCreateProject(ak, bk, qryClient),
		),
	}
}

// SimulateMsgCreateClass generates a MsgCreateClass with random values.
func SimulateMsgCreateClass(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		admin := accs[0]
		issuers := randomIssuers(r, accs)

		ctx := regentypes.Context{Context: sdkCtx}
		res, err := qryClient.Params(ctx, &ecocredit.QueryParamsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, err.Error()), nil, err
		}

		params := res.Params
		if params.AllowlistEnabled && !contains(params.AllowedClassCreators, admin.Address.String()) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, "not allowed to create credit class"), nil, nil // skip
		}

		spendable := bk.SpendableCoins(sdkCtx, admin.Address)
		if spendable.IsAllLTE(params.CreditClassFee) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, "not enough balance"), nil, nil
		}

		creditTypes := []string{"carbon", "biodiversity"}

		msg := &ecocredit.MsgCreateClass{
			Admin:          admin.Address.String(),
			Issuers:        issuers,
			Metadata:       []byte(simtypes.RandStringOfLength(r, 10)),
			CreditTypeName: creditTypes[r.Intn(len(creditTypes))],
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      admin,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgCreateProject generates a MsgCreateProject with random values.
func SimulateMsgCreateProject(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		issuer := accs[0]

		ctx := regentypes.Context{Context: sdkCtx}
		res, err := qryClient.Classes(ctx, &ecocredit.QueryClassesRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateProject, err.Error()), nil, err
		}

		classes := res.Classes
		if len(classes) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateProject, "no credit classes"), nil, nil
		}

		var classID string
		for _, class := range classes {
			if contains(class.Issuers, issuer.Address.String()) {
				classID = class.ClassId
				break
			}
		}

		if classID == "" {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateProject, "don't have permission to create project"), nil, nil
		}

		issuerAcc := ak.GetAccount(sdkCtx, issuer.Address)
		spendable := bk.SpendableCoins(sdkCtx, issuerAcc.GetAddress())

		msg := &ecocredit.MsgCreateProject{
			Issuer:          issuer.Address.String(),
			ClassId:         classID,
			Metadata:        []byte(simtypes.RandStringOfLength(r, 100)),
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

		return GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgCreateBatch generates a MsgCreateBatch with random values.
func SimulateMsgCreateBatch(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		issuer := accs[0]

		ctx := regentypes.Context{Context: sdkCtx}
		res, err := qryClient.Projects(ctx, &ecocredit.QueryProjectsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateBatch, err.Error()), nil, err
		}

		projects := res.Projects
		if len(projects) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateBatch, "no projects"), nil, nil
		}

		var projectID string
		for _, project := range projects {
			if project.Issuer == issuer.Address.String() {
				projectID = project.ProjectId
				break
			}
		}

		if projectID == "" {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateBatch, "don't have permission to create credit batch"), nil, nil
		}

		issuerAcc := ak.GetAccount(sdkCtx, issuer.Address)
		spendable := bk.SpendableCoins(sdkCtx, issuerAcc.GetAddress())

		now := ctx.BlockTime()
		tenHours := now.Add(10 * time.Hour)
		msg := &ecocredit.MsgCreateBatch{
			Issuer:    issuer.Address.String(),
			ProjectId: projectID,
			Issuance:  generateBatchIssuance(r, accs),
			StartDate: &now,
			EndDate:   &tenHours,
			Metadata:  []byte(simtypes.RandStringOfLength(r, 10)),
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

		return GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgSend generates a MsgSend with random values.
func SimulateMsgSend(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		ctx := regentypes.Context{Context: sdkCtx}
		class, op, err := getRandomClass(ctx, r, qryClient, TypeMsgSend)
		if class == nil {
			return op, nil, err
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgSend, class.ClassId)
		if project == nil {
			return op, nil, err
		}

		batch, op, err := getRandomBatchFromProject(ctx, r, qryClient, TypeMsgSend, class.ClassId)
		if batch == nil {
			return op, nil, err
		}

		balres, err := qryClient.Balance(ctx, &ecocredit.QueryBalanceRequest{
			Account:    project.Issuer,
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
		if project.Issuer == recipient.Address.String() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, "sender & recipient are same"), nil, nil
		}

		addr, err := sdk.AccAddressFromBech32(project.Issuer)
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
			Sender:    project.Issuer,
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

// SimulateMsgRetire generates a MsgRetire with random values.
func SimulateMsgRetire(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		ctx := regentypes.Context{Context: sdkCtx}
		class, op, err := getRandomClass(ctx, r, qryClient, TypeMsgRetire)
		if class == nil {
			return op, nil, err
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgRetire, class.ClassId)
		if project == nil {
			return op, nil, err
		}

		batch, op, err := getRandomBatchFromProject(ctx, r, qryClient, TypeMsgRetire, project.ProjectId)
		if batch == nil {
			return op, nil, err
		}

		balanceRes, err := qryClient.Balance(ctx, &ecocredit.QueryBalanceRequest{
			Account:    project.Issuer,
			BatchDenom: batch.BatchDenom,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, err
		}

		tradableBalance, err := math.NewNonNegativeDecFromString(balanceRes.TradableAmount)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, err.Error()), nil, err
		}

		if tradableBalance.IsZero() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, "balance is zero"), nil, nil
		}

		randSub := math.NewDecFromInt64(int64(simtypes.RandIntBetween(r, 1, 10)))
		spendable, account, op, err := getAccountAndSpendableCoins(sdkCtx, bk, accs, project.Issuer, TypeMsgRetire)
		if spendable == nil {
			return op, nil, err
		}

		if !spendable.IsAllPositive() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, "insufficient funds"), nil, nil
		}

		if tradableBalance.Cmp(randSub) != 1 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, "insufficient funds"), nil, nil
		}

		msg := &ecocredit.MsgRetire{
			Holder: account.Address.String(),
			Credits: []*ecocredit.MsgRetire_RetireCredits{
				{
					BatchDenom: batch.BatchDenom,
					Amount:     randSub.String(),
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
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgCancel generates a MsgCancel with random values.
func SimulateMsgCancel(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		ctx := regentypes.Context{Context: sdkCtx}
		class, op, err := getRandomClass(ctx, r, qryClient, TypeMsgCancel)
		if class == nil {
			return op, nil, err
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgRetire, class.ClassId)
		if project == nil {
			return op, nil, err
		}

		batch, op, err := getRandomBatchFromProject(ctx, r, qryClient, TypeMsgCancel, project.ProjectId)
		if batch == nil {
			return op, nil, err
		}

		spendable, account, op, err := getAccountAndSpendableCoins(sdkCtx, bk, accs, project.Issuer, TypeMsgCancel)
		if spendable == nil {
			return op, nil, err
		}

		balanceRes, err := qryClient.Balance(ctx, &ecocredit.QueryBalanceRequest{
			Account:    project.Issuer,
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
			Holder: project.Issuer,
			Credits: []*ecocredit.MsgCancel_CancelCredits{
				{
					BatchDenom: batch.BatchDenom,
					Amount:     balanceRes.TradableAmount,
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
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgUpdateClassAdmin generates a MsgUpdateClassAdmin with random values
func SimulateMsgUpdateClassAdmin(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		ctx := regentypes.Context{Context: sdkCtx}
		class, op, err := getRandomClass(ctx, r, qryClient, TypeMsgUpdateClassAdmin)
		if class == nil {
			return op, nil, err
		}

		spendable, account, op, err := getAccountAndSpendableCoins(sdkCtx, bk, accs, class.Admin, TypeMsgUpdateClassAdmin)
		if spendable == nil {
			return op, nil, err
		}

		newAdmin, _ := simtypes.RandomAcc(r, accs)
		if newAdmin.Address.String() == class.Admin {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateClassAdmin, "same account"), nil, nil // skip
		}

		msg := &ecocredit.MsgUpdateClassAdmin{
			Admin:    class.Admin,
			ClassId:  class.ClassId,
			NewAdmin: newAdmin.Address.String(),
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgUpdateClassMetadata generates a MsgUpdateClassMetadata with random metadata
func SimulateMsgUpdateClassMetadata(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		ctx := regentypes.Context{Context: sdkCtx}
		class, op, err := getRandomClass(ctx, r, qryClient, TypeMsgUpdateClassMetadata)
		if class == nil {
			return op, nil, err
		}

		spendable, account, op, err := getAccountAndSpendableCoins(sdkCtx, bk, accs, class.Admin, TypeMsgUpdateClassMetadata)
		if spendable == nil {
			return op, nil, err
		}

		msg := &ecocredit.MsgUpdateClassMetadata{
			Admin:    class.Admin,
			ClassId:  class.ClassId,
			Metadata: []byte(simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 10, 256))),
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgUpdateClassIssuers generates a MsgUpdateClassMetaData with random values
func SimulateMsgUpdateClassIssuers(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		ctx := regentypes.Context{Context: sdkCtx}
		class, op, err := getRandomClass(ctx, r, qryClient, TypeMsgUpdateClassIssuers)
		if class == nil {
			return op, nil, err
		}

		spendable, account, op, err := getAccountAndSpendableCoins(sdkCtx, bk, accs, class.Admin, TypeMsgUpdateClassIssuers)
		if spendable == nil {
			return op, nil, err
		}

		issuers := randomIssuers(r, accs)
		msg := &ecocredit.MsgUpdateClassIssuers{
			Admin:   class.Admin,
			ClassId: class.ClassId,
			Issuers: issuers,
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return GenAndDeliverTxWithRandFees(txCtx)
	}
}

func getAccountAndSpendableCoins(ctx sdk.Context, bk ecocredit.BankKeeper,
	accs []simtypes.Account, addr, msgType string) (sdk.Coins, *simtypes.Account, simtypes.OperationMsg, error) {
	accAddr, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, err.Error()), err
	}

	account, found := simtypes.FindAccount(accs, accAddr)
	if !found {
		return nil, &account, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, "account not found"), nil
	}

	spendable := bk.SpendableCoins(ctx, accAddr)
	return spendable, &account, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil

}

func getRandomClass(ctx regentypes.Context, r *rand.Rand, qryClient ecocredit.QueryClient, msgType string) (*ecocredit.ClassInfo, simtypes.OperationMsg, error) {
	res, err := qryClient.Classes(ctx, &ecocredit.QueryClassesRequest{})
	if err != nil {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, err.Error()), err
	}

	classes := res.Classes
	if len(classes) == 0 {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, "no credit class found"), nil
	}

	return classes[r.Intn(len(classes))], simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
}

func getRandomProjectFromClass(ctx regentypes.Context, r *rand.Rand, qryClient ecocredit.QueryClient, msgType, classID string) (*ecocredit.ProjectInfo, simtypes.OperationMsg, error) {
	res, err := qryClient.Projects(ctx, &ecocredit.QueryProjectsRequest{
		ClassId: classID,
	})
	if err != nil {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, err.Error()), err
	}

	projects := res.Projects
	if len(projects) == 0 {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, "no project found"), nil
	}

	return projects[r.Intn(len(projects))], simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
}

func getRandomBatchFromProject(ctx regentypes.Context, r *rand.Rand, qryClient ecocredit.QueryClient, msgType, projectID string) (*ecocredit.BatchInfo, simtypes.OperationMsg, error) {
	res, err := qryClient.Batches(ctx, &ecocredit.QueryBatchesRequest{
		ProjectId: projectID,
	})
	if err != nil {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, err.Error()), err
	}

	batches := res.Batches
	if len(batches) == 0 {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, "no batch found"), nil
	}

	return batches[r.Intn(len(batches))], simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
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
	n := simtypes.RandIntBetween(r, 3, 10)
	issuers := make([]string, n)
	for i := 0; i < n; i++ {
		acc, _ := simtypes.RandomAcc(r, accounts)
		issuers[i] = acc.Address.String()
	}

	return issuers
}

func generateBatchIssuance(r *rand.Rand, accs []simtypes.Account) []*ecocredit.MsgCreateBatch_BatchIssuance {
	numIssuances := simtypes.RandIntBetween(r, 3, 10)
	res := make([]*ecocredit.MsgCreateBatch_BatchIssuance, numIssuances)

	for i := 0; i < numIssuances; i++ {
		recipient := accs[i]
		res[i] = &ecocredit.MsgCreateBatch_BatchIssuance{
			Recipient:          recipient.Address.String(),
			TradableAmount:     fmt.Sprintf("%d", simtypes.RandIntBetween(r, 500, 100000)),
			RetiredAmount:      fmt.Sprintf("%d", simtypes.RandIntBetween(r, 500, 10000)),
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
		10000000,
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
		return simtypes.NoOpMsg(txCtx.ModuleName, txCtx.MsgType, "unable to deliver tx"), nil, err
	}

	return simtypes.NewOperationMsg(txCtx.Msg, true, "", txCtx.Cdc), nil, nil
}
