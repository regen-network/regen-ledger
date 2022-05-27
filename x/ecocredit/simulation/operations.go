package simulation

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	regentypes "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	basketsims "github.com/regen-network/regen-ledger/x/ecocredit/simulation/basket"
	marketplacesims "github.com/regen-network/regen-ledger/x/ecocredit/simulation/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation/utils"
)

// Simulation operation weights constants
const (
	OpWeightMsgCreateClass           = "op_weight_msg_create_class"
	OpWeightMsgCreateBatch           = "op_weight_msg_create_batch"
	OpWeightMsgSend                  = "op_weight_msg_send"
	OpWeightMsgRetire                = "op_weight_msg_retire"
	OpWeightMsgCancel                = "op_weight_msg_cancel"
	OpWeightMsgUpdateClassAdmin      = "op_weight_msg_update_class_admin"
	OpWeightMsgUpdateClassMetadata   = "op_weight_msg_update_class_metadata"
	OpWeightMsgUpdateClassIssuers    = "op_weight_msg_update_class_issuers"
	OpWeightMsgCreateProject         = "op_weight_msg_create_project"
	OpWeightMsgUpdateProjectAdmin    = "op_weight_msg_update_project_admin"
	OpWeightMsgUpdateProjectMetadata = "op_weight_msg_update_project_metadata"
)

// ecocredit operations weights
const (
	WeightCreateClass           = 10
	WeightCreateProject         = 20
	WeightCreateBatch           = 50
	WeightSend                  = 100
	WeightRetire                = 80
	WeightCancel                = 30
	WeightUpdateClass           = 30
	WeightUpdateProjectAdmin    = 30
	WeightUpdateProjectMetadata = 30
)

// ecocredit message types
var (
	TypeMsgCreateClass           = sdk.MsgTypeURL(&core.MsgCreateClass{})
	TypeMsgCreateProject         = sdk.MsgTypeURL(&core.MsgCreateProject{})
	TypeMsgCreateBatch           = sdk.MsgTypeURL(&core.MsgCreateBatch{})
	TypeMsgSend                  = sdk.MsgTypeURL(&core.MsgSend{})
	TypeMsgRetire                = sdk.MsgTypeURL(&core.MsgRetire{})
	TypeMsgCancel                = sdk.MsgTypeURL(&core.MsgCancel{})
	TypeMsgUpdateClassAdmin      = sdk.MsgTypeURL(&core.MsgUpdateClassAdmin{})
	TypeMsgUpdateClassIssuers    = sdk.MsgTypeURL(&core.MsgUpdateClassIssuers{})
	TypeMsgUpdateClassMetadata   = sdk.MsgTypeURL(&core.MsgUpdateClassMetadata{})
	TypeMsgUpdateProjectMetadata = sdk.MsgTypeURL(&core.MsgUpdateProjectMetadata{})
	TypeMsgUpdateProjectAdmin    = sdk.MsgTypeURL(&core.MsgUpdateProjectAdmin{})
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryClient, basketQryClient basket.QueryClient,
	mktQryClient marketplace.QueryClient) simulation.WeightedOperations {

	var (
		weightMsgCreateClass           int
		weightMsgCreateBatch           int
		weightMsgSend                  int
		weightMsgRetire                int
		weightMsgCancel                int
		weightMsgUpdateClassAdmin      int
		weightMsgUpdateClassIssuers    int
		weightMsgUpdateClassMetadata   int
		weightMsgCreateProject         int
		weightMsgUpdateProjectMetadata int
		weightMsgUpdateProjectAdmin    int
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCreateClass, &weightMsgCreateClass, nil,
		func(_ *rand.Rand) {
			weightMsgCreateClass = WeightCreateClass
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCreateProject, &weightMsgCreateProject, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProject = WeightCreateProject
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

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateProjectAdmin, &weightMsgUpdateProjectAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProjectAdmin = WeightUpdateProjectAdmin
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateProjectMetadata, &weightMsgUpdateProjectMetadata, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProjectMetadata = WeightUpdateProjectMetadata
		},
	)

	ops := simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgCreateClass,
			SimulateMsgCreateClass(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCreateProject,
			SimulateMsgCreateProject(ak, bk, qryClient),
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
			weightMsgUpdateClassAdmin,
			SimulateMsgUpdateProjectAdmin(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateProjectMetadata,
			SimulateMsgUpdateProjectMetadata(ak, bk, qryClient),
		),
	}

	basketOps := basketsims.WeightedOperations(appParams, cdc, ak, bk, qryClient, basketQryClient)
	marketplaceOps := marketplacesims.WeightedOperations(appParams, cdc, ak, bk, qryClient, mktQryClient)

	ops = append(ops, basketOps...)
	return append(ops, marketplaceOps...)
}

