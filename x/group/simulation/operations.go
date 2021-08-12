package simulation

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	gogotypes "github.com/gogo/protobuf/types"

	regentypes "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/group"
	"github.com/regen-network/regen-ledger/x/group/exported"
)

// group message types
var (
	TypeMsgCreateGroup                      = sdk.MsgTypeURL(&group.MsgCreateGroup{})
	TypeMsgUpdateGroupMembers               = sdk.MsgTypeURL(&group.MsgUpdateGroupMembers{})
	TypeMsgUpdateGroupAdmin                 = sdk.MsgTypeURL(&group.MsgUpdateGroupAdmin{})
	TypeMsgUpdateGroupMetadata              = sdk.MsgTypeURL(&group.MsgUpdateGroupMetadata{})
	TypeMsgCreateGroupAccount               = sdk.MsgTypeURL(&group.MsgCreateGroupAccount{})
	TypeMsgUpdateGroupAccountAdmin          = sdk.MsgTypeURL(&group.MsgUpdateGroupAccountAdmin{})
	TypeMsgUpdateGroupAccountDecisionPolicy = sdk.MsgTypeURL(&group.MsgUpdateGroupAccountDecisionPolicy{})
	TypeMsgUpdateGroupAccountMetadata       = sdk.MsgTypeURL(&group.MsgUpdateGroupAccountMetadata{})
	TypeMsgCreateProposal                   = sdk.MsgTypeURL(&group.MsgCreateProposal{})
	TypeMsgVote                             = sdk.MsgTypeURL(&group.MsgVote{})
	TypeMsgExec                             = sdk.MsgTypeURL(&group.MsgExec{})
)

// Simulation operation weights constants
const (
	OpMsgCreateGroup                      = "op_weight_msg_create_group"
	OpMsgUpdateGroupAdmin                 = "op_weight_msg_update_group_admin"
	OpMsgUpdateGroupMetadata              = "op_wieght_msg_update_group_metadata"
	OpMsgUpdateGroupMembers               = "op_weight_msg_update_group_members"
	OpMsgCreateGroupAccountRequest        = "op_weight_msg_create_group_account"
	OpMsgUpdateGroupAccountAdmin          = "op_weight_msg_update_group_account_admin"
	OpMsgUpdateGroupAccountDecisionPolicy = "op_weight_msg_update_group_account_decision_policy"
	OpMsgUpdateGroupAccountMetaData       = "op_weight_msg_update_group_account_metadata"
	OpMsgCreateProposal                   = "op_weight_msg_create_proposal"
	OpMsgVote                             = "op_weight_msg_vote"
	OpMsgExec                             = "ops_weight_msg_exec"
)

