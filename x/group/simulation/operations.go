package simulation

import (
	"math/rand"

	"fmt"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	proto "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	regentypes "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/group"
	"github.com/regen-network/regen-ledger/x/group/exported"
)

// Simulation operation weights constants
const (
	OpMsgCreateGroupRequest               = "op_weight_msg_create_group"
	OpMsgUpdateGroupAdminRequest          = "op_weight_msg_update_group_admin"
	OpMsgUpdateGroupMetadata              = "op_wieght_msg_update_group_metadata"
	OpMsgUpdateGroupMembers               = "op_weight_msg_update_group_members"
	OpMsgCreateGroupAccountRequest        = "op_weight_msg_create_group_account"
	OpMsgUpdateGroupAccountAdmin          = "op_weight_msg_update_group_account_admin"
	OpMsgUpdateGroupAccountDecisionPolicy = "op_weight_msg_update_group_account_decision_policy"
	OpMsgUpdateGroupAccountComment        = "op_weight_msg_update_group_account_comment"
	OpMsgCreateProposal                   = "op_weight_msg_create_proposal"
	OpMsgVote                             = "op_weight_msg_vote"
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONMarshaler, ak exported.AccountKeeper,
	bk exported.BankKeeper, govk exported.GovKeeper, qryClient group.QueryClient) simulation.WeightedOperations {
	var (
		weightMsgCreateGroup                      int
		weightMsgUpdateGroupAdmin                 int
		weightMsgUpdateGroupMetadata              int
		weightMsgUpdateGroupMembers               int
		weightMsgCreateGroupAccount               int
		weightMsgUpdateGroupAccountAdmin          int
		weightMsgUpdateGroupAccountDecisionPolicy int
		weightMsgUpdateGroupAccountComment        int
		weightMsgCreateProposal                   int
		weightMsgVote                             int
	)

	appParams.GetOrGenerate(cdc, OpMsgCreateGroupRequest, &weightMsgCreateGroup, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGroup = simappparams.DefaultWeightMsgCreateValidator
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupAdminRequest, &weightMsgUpdateGroupAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupAdmin = simappparams.DefaultWeightMsgCreateValidator
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupMetadata, &weightMsgUpdateGroupMetadata, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupMetadata = simappparams.DefaultWeightMsgCreateValidator
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgCreateGroupAccountRequest, &weightMsgCreateGroupAccount, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGroupAccount = simappparams.DefaultWeightMsgCreateValidator
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgCreateProposal, &weightMsgCreateProposal, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProposal = simappparams.DefaultWeightMsgCreateValidator
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupMembers, &weightMsgUpdateGroupMembers, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupMembers = simappparams.DefaultWeightMsgCreateValidator
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupAccountAdmin, &weightMsgUpdateGroupAccountAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupAccountAdmin = simappparams.DefaultWeightMsgCreateValidator
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupAccountDecisionPolicy, &weightMsgUpdateGroupAccountDecisionPolicy, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupAccountDecisionPolicy = simappparams.DefaultWeightMsgCreateValidator
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupAccountComment, &weightMsgUpdateGroupAccountComment, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupAccountComment = simappparams.DefaultWeightMsgCreateValidator
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgVote, &weightMsgVote, nil,
		func(_ *rand.Rand) {
			weightMsgVote = simappparams.DefaultWeightMsgCreateValidator
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgCreateGroup,
			SimulateMsgCreateGroup(ak, bk),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupAdmin,
			SimulateMsgUpdateGroupAdmin(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupMetadata,
			SimulateMsgUpdateGroupMetadata(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCreateGroupAccount,
			SimulateMsgCreateGroupAccount(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCreateProposal,
			SimulateMsgCreateProposal(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupAdmin,
			SimulateMsgUpdateGroupAdmin(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupMembers,
			SimulateMsgUpdateGroupMembers(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupAccountAdmin,
			SimulateMsgUpdateGroupAccountAdmin(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupAccountDecisionPolicy,
			SimulateMsgUpdateGroupAccountDecisionPolicy(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupAccountComment,
			SimulateMsgUpdateGroupAccountComment(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgVote,
			SimulateMsgVote(ak, bk, govk, qryClient),
		),
	}
}

func SimulateMsgCreateGroup(ak exported.AccountKeeper, bk exported.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc := accounts[0]

		account := ak.GetAccount(ctx, acc.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateGroup, "fee error"), nil, err
		}

		members := []group.Member{
			{
				Address:  acc.Address.String(),
				Weight:   fmt.Sprintf("%d", simappparams.DefaultWeightMsgCreateValidator),
				Metadata: []byte(simtypes.RandStringOfLength(r, 10)),
			},
		}

		msg := &group.MsgCreateGroupRequest{Admin: acc.Address.String(), Members: members, Metadata: []byte(simtypes.RandStringOfLength(r, 10))}

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
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateGroup, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, err
	}
}

func SimulateMsgCreateGroupAccount(ak exported.AccountKeeper, bk exported.BankKeeper, qryClient group.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc := accounts[0]

		account := ak.GetAccount(ctx, acc.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateGroupAccount, "fee error"), nil, err
		}

		ctx1 := regentypes.Context{Context: ctx}
		result, err := qryClient.GroupAccountInfo(ctx1, &group.QueryGroupAccountInfoRequest{GroupAccount: acc.Address.String()})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateGroupAccount, "fail to query group info"), nil, err
		}

		if result.Info == nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateGroupAccount, "no group account found"), nil, nil
		}

		addr, err := sdk.AccAddressFromBech32(result.Info.Admin)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateGroupAccount, "fail to decode acc address"), nil, err
		}

		msg, err := group.NewMsgCreateGroupAccountRequest(
			addr,
			uint64(simtypes.RandIntBetween(r, 1, 10)),
			[]byte(simtypes.RandStringOfLength(r, 10)),
			&group.ThresholdDecisionPolicy{
				Threshold: fmt.Sprintf("%d", simtypes.RandIntBetween(r, 1, 100)),
				Timeout:   proto.Duration{Seconds: int64(simtypes.RandIntBetween(r, 100, 1000))},
			},
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateGroupAccount, err.Error()), nil, err
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
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateGroupAccount, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, err
	}
}