// SimulateMsgUpdateProjectMetadata generates a MsgUpdateProjectMetadata with random values.
func SimulateMsgUpdateProjectMetadata(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, qryClient core.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgUpdateProjectMetadata)
		if err != nil {
			return op, nil, err
		}

		ctx := regentypes.Context{Context: sdkCtx}
		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgUpdateProjectMetadata, class.Id)
		if project == nil {
			return op, nil, err
		}

		admin, err := sdk.AccAddressFromBech32(project.Admin)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateProjectMetadata, err.Error()), nil, err
		}

		msg := &core.MsgUpdateProjectMetadata{
			Admin:       admin.String(),
			ProjectId:   project.Id,
			NewMetadata: simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 10, core.MaxMetadataLength)),
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, admin.String(), TypeMsgUpdateProjectMetadata)
		if spendable == nil {
			return op, nil, err
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

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgUpdateProjectAdmin generates a MsgUpdateProjectAdmin with random values.
func SimulateMsgUpdateProjectAdmin(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, qryClient core.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgUpdateProjectAdmin)
		if err != nil {
			return op, nil, err
		}

		ctx := regentypes.Context{Context: sdkCtx}
		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgUpdateProjectAdmin, class.Id)
		if project == nil {
			return op, nil, err
		}

		newAdmin, _ := simtypes.RandomAcc(r, accs)
		if project.Admin == newAdmin.Address.String() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateProjectAdmin, "old and new admin are same"), nil, nil
		}

		msg := &core.MsgUpdateProjectAdmin{
			Admin:     project.Admin,
			NewAdmin:  newAdmin.Address.String(),
			ProjectId: project.Id,
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, project.Admin, TypeMsgUpdateProjectAdmin)
		if spendable == nil {
			return op, nil, err
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

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgCreateClass generates a MsgCreateClass with random values.
func SimulateMsgCreateClass(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		admin, _ := simtypes.RandomAcc(r, accs)
		issuers := randomIssuers(r, accs)

		ctx := regentypes.Context{Context: sdkCtx}
		res, err := qryClient.Params(ctx, &core.QueryParamsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, err.Error()), nil, err
		}

		params := res.Params
		if params.AllowlistEnabled && !utils.Contains(params.AllowedClassCreators, admin.Address.String()) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, "not allowed to create credit class"), nil, nil // skip
		}

		spendable, neg := bk.SpendableCoins(sdkCtx, admin.Address).SafeSub(params.CreditClassFee)
		if neg {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, "not enough balance"), nil, nil
		}

		creditTypes := []string{"C", "BIO"}
		msg := &core.MsgCreateClass{
			Admin:            admin.Address.String(),
			Issuers:          issuers,
			Metadata:         simtypes.RandStringOfLength(r, 10),
			CreditTypeAbbrev: creditTypes[r.Intn(len(creditTypes))],
			Fee:              &params.CreditClassFee[0],
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

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgCreateProject generates a MsgCreateProject with random values.
func SimulateMsgCreateProject(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		issuer, _ := simtypes.RandomAcc(r, accs)

		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgCreateProject)
		if class == nil {
			return op, nil, err
		}

		issuers, op, err := getClassIssuers(sdkCtx, r, qryClient, class.Id, TypeMsgCreateProject)
		if len(issuers) == 0 {
			return op, nil, err
		}

		if !utils.Contains(issuers, issuer.Address.String()) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateProject, "don't have permission to create project"), nil, nil
		}

		issuerAcc := ak.GetAccount(sdkCtx, issuer.Address)
		spendable := bk.SpendableCoins(sdkCtx, issuerAcc.GetAddress())

		msg := &core.MsgCreateProject{
			Issuer:       issuer.Address.String(),
			ClassId:      class.Id,
			Metadata:     simtypes.RandStringOfLength(r, 100),
			Jurisdiction: "AB-CDE FG1 345",
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

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgCreateBatch generates a MsgCreateBatch with random values.
func SimulateMsgCreateBatch(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		issuer, _ := simtypes.RandomAcc(r, accs)

		ctx := regentypes.Context{Context: sdkCtx}
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgCreateBatch)
		if class == nil {
			return op, nil, err
		}

		result, err := qryClient.ClassIssuers(ctx, &core.QueryClassIssuersRequest{ClassId: class.Id})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateBatch, err.Error()), nil, err
		}

		classIssuers := result.Issuers
		if len(classIssuers) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateBatch, "no issuers"), nil, nil
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgCreateBatch, class.Id)
		if project == nil {
			return op, nil, err
		}

		if !utils.Contains(classIssuers, issuer.Address.String()) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateBatch, "don't have permission to create credit batch"), nil, nil
		}

		issuerAcc := ak.GetAccount(sdkCtx, issuer.Address)
		spendable := bk.SpendableCoins(sdkCtx, issuerAcc.GetAddress())

		now := sdkCtx.BlockTime()
		tenHours := now.Add(10 * time.Hour)

		msg := &core.MsgCreateBatch{
			Issuer:    issuer.Address.String(),
			ProjectId: project.Id,
			Issuance:  generateBatchIssuance(r, accs),
			StartDate: &now,
			EndDate:   &tenHours,
			Metadata:  simtypes.RandStringOfLength(r, 10),
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

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgSend generates a MsgSend with random values.
func SimulateMsgSend(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		ctx := regentypes.Context{Context: sdkCtx}
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgSend)
		if class == nil {
			return op, nil, err
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgSend, class.Id)
		if project == nil {
			return op, nil, err
		}

		batch, op, err := getRandomBatchFromProject(ctx, r, qryClient, TypeMsgSend, class.Id)
		if batch == nil {
			return op, nil, err
		}

		admin := sdk.AccAddress(project.Admin).String()
		balres, err := qryClient.Balance(ctx, &core.QueryBalanceRequest{
			Account:    admin,
			BatchDenom: batch.Denom,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, err
		}

		tradableBalance, err := math.NewNonNegativeDecFromString(balres.Balance.TradableAmount)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, err
		}

		retiredBalance, err := math.NewNonNegativeDecFromString(balres.Balance.RetiredAmount)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, err
		}

		if tradableBalance.IsZero() || retiredBalance.IsZero() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, "balance is zero"), nil, nil
		}

		recipient, _ := simtypes.RandomAcc(r, accs)
		if admin == recipient.Address.String() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, "sender & recipient are same"), nil, nil
		}

		addr := sdk.AccAddress(project.Admin)
		acc, found := simtypes.FindAccount(accs, addr)
		if !found {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, "account not found"), nil, nil
		}

		issuer := ak.GetAccount(sdkCtx, acc.Address)
		spendable := bk.SpendableCoins(sdkCtx, issuer.GetAddress())

		var tradable int
		var retired int
		var retirementJurisdiction string
		if !tradableBalance.IsZero() {
			i64, err := tradableBalance.Int64()
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, nil
			}
			if i64 > 1 {
				tradable = simtypes.RandIntBetween(r, 1, int(i64))
				retired = simtypes.RandIntBetween(r, 0, tradable)
				if retired != 0 {
					retirementJurisdiction = "AQ"
				}
			} else {
				tradable = int(i64)
			}
		}

		if retired+tradable > tradable {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, "insufficient credit balance"), nil, nil
		}

		msg := &core.MsgSend{
			Sender:    admin,
			Recipient: recipient.Address.String(),
			Credits: []*core.MsgSend_SendCredits{
				{
					BatchDenom:             batch.Denom,
					TradableAmount:         fmt.Sprintf("%d", tradable),
					RetiredAmount:          fmt.Sprintf("%d", retired),
					RetirementJurisdiction: retirementJurisdiction,
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

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgRetire generates a MsgRetire with random values.
func SimulateMsgRetire(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		ctx := regentypes.Context{Context: sdkCtx}
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgRetire)
		if class == nil {
			return op, nil, err
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgRetire, class.Id)
		if project == nil {
			return op, nil, err
		}

		batch, op, err := getRandomBatchFromProject(ctx, r, qryClient, TypeMsgRetire, project.Id)
		if batch == nil {
			return op, nil, err
		}

		admin := sdk.AccAddress(project.Admin).String()
		balanceRes, err := qryClient.Balance(ctx, &core.QueryBalanceRequest{
			Account:    admin,
			BatchDenom: batch.Denom,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, err.Error()), nil, err
		}

		tradableBalance, err := math.NewNonNegativeDecFromString(balanceRes.Balance.TradableAmount)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, err.Error()), nil, err
		}

		if tradableBalance.IsZero() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, "balance is zero"), nil, nil
		}

		randSub := math.NewDecFromInt64(int64(simtypes.RandIntBetween(r, 1, 10)))
		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, admin, TypeMsgRetire)
		if spendable == nil {
			return op, nil, err
		}

		if !spendable.IsAllPositive() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, "insufficient funds"), nil, nil
		}

		if tradableBalance.Cmp(randSub) != 1 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, "insufficient funds"), nil, nil
		}

		msg := &core.MsgRetire{
			Holder: account.Address.String(),
			Credits: []*core.MsgRetire_RetireCredits{
				{
					BatchDenom: batch.Denom,
					Amount:     randSub.String(),
				},
			},
			Jurisdiction: "ST-UVW XY Z12",
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

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgCancel generates a MsgCancel with random values.
func SimulateMsgCancel(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		ctx := regentypes.Context{Context: sdkCtx}
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgCancel)
		if class == nil {
			return op, nil, err
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgCancel, class.Id)
		if project == nil {
			return op, nil, err
		}

		batch, op, err := getRandomBatchFromProject(ctx, r, qryClient, TypeMsgCancel, project.Id)
		if batch == nil {
			return op, nil, err
		}

		admin := sdk.AccAddress(project.Admin).String()
		balanceRes, err := qryClient.Balance(ctx, &core.QueryBalanceRequest{
			Account:    admin,
			BatchDenom: batch.Denom,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancel, err.Error()), nil, err
		}

		tradableBalance, err := math.NewNonNegativeDecFromString(balanceRes.Balance.TradableAmount)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancel, err.Error()), nil, err
		}

		if tradableBalance.IsZero() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancel, "balance is zero"), nil, nil
		}

		msg := &core.MsgCancel{
			Holder: admin,
			Credits: []*core.MsgCancel_CancelCredits{
				{
					BatchDenom: batch.Denom,
					Amount:     balanceRes.Balance.TradableAmount,
				},
			},
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, admin, TypeMsgCancel)
		if spendable == nil {
			return op, nil, err
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

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgUpdateClassAdmin generates a MsgUpdateClassAdmin with random values
func SimulateMsgUpdateClassAdmin(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgUpdateClassAdmin)
		if class == nil {
			return op, nil, err
		}

		admin := sdk.AccAddress(class.Admin)
		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, admin.String(), TypeMsgUpdateClassAdmin)
		if spendable == nil {
			return op, nil, err
		}

		newAdmin, _ := simtypes.RandomAcc(r, accs)
		if newAdmin.Address.String() == admin.String() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateClassAdmin, "old and new account is same"), nil, nil // skip
		}

		msg := &core.MsgUpdateClassAdmin{
			Admin:    admin.String(),
			ClassId:  class.Id,
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

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgUpdateClassMetadata generates a MsgUpdateClassMetadata with random metadata
func SimulateMsgUpdateClassMetadata(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgUpdateClassMetadata)
		if class == nil {
			return op, nil, err
		}

		admin := sdk.AccAddress(class.Admin)
		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, admin.String(), TypeMsgUpdateClassMetadata)
		if spendable == nil {
			return op, nil, err
		}

		msg := &core.MsgUpdateClassMetadata{
			Admin:    admin.String(),
			ClassId:  class.Id,
			Metadata: simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 10, 256)),
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

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgUpdateClassIssuers generates a MsgUpdateClassMetaData with random values
func SimulateMsgUpdateClassIssuers(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgUpdateClassIssuers)
		if class == nil {
			return op, nil, err
		}

		admin := sdk.AccAddress(class.Admin)
		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, admin.String(), TypeMsgUpdateClassIssuers)
		if spendable == nil {
			return op, nil, err
		}

		issuersRes, err := qryClient.ClassIssuers(sdk.WrapSDKContext(sdkCtx), &core.QueryClassIssuersRequest{ClassId: class.Id})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateClassIssuers, err.Error()), nil, err
		}
		classIssuers := issuersRes.Issuers

		var addIssuers []string
		var removeIssuers []string

		issuers := randomIssuers(r, accs)
		if len(issuers) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateClassIssuers, "empty issuers"), nil, nil
		}

		for _, i := range issuers {
			if utils.Contains(classIssuers, i) {
				removeIssuers = append(removeIssuers, i)
			} else {
				addIssuers = append(addIssuers, i)
			}
		}

		msg := &core.MsgUpdateClassIssuers{
			Admin:         admin.String(),
			ClassId:       class.Id,
			AddIssuers:    addIssuers,
			RemoveIssuers: removeIssuers,
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

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func getClassIssuers(ctx sdk.Context, r *rand.Rand, qryClient core.QueryClient, className string, msgType string) ([]string, simtypes.OperationMsg, error) {
	classIssuers, err := qryClient.ClassIssuers(sdk.WrapSDKContext(ctx), &core.QueryClassIssuersRequest{ClassId: className})
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return []string{}, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, "no credit classes"), nil
		}

		return []string{}, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, err.Error()), err
	}

	return classIssuers.Issuers, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
}

