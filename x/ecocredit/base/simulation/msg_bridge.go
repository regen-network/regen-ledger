package simulation

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/types/v2/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/simulation/utils"
)

const OpWeightMsgBridge = "op_weight_msg_bridge" //nolint:gosec

var TypeMsgBridge = sdk.MsgTypeURL(&types.MsgBridge{})

const WeightBridge = 33

// SimulateMsgBridge generates a MsgBridge with random values.
func SimulateMsgBridge(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		ctx := sdk.WrapSDKContext(sdkCtx)
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgBridge)
		if class == nil {
			return op, nil, err
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgBridge, class.Id)
		if project == nil {
			return op, nil, err
		}

		batch, op, err := getRandomBatchFromProject(ctx, r, qryClient, TypeMsgBridge, project.Id)
		if batch == nil {
			return op, nil, err
		}

		issuersRes, err := qryClient.ClassIssuers(ctx, &types.QueryClassIssuersRequest{
			ClassId: class.Id,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBridge, err.Error()), nil, err
		}

		issuers := issuersRes.Issuers
		owner := issuers[r.Intn(len(issuers))]
		ownerAddr, err := sdk.AccAddressFromBech32(owner)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBridge, err.Error()), nil, err
		}

		_, found := simtypes.FindAccount(accs, ownerAddr)
		if !found {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBridge, "not a simulation account"), nil, nil
		}

		balanceRes, err := qryClient.Balance(ctx, &types.QueryBalanceRequest{
			Address:    owner,
			BatchDenom: batch.Denom,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBridge, err.Error()), nil, err
		}

		tradableBalance, err := math.NewNonNegativeDecFromString(balanceRes.Balance.TradableAmount)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBridge, err.Error()), nil, err
		}

		if tradableBalance.IsZero() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBridge, "balance is zero"), nil, nil
		}

		tradable, err := tradableBalance.Int64()
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBridge, err.Error()), nil, nil
		}

		amount := 1
		if tradable > 1 {
			amount = simtypes.RandIntBetween(r, 1, int(tradable))
		}

		msg := &types.MsgBridge{
			Target:    "polygon",
			Recipient: "0x323b5d4c32345ced77393b3530b1eed0f346429d",
			Owner:     owner,
			Credits: []*types.Credits{
				{
					BatchDenom: batch.Denom,
					Amount:     fmt.Sprintf("%d", amount),
				},
			},
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, owner, TypeMsgBridge)
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

		fees, err := simtypes.RandomFees(r, sdkCtx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBridge, "fee error"), nil, err
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		acc := txCtx.AccountKeeper.GetAccount(txCtx.Context, txCtx.SimAccount.Address)

		tx, err := helpers.GenSignedMockTx(
			r,
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{acc.GetAccountNumber()},
			[]uint64{acc.GetSequence()},
			account.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBridge, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
		if err != nil {
			if !strings.Contains(err.Error(), "only credits previously bridged from another chain") {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBridge, "unable to deliver tx"), nil, err
			}
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}
