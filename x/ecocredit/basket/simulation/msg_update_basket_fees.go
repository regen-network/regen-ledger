package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/x/ecocredit/v4"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/basket/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/simulation/utils"
)

const OpWeightMsgUpdateBasketFee = "op_weight_msg_update_basket_fee" //nolint:gosec

var TypeMsgUpdateBasketFee = sdk.MsgTypeURL(&types.MsgUpdateBasketFee{})

const WeightUpdateBasketFees = 100

// SimulateMsgUpdateBasketFee generates a Basket/MsgUpdateBasketFee with random values.
func SimulateMsgUpdateBasketFee(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, _ basetypes.QueryServer,
	_ types.QueryServer, govk ecocredit.GovKeeper, authority sdk.AccAddress) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		proposer, _ := simtypes.RandomAcc(r, accs)
		proposerAddr := proposer.Address.String()

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, proposerAddr, TypeMsgUpdateBasketFee)
		if spendable == nil {
			return op, nil, err
		}

		params := govk.GetParams(sdkCtx)
		deposit, skip, err := utils.RandomDeposit(r, sdkCtx, ak, bk, params.MinDeposit, proposer.Address)
		switch {
		case skip:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateBasketFee, "skip deposit"), nil, nil
		case err != nil:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateBasketFee, "unable to generate deposit"), nil, err
		}

		fee := utils.RandomFee(r)
		if fee.Amount.IsZero() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateBasketFee, "invalid proposal message"), nil, err
		}

		msg := types.MsgUpdateBasketFee{
			Authority: authority.String(),
			Fee:       &fee,
		}

		anyMsg, err := codectypes.NewAnyWithValue(&msg)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateBasketFee, err.Error()), nil, err
		}

		proposalMsg := govtypes.MsgSubmitProposal{
			Title:          simtypes.RandStringOfLength(r, 10),
			Summary:        simtypes.RandStringOfLength(r, 10),
			InitialDeposit: deposit,
			Proposer:       proposerAddr,
			Metadata:       simtypes.RandStringOfLength(r, 10),
			Messages:       []*codectypes.Any{anyMsg},
		}

		txCtx := simulation.OperationInput{
			R:     r,
			App:   app,
			TxGen: moduletestutil.MakeTestEncodingConfig().TxConfig,
			Cdc:   nil,
			Msg:   &proposalMsg,
			// MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}