//  If update group or group account txn's executed, `SimulateMsgVote` & `SimulateMsgExec` txn's returns `noOp`.
//  That's why we have less weight for update group & group-account txn's.
const (
	WeightCreateGroup                      = 100
	WeightCreateGroupAccount               = 100
	WeightCreateProposal                   = 90
	WeightMsgVote                          = 90
	WeightMsgExec                          = 90
	WeightUpdateGroupMetadata              = 5
	WeightUpdateGroupAdmin                 = 5
	WeightUpdateGroupMembers               = 5
	WeightUpdateGroupAccountAdmin          = 5
	WeightUpdateGroupAccountDecisionPolicy = 5
	WeightUpdateGroupAccountMetadata       = 5
	GroupMemberWeight                      = 40
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec, ak exported.AccountKeeper,
	bk exported.BankKeeper, qryClient group.QueryClient, protoCdc *codec.ProtoCodec) simulation.WeightedOperations {
	var (
		weightMsgCreateGroup                      int
		weightMsgUpdateGroupAdmin                 int
		weightMsgUpdateGroupMetadata              int
		weightMsgUpdateGroupMembers               int
		weightMsgCreateGroupAccount               int
		weightMsgUpdateGroupAccountAdmin          int
		weightMsgUpdateGroupAccountDecisionPolicy int
		weightMsgUpdateGroupAccountMetadata       int
		weightMsgCreateProposal                   int
		weightMsgVote                             int
		weightMsgExec                             int
	)

	appParams.GetOrGenerate(cdc, OpMsgCreateGroup, &weightMsgCreateGroup, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGroup = WeightCreateGroup
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgCreateGroupAccountRequest, &weightMsgCreateGroupAccount, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGroupAccount = WeightCreateGroupAccount
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgCreateProposal, &weightMsgCreateProposal, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProposal = WeightCreateProposal
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgVote, &weightMsgVote, nil,
		func(_ *rand.Rand) {
			weightMsgVote = WeightMsgVote
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgExec, &weightMsgExec, nil,
		func(_ *rand.Rand) {
			weightMsgExec = WeightMsgExec
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupMetadata, &weightMsgUpdateGroupMetadata, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupMetadata = WeightUpdateGroupMetadata
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupAdmin, &weightMsgUpdateGroupAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupAdmin = WeightUpdateGroupAdmin
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupMembers, &weightMsgUpdateGroupMembers, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupMembers = WeightUpdateGroupMembers
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupAccountAdmin, &weightMsgUpdateGroupAccountAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupAccountAdmin = WeightUpdateGroupAccountAdmin
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupAccountDecisionPolicy, &weightMsgUpdateGroupAccountDecisionPolicy, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupAccountDecisionPolicy = WeightUpdateGroupAccountDecisionPolicy
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupAccountMetaData, &weightMsgUpdateGroupAccountMetadata, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupAccountMetadata = WeightUpdateGroupAccountMetadata
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgCreateGroup,
			SimulateMsgCreateGroup(ak, bk, protoCdc),
		),
		simulation.NewWeightedOperation(
			weightMsgCreateGroupAccount,
			SimulateMsgCreateGroupAccount(ak, bk, qryClient, protoCdc),
		),
		simulation.NewWeightedOperation(
			weightMsgCreateProposal,
			SimulateMsgCreateProposal(ak, bk, qryClient, protoCdc),
		),
		simulation.NewWeightedOperation(
			weightMsgVote,
			SimulateMsgVote(ak, bk, qryClient, protoCdc),
		),
		simulation.NewWeightedOperation(
			weightMsgExec,
			SimulateMsgExec(ak, bk, qryClient, protoCdc),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupMetadata,
			SimulateMsgUpdateGroupMetadata(ak, bk, qryClient, protoCdc),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupAdmin,
			SimulateMsgUpdateGroupAdmin(ak, bk, qryClient, protoCdc),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupMembers,
			SimulateMsgUpdateGroupMembers(ak, bk, qryClient, protoCdc),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupAccountAdmin,
			SimulateMsgUpdateGroupAccountAdmin(ak, bk, qryClient, protoCdc),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupAccountDecisionPolicy,
			SimulateMsgUpdateGroupAccountDecisionPolicy(ak, bk, qryClient, protoCdc),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupAccountMetadata,
			SimulateMsgUpdateGroupAccountMetadata(ak, bk, qryClient, protoCdc),
		),
	}
}

// SimulateMsgCreateGroup generates a MsgCreateGroup with random values
func SimulateMsgCreateGroup(ak exported.AccountKeeper, bk exported.BankKeeper, protoCdc *codec.ProtoCodec) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc, _ := simtypes.RandomAcc(r, accounts)
		account := ak.GetAccount(ctx, acc.Address)
		accAddr := acc.Address.String()

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroup, "fee error"), nil, err
		}

		members := []group.Member{
			{
				Address:  accAddr,
				Weight:   fmt.Sprintf("%d", GroupMemberWeight),
				Metadata: []byte(simtypes.RandStringOfLength(r, 10)),
			},
		}

		msg := &group.MsgCreateGroup{Admin: accAddr, Members: members, Metadata: []byte(simtypes.RandStringOfLength(r, 10))}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroup, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "", protoCdc), nil, err
	}
}

