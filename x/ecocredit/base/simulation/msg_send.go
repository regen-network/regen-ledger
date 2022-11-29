package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/types/v2/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/simulation/utils"
)

const OpWeightMsgSend = "op_weight_msg_send" //nolint:gosec

var TypeMsgSend = sdk.MsgTypeURL(&types.MsgSend{})

const WeightSend = 100

// SimulateMsgSend generates a MsgSend with random values.
func SimulateMsgSend(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		ctx := sdk.WrapSDKContext(sdkCtx)
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
		balres, err := qryClient.Balance(ctx, &types.QueryBalanceRequest{
			Address:    admin,
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

		msg := &types.MsgSend{
			Sender:    admin,
			Recipient: recipient.Address.String(),
			Credits: []*types.MsgSend_SendCredits{
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

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}