func getRandomProjectFromClass(ctx regentypes.Context, r *rand.Rand, qryClient core.QueryClient, msgType, classID string) (*core.ProjectInfo, simtypes.OperationMsg, error) {
	res, err := qryClient.Projects(ctx, &core.QueryProjectsRequest{
		ClassId: classID,
	})
	if err != nil {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, err.Error()), err
	}

	projects := res.Projects
	if len(projects) == 0 {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, "no projects"), nil
	}

	return projects[r.Intn(len(projects))], simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
}

func getRandomBatchFromProject(ctx regentypes.Context, r *rand.Rand, qryClient core.QueryClient, msgType, projectID string) (*core.BatchInfo, simtypes.OperationMsg, error) {
	res, err := qryClient.BatchesByProject(ctx, &core.QueryBatchesByProjectRequest{
		ProjectId: projectID,
	})
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, fmt.Sprintf("no credit batches for %s project", projectID)), nil
		}
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, err.Error()), err
	}

	batches := res.Batches
	if len(batches) == 0 {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, fmt.Sprintf("no credit batches for %s project", projectID)), nil
	}
	return batches[r.Intn(len(batches))], simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
}

func randomIssuers(r *rand.Rand, accounts []simtypes.Account) []string {
	n := simtypes.RandIntBetween(r, 3, 10)

	var issuers []string
	issuersMap := make(map[string]bool)
	for i := 0; i < n; i++ {
		acc, _ := simtypes.RandomAcc(r, accounts)
		addr := acc.Address.String()
		if _, ok := issuersMap[addr]; ok {
			continue
		}
		issuersMap[acc.Address.String()] = true
		issuers = append(issuers, addr)
	}

	return issuers
}

func generateBatchIssuance(r *rand.Rand, accs []simtypes.Account) []*core.BatchIssuance {
	numIssuances := simtypes.RandIntBetween(r, 3, 10)
	res := make([]*core.BatchIssuance, numIssuances)

	for i := 0; i < numIssuances; i++ {
		recipient := accs[i]
		retiredAmount := simtypes.RandIntBetween(r, 0, 100)
		var retirementJurisdiction string
		if retiredAmount > 0 {
			retirementJurisdiction = "AD"
		}
		res[i] = &core.BatchIssuance{
			Recipient:              recipient.Address.String(),
			TradableAmount:         fmt.Sprintf("%d", simtypes.RandIntBetween(r, 10, 1000)),
			RetiredAmount:          fmt.Sprintf("%d", retiredAmount),
			RetirementJurisdiction: retirementJurisdiction,
		}
	}

	return res
}
