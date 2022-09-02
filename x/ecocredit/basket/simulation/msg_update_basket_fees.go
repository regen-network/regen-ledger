package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation/utils"
)

const OpWeightMsgUpdateBasketFees = "op_weight_msg_update_basket_fees" //nolint:gosec

var TypeMsgUpdateBasketFees = types.MsgUpdateBasketFees{}.Route()

const WeightUpdateBasketFees = 100

// SimulateMsgCreate generates a Basket/MsgUpdateBasketFees with random values.
func SimulateMsgUpdateBasketFees(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, qryClient basetypes.QueryServer,
	basketQryClient types.QueryServer, govk ecocredit.GovKeeper, authority sdk.AccAddress) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		proposer, _ := simtypes.RandomAcc(r, accs)
		proposerAddr := proposer.Address.String()

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, proposerAddr, TypeMsgUpdateBasketFees)
		if spendable == nil {
			return op, nil, err
		}

		params := govk.GetDepositParams(sdkCtx)
		deposit, skip, err := utils.RandomDeposit(r, sdkCtx, ak, bk, params, proposer.Address)
		switch {
		case skip:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateBasketFees, "skip deposit"), nil, nil
		case err != nil:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateBasketFees, "unable to generate deposit"), nil, err
		}

		fees := utils.RandomFees(r)
		msg := types.MsgUpdateBasketFees{
			Authority:  authority.String(),
			BasketFees: fees,
		}

		any, err := codectypes.NewAnyWithValue(&msg)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateBasketFees, err.Error()), nil, err
		}

		proposalMsg := govtypes.MsgSubmitProposal{
			InitialDeposit: deposit,
			Proposer:       proposerAddr,
			Metadata:       simtypes.RandStringOfLength(r, 10),
			Messages:       []*codectypes.Any{any},
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simapp.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             &proposalMsg,
			MsgType:         msg.Type(),
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
