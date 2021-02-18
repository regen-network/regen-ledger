package simulation

import (
	"math/rand"

	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	proto "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	regentypes "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/group"
	"github.com/regen-network/regen-ledger/x/group/server"
)

// Simulation operation weights constants
const (
	OpMsgCreateGroupRequest        = "op_weight_msg_create_group"
	OpMsgCreateGroupAccountRequest = "op_weight_msg_create_group_account"
	OpMsgCreateProposal            = "op_weight_msg_create_proposal"
)

// group message types
const (
	TypeMsgCreateGroup        = "/regen.group.v1alpha1.Msg/CreateGroup"
	TypeMsgCreateGroupAccount = "/regen.group.v1alpha1.Msg/CreateGroupAccount"
	TypeMsgCreateProposal     = "/regen.group.v1alpha1.Msg/CreateProposal"
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONMarshaler, ak server.AccountKeeper,
	bk server.BankKeeper, protoCdc *codec.ProtoCodec, qryClient group.QueryClient) simulation.WeightedOperations {
	var (
		weightMsgCreateGroup        int
		weightMsgCreateGroupAccount int
		weightMsgCreateProposal     int
	)

	appParams.GetOrGenerate(cdc, OpMsgCreateGroupRequest, &weightMsgCreateGroup, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGroup = simappparams.DefaultWeightMsgCreateValidator
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

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgCreateGroup,
			SimulateMsgCreateGroup(ak, bk, protoCdc),
		),
		simulation.NewWeightedOperation(
			weightMsgCreateGroupAccount,
			SimulateMsgCreateGroupAccount(ak, bk, protoCdc, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCreateProposal,
			SimulateMsgCreateProposal(ak, bk, protoCdc, qryClient),
		),
	}
}

func SimulateMsgCreateGroup(ak server.AccountKeeper, bk server.BankKeeper, protoCdc *codec.ProtoCodec) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc := accounts[0]

		account := ak.GetAccount(ctx, acc.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroup, "fee error"), nil, err
		}

		members := []group.Member{
			{
				Address:  acc.Address.String(),
				Weight:   fmt.Sprintf("%d", simappparams.DefaultWeightMsgCreateValidator),
				Metadata: []byte(simtypes.RandStringOfLength(r, 10)),
			},
		}
		srvMsg := &group.MsgCreateGroupRequest{Admin: acc.Address.String(), Members: members, Metadata: []byte(simtypes.RandStringOfLength(r, 10))}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		svcMsgClientConn := &msgservice.ServiceMsgClientConn{}
		msgClient := group.NewMsgClient(svcMsgClientConn)
		_, err = msgClient.CreateGroup(context.Background(), srvMsg)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroup, err.Error()), nil, err
		}
		tx, err := helpers.GenTx(
			txGen,
			svcMsgClientConn.GetMsgs(),
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
			return simtypes.NoOpMsg(group.ModuleName, svcMsgClientConn.GetMsgs()[0].Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(svcMsgClientConn.GetMsgs()[0], true, "", protoCdc), nil, err
	}
}

func SimulateMsgCreateGroupAccount(ak server.AccountKeeper, bk server.BankKeeper, protoCdc *codec.ProtoCodec, qryClient group.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc := accounts[0]

		account := ak.GetAccount(ctx, acc.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupAccount, "fee error"), nil, err
		}

		ctx1 := regentypes.Context{Context: ctx}
		result, err := qryClient.GroupAccountInfo(ctx1, &group.QueryGroupAccountInfoRequest{GroupAccount: acc.Address.String()})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupAccount, "fail to query group info"), nil, err
		}

		if result.Info == nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupAccount, "no group account found"), nil, nil
		}

		addr, err := sdk.AccAddressFromBech32(result.Info.Admin)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupAccount, "fail to decode acc address"), nil, err
		}

		msg, err := group.NewMsgCreateGroupAccountRequest(
			addr,
			group.ID(simtypes.RandIntBetween(r, 1, 10)), // TODO: replace with existed group-id
			[]byte(simtypes.RandStringOfLength(r, 10)),
			&group.ThresholdDecisionPolicy{
				Threshold: fmt.Sprintf("%d", simtypes.RandIntBetween(r, 1, 100)),
				Timeout:   proto.Duration{Seconds: int64(simtypes.RandIntBetween(r, 100, 1000))},
			},
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupAccount, err.Error()), nil, err
		}
		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		svcMsgClientConn := &msgservice.ServiceMsgClientConn{}
		msgClient := group.NewMsgClient(svcMsgClientConn)
		_, err = msgClient.CreateGroupAccount(context.Background(), msg)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupAccount, err.Error()), nil, err
		}
		tx, err := helpers.GenTx(
			txGen,
			svcMsgClientConn.GetMsgs(),
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
			return simtypes.NoOpMsg(group.ModuleName, svcMsgClientConn.GetMsgs()[0].Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(svcMsgClientConn.GetMsgs()[0], true, "", protoCdc), nil, err
	}
}

func SimulateMsgCreateProposal(ak server.AccountKeeper, bk server.BankKeeper, protoCdc *codec.ProtoCodec, queryClient group.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc := accounts[0]

		account := ak.GetAccount(ctx, acc.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateProposal, "fee error"), nil, err
		}

		ctx1 := regentypes.Context{Context: ctx}
		result, err := queryClient.GroupAccountInfo(ctx1, &group.QueryGroupAccountInfoRequest{GroupAccount: acc.Address.String()})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupAccount, "fail to query group info"), nil, err
		}

		if result.Info == nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupAccount, "no group account found"), nil, nil
		}

		msg := group.MsgCreateProposalRequest{
			GroupAccount: result.Info.GroupAccount,
			Proposers:    []string{acc.Address.String()},
			Metadata:     []byte(simtypes.RandStringOfLength(r, 10)),
		}
		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		svcMsgClientConn := &msgservice.ServiceMsgClientConn{}
		msgClient := group.NewMsgClient(svcMsgClientConn)

		_, err = msgClient.CreateProposal(context.Background(), &msg)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateProposal, err.Error()), nil, err
		}
		tx, err := helpers.GenTx(
			txGen,
			svcMsgClientConn.GetMsgs(),
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
			return simtypes.NoOpMsg(group.ModuleName, svcMsgClientConn.GetMsgs()[0].Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(svcMsgClientConn.GetMsgs()[0], true, "", protoCdc), nil, err
	}
}