func SimulateMsgCreateProposal(ak exported.AccountKeeper, bk exported.BankKeeper, queryClient group.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc := accounts[0]

		account := ak.GetAccount(ctx, acc.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateProposal, "fee error"), nil, err
		}

		ctx1 := regentypes.Context{Context: ctx}
		result, err := queryClient.GroupAccountInfo(ctx1, &group.QueryGroupAccountInfoRequest{GroupAccount: acc.Address.String()})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateProposal, "fail to query group info"), nil, err
		}

		if result.Info == nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateProposal, "no group account found"), nil, nil
		}

		msg := group.MsgCreateProposalRequest{
			GroupAccount: result.Info.GroupAccount,
			Proposers:    []string{acc.Address.String()},
			Metadata:     []byte(simtypes.RandStringOfLength(r, 10)),
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
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateProposal, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(&msg, true, "", nil), nil, err
	}
}

func SimulateMsgUpdateGroupAdmin(ak exported.AccountKeeper, bk exported.BankKeeper, queryClient group.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc1 := accounts[0]
		acc2 := accounts[1]

		account2 := ak.GetAccount(ctx, acc2.Address)

		spendableCoins := bk.SpendableCoins(ctx, account2.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAdmin, "fee error"), nil, err
		}

		ctx1 := regentypes.Context{Context: ctx}
		result, err := queryClient.GroupAccountInfo(ctx1, &group.QueryGroupAccountInfoRequest{GroupAccount: acc1.Address.String()})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAdmin, "fail to query group info"), nil, err
		}

		if result.Info == nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAdmin, "no group account found"), nil, nil
		}

		msg := group.MsgUpdateGroupAccountAdminRequest{
			GroupAccount: result.Info.GroupAccount,
			Admin:        result.Info.Admin,
			NewAdmin:     account2.GetAddress().String(),
		}
		txGen := simappparams.MakeTestEncodingConfig().TxConfig

		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account2.GetAccountNumber()},
			[]uint64{account2.GetSequence()},
			acc2.PrivKey,
		)

		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAdmin, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(&msg, true, "", nil), nil, err
	}
}

func SimulateMsgUpdateGroupMetadata(ak exported.AccountKeeper, bk exported.BankKeeper, queryClient group.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc := accounts[0]

		account := ak.GetAccount(ctx, acc.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupComment, "fee error"), nil, err
		}

		ctx1 := regentypes.Context{Context: ctx}
		result, err := queryClient.GroupAccountInfo(ctx1, &group.QueryGroupAccountInfoRequest{GroupAccount: acc.Address.String()})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupComment, "fail to query group info"), nil, err
		}

		if result.Info == nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupComment, "no group account found"), nil, nil
		}

		msg := group.MsgUpdateGroupMetadataRequest{
			GroupId:  result.Info.GroupId,
			Admin:    result.Info.Admin,
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
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupComment, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(&msg, true, "", nil), nil, err
	}
}

func SimulateMsgUpdateGroupMembers(ak exported.AccountKeeper,
	bk exported.BankKeeper, queryClient group.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc1 := accounts[0]
		acc2 := accounts[0]
		acc3 := accounts[0]

		account := ak.GetAccount(ctx, acc1.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupMembers, "fee error"), nil, err
		}

		ctx1 := regentypes.Context{Context: ctx}
		result, err := queryClient.GroupAccountInfo(ctx1, &group.QueryGroupAccountInfoRequest{GroupAccount: acc1.Address.String()})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupMembers, "fail to query group info"), nil, err
		}

		if result.Info == nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupMembers, "no group account found"), nil, nil
		}

		members := []group.Member{
			{
				Address:  acc2.Address.String(),
				Weight:   fmt.Sprintf("%d", simappparams.DefaultWeightMsgCreateValidator),
				Metadata: []byte(simtypes.RandStringOfLength(r, 10)),
			},
			{
				Address:  acc3.Address.String(),
				Weight:   fmt.Sprintf("%d", simappparams.DefaultWeightCommunitySpendProposal),
				Metadata: []byte(simtypes.RandStringOfLength(r, 10)),
			},
		}

		msg := group.MsgUpdateGroupMembersRequest{
			GroupId:       result.Info.GroupId,
			Admin:         result.Info.Admin,
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
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupMembers, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(&msg, true, "", nil), nil, err
	}
}