// SimulateMsgCreateGroupAccount generates a NewMsgCreateGroupAccount with random values
func SimulateMsgCreateGroupAccount(ak exported.AccountKeeper, bk exported.BankKeeper, qryClient group.QueryClient, protoCdc *codec.ProtoCodec) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc := accounts[0]
		account := ak.GetAccount(ctx, acc.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupAccount, "fee error"), nil, err
		}

		groupAdmin, groupID, op, err := getGroupDetails(ctx, qryClient, acc)
		if err != nil {
			return op, nil, err
		}
		if groupAdmin == "" {
			return op, nil, nil
		}

		addr, err := sdk.AccAddressFromBech32(groupAdmin)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupAccount, "fail to decode acc address"), nil, err
		}

		msg, err := group.NewMsgCreateGroupAccount(
			addr,
			groupID,
			[]byte(simtypes.RandStringOfLength(r, 10)),
			&group.ThresholdDecisionPolicy{
				Threshold: "20",
				Timeout:   gogotypes.Duration{Seconds: int64(30 * 24 * 60 * 60)},
			},
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupAccount, err.Error()), nil, err
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupAccount, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(msg, true, "", protoCdc), nil, err
	}
}

func getGroupDetails(sdkCtx sdk.Context, qryClient group.QueryClient, acc simtypes.Account) (groupAdmin string, groupID uint64, op simtypes.OperationMsg, err error) {
	ctx := regentypes.Context{Context: sdkCtx}

	groups, err := qryClient.GroupsByAdmin(ctx, &group.QueryGroupsByAdminRequest{Admin: acc.Address.String()})
	if err != nil {
		return "", 0, simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupAccount, "fail to query groups"), err
	}

	if len(groups.Groups) == 0 {
		return "", 0, simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupAccount, ""), nil
	}

	groupAdmin = groups.Groups[0].Admin
	groupID = groups.Groups[0].GroupId

	return groupAdmin, groupID, simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupAccount, ""), nil
}

// SimulateMsgCreateProposal generates a NewMsgCreateProposal with random values
func SimulateMsgCreateProposal(ak exported.AccountKeeper, bk exported.BankKeeper, queryClient group.QueryClient, protoCdc *codec.ProtoCodec) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc := accounts[0]
		account := ak.GetAccount(sdkCtx, acc.Address)

		spendableCoins := bk.SpendableCoins(sdkCtx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, sdkCtx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateProposal, "fee error"), nil, err
		}

		ctx := regentypes.Context{Context: sdkCtx}
		groupAdmin, _, op, err := getGroupDetails(sdkCtx, queryClient, acc)
		if err != nil {
			return op, nil, err
		}
		if groupAdmin == "" {
			return op, nil, nil
		}

		groupAccounts, opMsg, err := groupAccountsByAdmin(ctx, queryClient, groupAdmin)
		if groupAccounts == nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateProposal, opMsg), nil, err
		}

		msg := group.MsgCreateProposal{
			Address:   groupAccounts[0].Address,
			Proposers: []string{acc.Address.String()},
			Metadata:  []byte(simtypes.RandStringOfLength(r, 10)),
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateProposal, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(&msg, true, "", protoCdc), nil, err
	}
}

// SimulateMsgUpdateGroupAdmin generates a MsgUpdateGroupAccountAdmin with random values
func SimulateMsgUpdateGroupAdmin(ak exported.AccountKeeper, bk exported.BankKeeper, queryClient group.QueryClient, protoCdc *codec.ProtoCodec) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc1 := accounts[0]
		acc2 := accounts[1]

		account := ak.GetAccount(sdkCtx, acc1.Address)

		spendableCoins := bk.SpendableCoins(sdkCtx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, sdkCtx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupAdmin, "fee error"), nil, err
		}

		ctx := regentypes.Context{Context: sdkCtx}

		groupAdmin, groupID, op, err := getGroupDetails(sdkCtx, queryClient, acc1)
		if err != nil {
			return op, nil, err
		}
		if groupAdmin == "" {
			return op, nil, nil
		}

		groupAccounts, opMsg, err := groupAccountsByAdmin(ctx, queryClient, groupAdmin)
		if groupAccounts == nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupAdmin, opMsg), nil, err
		}

		msg := group.MsgUpdateGroupAdmin{
			GroupId:  groupID,
			Admin:    groupAccounts[0].Admin,
			NewAdmin: acc2.Address.String(),
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc1.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupAdmin, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(&msg, true, "", protoCdc), nil, err
	}
}