func SimulateMsgUpdateGroupAccountAdmin(ak exported.AccountKeeper,
	bk exported.BankKeeper, queryClient group.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc1 := accounts[0]
		acc2 := accounts[0]

		account := ak.GetAccount(ctx, acc1.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountAdmin, "fee error"), nil, err
		}

		ctx1 := regentypes.Context{Context: ctx}
		result, err := queryClient.GroupAccountInfo(ctx1, &group.QueryGroupAccountInfoRequest{GroupAccount: acc1.Address.String()})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountAdmin, "fail to query group info"), nil, err
		}

		if result.Info == nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountAdmin, "no group account found"), nil, nil
		}

		msg := group.MsgUpdateGroupAccountAdminRequest{
			Admin:        result.Info.Admin,
			GroupAccount: result.Info.GroupAccount,
			NewAdmin:     acc2.Address.String(),
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
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountAdmin, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(&msg, true, "", nil), nil, err
	}
}

func SimulateMsgUpdateGroupAccountDecisionPolicy(ak exported.AccountKeeper,
	bk exported.BankKeeper, queryClient group.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc1 := accounts[0]

		account := ak.GetAccount(ctx, acc1.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountDecisionPolicy, "fee error"), nil, err
		}

		ctx1 := regentypes.Context{Context: ctx}
		result, err := queryClient.GroupAccountInfo(ctx1, &group.QueryGroupAccountInfoRequest{GroupAccount: acc1.Address.String()})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountDecisionPolicy, "fail to query group info"), nil, err
		}

		if result.Info == nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountDecisionPolicy, "no group account found"), nil, nil
		}

		admin, err := sdk.AccAddressFromBech32(result.Info.Admin)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountDecisionPolicy, fmt.Sprintf("fail to decide bech32 address: %s", err.Error())), nil, nil
		}

		groupAccount, err := sdk.AccAddressFromBech32(result.Info.GroupAccount)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountDecisionPolicy, fmt.Sprintf("fail to decide bech32 address: %s", err.Error())), nil, nil
		}

		msg, err := group.NewMsgUpdateGroupAccountDecisionPolicyRequest(admin, groupAccount, &group.ThresholdDecisionPolicy{
			Threshold: fmt.Sprintf("%d", simtypes.RandIntBetween(r, 1, 100)),
			Timeout:   proto.Duration{Seconds: int64(simtypes.RandIntBetween(r, 100, 1000))},
		})

		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountDecisionPolicy, err.Error()), nil, err
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
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountDecisionPolicy, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, err
	}
}

func SimulateMsgUpdateGroupAccountComment(ak exported.AccountKeeper,
	bk exported.BankKeeper, queryClient group.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc1 := accounts[0]

		account := ak.GetAccount(ctx, acc1.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountComment, "fee error"), nil, err
		}

		ctx1 := regentypes.Context{Context: ctx}
		result, err := queryClient.GroupAccountInfo(ctx1, &group.QueryGroupAccountInfoRequest{GroupAccount: acc1.Address.String()})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountComment, "fail to query group info"), nil, err
		}

		if result.Info == nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountComment, "no group account found"), nil, nil
		}

		msg := group.MsgUpdateGroupAccountMetadataRequest{
			Admin:        result.Info.Admin,
			GroupAccount: result.Info.GroupAccount,
			Metadata:     []byte(simtypes.RandStringOfLength(r, 10)),
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
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountComment, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(&msg, true, "", nil), nil, err
	}
}

func SimulateMsgVote(ak exported.AccountKeeper,
	bk exported.BankKeeper, govk exported.GovKeeper, queryClient group.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc1 := accounts[0]
		acc2 := accounts[1]

		account := ak.GetAccount(ctx, acc1.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgVote, "fee error"), nil, err
		}

		ctx1 := regentypes.Context{Context: ctx}
		accountInfo, err := queryClient.GroupAccountInfo(ctx1, &group.QueryGroupAccountInfoRequest{GroupAccount: acc1.Address.String()})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgVote, "fail to query group info"), nil, err
		}

		if accountInfo.Info == nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgVote, "no group account found"), nil, nil
		}

		testProposal := govtypes.NewTextProposal("Test", "description")

		proposal, err := govk.SubmitProposal(ctx, testProposal)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgVote, "fail to create proposal"), nil, nil
		}

		msg := group.MsgVoteRequest{
			ProposalId: proposal.ProposalId,
			Voter:      acc2.Address.String(),
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
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgVote, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(&msg, true, "", nil), nil, err
	}
}