// SimulateMsgUpdateGroupMetadata generates a MsgUpdateGroupMetadata with random values
func SimulateMsgUpdateGroupMetadata(ak exported.AccountKeeper, bk exported.BankKeeper, queryClient group.QueryClient, protoCdc *codec.ProtoCodec) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc := accounts[0]
		account := ak.GetAccount(sdkCtx, acc.Address)

		spendableCoins := bk.SpendableCoins(sdkCtx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, sdkCtx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupMetadata, "fee error"), nil, err
		}

		groupAdmin, groupID, op, err := getGroupDetails(sdkCtx, queryClient, acc)
		if err != nil {
			return op, nil, err
		}
		if groupAdmin == "" {
			return op, nil, nil
		}

		msg := group.MsgUpdateGroupMetadata{
			GroupId:  groupID,
			Admin:    groupAdmin,
			Metadata: []byte(simtypes.RandStringOfLength(r, 10)),
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupMetadata, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(&msg, true, "", protoCdc), nil, err
	}
}

// SimulateMsgUpdateGroupMembers generates a MsgUpdateGroupMembers with random values
func SimulateMsgUpdateGroupMembers(ak exported.AccountKeeper,
	bk exported.BankKeeper, queryClient group.QueryClient, protoCdc *codec.ProtoCodec) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc1 := accounts[0]
		acc2 := accounts[1]
		acc3 := accounts[2]
		account := ak.GetAccount(sdkCtx, acc1.Address)

		spendableCoins := bk.SpendableCoins(sdkCtx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, sdkCtx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupMembers, "fee error"), nil, err
		}

		ctx := regentypes.Context{Context: sdkCtx}
		groupAdmin, groupID, op, err := getGroupDetails(sdkCtx, queryClient, acc1)
		if err != nil {
			return op, nil, err
		}
		if groupAdmin == "" {
			return op, nil, nil
		}

		groupAccounts, opMsg, err := groupAccountsByAdmin(ctx, queryClient, groupAdmin)
		if groupAccounts == nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupMembers, opMsg), nil, err
		}

		members := []group.Member{
			{
				Address:  acc2.Address.String(),
				Weight:   fmt.Sprintf("%d", GroupMemberWeight),
				Metadata: []byte(simtypes.RandStringOfLength(r, 10)),
			},
			{
				Address:  acc3.Address.String(),
				Weight:   fmt.Sprintf("%d", GroupMemberWeight),
				Metadata: []byte(simtypes.RandStringOfLength(r, 10)),
			},
		}

		msg := group.MsgUpdateGroupMembers{
			GroupId:       groupID,
			Admin:         groupAdmin,
			MemberUpdates: members,
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc1.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupMembers, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(&msg, true, "", protoCdc), nil, err
	}
}

// SimulateMsgUpdateGroupAccountAdmin generates a MsgUpdateGroupAccountAdmin with random values
func SimulateMsgUpdateGroupAccountAdmin(ak exported.AccountKeeper, bk exported.BankKeeper, queryClient group.QueryClient, protoCdc *codec.ProtoCodec) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc1 := accounts[0]
		acc2 := accounts[1]

		account := ak.GetAccount(sdkCtx, acc1.Address)

		spendableCoins := bk.SpendableCoins(sdkCtx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, sdkCtx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupAccountAdmin, "fee error"), nil, err
		}

		ctx := regentypes.Context{Context: sdkCtx}

		groupAdmin, _, op, err := getGroupDetails(sdkCtx, queryClient, acc1)
		if err != nil {
			return op, nil, err
		}
		if groupAdmin == "" {
			return op, nil, nil
		}

		groupAccounts, opMsg, err := groupAccountsByAdmin(ctx, queryClient, groupAdmin)
		if groupAccounts == nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupAccountAdmin, opMsg), nil, err
		}

		msg := group.MsgUpdateGroupAccountAdmin{
			Admin:    groupAccounts[0].Admin,
			Address:  groupAccounts[0].Address,
			NewAdmin: acc2.Address.String(),
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc1.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupAccountAdmin, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(&msg, true, "", protoCdc), nil, err
	}
}

// SimulateMsgUpdateGroupAccountDecisionPolicy generates a NewMsgUpdateGroupAccountDecisionPolicyRequest with random values
func SimulateMsgUpdateGroupAccountDecisionPolicy(ak exported.AccountKeeper,
	bk exported.BankKeeper, queryClient group.QueryClient, protoCdc *codec.ProtoCodec) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc1 := accounts[0]
		account := ak.GetAccount(sdkCtx, acc1.Address)

		spendableCoins := bk.SpendableCoins(sdkCtx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, sdkCtx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupAccountDecisionPolicy, "fee error"), nil, err
		}

		ctx := regentypes.Context{Context: sdkCtx}
		groupAdmin, _, op, err := getGroupDetails(sdkCtx, queryClient, acc1)
		if err != nil {
			return op, nil, err
		}
		if groupAdmin == "" {
			return op, nil, nil
		}

		groupAccounts, opMsg, err := groupAccountsByAdmin(ctx, queryClient, groupAdmin)
		if groupAccounts == nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupAccountDecisionPolicy, opMsg), nil, err
		}

		adminBech32, err := sdk.AccAddressFromBech32(groupAccounts[0].Admin)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupAccountDecisionPolicy, fmt.Sprintf("fail to decide bech32 address: %s", err.Error())), nil, nil
		}

		groupAccountBech32, err := sdk.AccAddressFromBech32(groupAccounts[0].Address)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupAccountDecisionPolicy, fmt.Sprintf("fail to decide bech32 address: %s", err.Error())), nil, nil
		}

		msg, err := group.NewMsgUpdateGroupAccountDecisionPolicyRequest(adminBech32, groupAccountBech32, &group.ThresholdDecisionPolicy{
			Threshold: fmt.Sprintf("%d", simtypes.RandIntBetween(r, 1, 20)),
			Timeout:   gogotypes.Duration{Seconds: int64(simtypes.RandIntBetween(r, 100, 1000))},
		})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupAccountDecisionPolicy, err.Error()), nil, err
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc1.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupAccountDecisionPolicy, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "", protoCdc), nil, err
	}
}

// SimulateMsgUpdateGroupAccountMetadata generates a MsgUpdateGroupAccountMetadata with random values
func SimulateMsgUpdateGroupAccountMetadata(ak exported.AccountKeeper,
	bk exported.BankKeeper, queryClient group.QueryClient, protoCdc *codec.ProtoCodec) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc1 := accounts[0]

		account := ak.GetAccount(sdkCtx, acc1.Address)

		spendableCoins := bk.SpendableCoins(sdkCtx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, sdkCtx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupAccountMetadata, "fee error"), nil, err
		}

		ctx := regentypes.Context{Context: sdkCtx}
		groupAdmin, _, op, err := getGroupDetails(sdkCtx, queryClient, acc1)
		if err != nil {
			return op, nil, err
		}
		if groupAdmin == "" {
			return op, nil, nil
		}

		groupAccounts, opMsg, err := groupAccountsByAdmin(ctx, queryClient, groupAdmin)
		if groupAccounts == nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupAccountMetadata, opMsg), nil, err
		}

		msg := group.MsgUpdateGroupAccountMetadata{
			Admin:    groupAccounts[0].Admin,
			Address:  groupAccounts[0].Address,
			Metadata: []byte(simtypes.RandStringOfLength(r, 10)),
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc1.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupAccountMetadata, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(&msg, true, "", protoCdc), nil, err
	}
}

// SimulateMsgVote generates a MsgVote with random values
func SimulateMsgVote(ak exported.AccountKeeper,
	bk exported.BankKeeper, queryClient group.QueryClient, protoCdc *codec.ProtoCodec) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc1 := accounts[0]

		account := ak.GetAccount(sdkCtx, acc1.Address)

		spendableCoins := bk.SpendableCoins(sdkCtx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, sdkCtx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "fee error"), nil, err
		}
		ctx := regentypes.Context{Context: sdkCtx}
		groupAdmin, _, op, err := getGroupDetails(sdkCtx, queryClient, acc1)
		if err != nil {
			return op, nil, err
		}
		if groupAdmin == "" {
			return op, nil, nil
		}

		groupAccounts, opMsg, err := groupAccountsByAdmin(ctx, queryClient, groupAdmin)
		if groupAccounts == nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, opMsg), nil, err
		}

		proposalsResult, err := queryClient.ProposalsByGroupAccount(ctx, &group.QueryProposalsByGroupAccountRequest{Address: groupAccounts[0].Address})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "fail to query group info"), nil, err
		}
		proposals := proposalsResult.GetProposals()
		if len(proposals) == 0 {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "no proposals found"), nil, nil
		}

		proposalID := -1

		for _, proposal := range proposals {
			if proposal.Status == group.ProposalStatusSubmitted {
				votingPeriodEnd, err := gogotypes.TimestampFromProto(&proposal.Timeout)
				if err != nil {
					return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "error: while converting to timestamp"), nil, err
				}
				proposalID = int(proposal.ProposalId)
				if votingPeriodEnd.Before(ctx.BlockTime()) || votingPeriodEnd.Equal(ctx.BlockTime()) {
					return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "voting period ended: skipping"), nil, nil
				}
				break
			}
		}

		// return no-op if no proposal found
		if proposalID == -1 {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "no proposals found"), nil, nil
		}

		msg := group.MsgVote{
			ProposalId: uint64(proposalID),
			Voter:      acc1.Address.String(),
			Choice:     group.Choice_CHOICE_YES,
			Metadata:   []byte(simtypes.RandStringOfLength(r, 10)),
		}
		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc1.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupAccountMetadata, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)

		if err != nil {
			if strings.Contains(err.Error(), "group was modified") || strings.Contains(err.Error(), "group account was modified") {
				return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "no-op:group/group-account was modified"), nil, nil
			}
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(&msg, true, "", protoCdc), nil, err
	}
}

// SimulateMsgExec generates a MsgExec with random values
func SimulateMsgExec(ak exported.AccountKeeper,
	bk exported.BankKeeper, queryClient group.QueryClient, protoCdc *codec.ProtoCodec) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc1 := accounts[0]

		account := ak.GetAccount(sdkCtx, acc1.Address)

		spendableCoins := bk.SpendableCoins(sdkCtx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, sdkCtx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "fee error"), nil, err
		}
		ctx := regentypes.Context{Context: sdkCtx}
		groupAdmin, _, op, err := getGroupDetails(sdkCtx, queryClient, acc1)
		if err != nil {
			return op, nil, err
		}
		if groupAdmin == "" {
			return op, nil, nil
		}

		groupAccounts, opMsg, err := groupAccountsByAdmin(ctx, queryClient, groupAdmin)
		if groupAccounts == nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, opMsg), nil, err
		}

		proposalsResult, err := queryClient.ProposalsByGroupAccount(ctx, &group.QueryProposalsByGroupAccountRequest{Address: groupAccounts[0].Address})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "fail to query group info"), nil, err
		}
		proposals := proposalsResult.GetProposals()
		if len(proposals) == 0 {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "no proposals found"), nil, nil
		}

		proposalID := -1

		for _, proposal := range proposals {
			if proposal.Status == group.ProposalStatusSubmitted {
				proposalID = int(proposal.ProposalId)
				break
			}
		}

		// return no-op if no proposal found
		if proposalID == -1 {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "no proposals found"), nil, nil
		}

		msg := group.MsgExec{
			ProposalId: uint64(proposalID),
			Signer:     acc1.Address.String(),
		}
		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc1.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupAccountMetadata, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			if strings.Contains(err.Error(), "group was modified") || strings.Contains(err.Error(), "group account was modified") {
				return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "no-op:group/group-account was modified"), nil, nil
			}
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(&msg, true, "", protoCdc), nil, err
	}
}

func groupAccountsByAdmin(ctx regentypes.Context, qryClient group.QueryClient, admin string) ([]*group.GroupAccountInfo, string, error) {
	result, err := qryClient.GroupAccountsByAdmin(ctx, &group.QueryGroupAccountsByAdminRequest{Admin: admin})
	if err != nil {
		return nil, "fail to query group info", err
	}

	groupAccounts := result.GetGroupAccounts()
	if len(groupAccounts) == 0 {
		return nil, "no group account found", nil
	}
	return groupAccounts, "", nil
}
