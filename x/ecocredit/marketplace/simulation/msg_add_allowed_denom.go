package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/simulation/utils"
)

const OpWeightMsgAddAllowedDenom = "op_weight_msg_add_allowed_denom" //nolint:gosec

const WeightAddAllowedDenom = 100

var TypeMsgAddAllowedDenom = types.MsgAddAllowedDenom{}.Route()

func SimulateMsgAddAllowedDenom(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient types.QueryServer, govk ecocredit.GovKeeper, authority sdk.AccAddress) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		proposer, _ := simtypes.RandomAcc(r, accs)
		proposerAddr := proposer.Address.String()

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, proposerAddr, TypeMsgAddAllowedDenom)
		if spendable == nil {
			return op, nil, err
		}

		bankDenom := simtypes.RandStringOfLength(r, 4)
		res, err := qryClient.AllowedDenoms(sdkCtx, &types.QueryAllowedDenomsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddAllowedDenom, err.Error()), nil, err
		}

		if isDenomExists(res.AllowedDenoms, bankDenom) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddAllowedDenom, fmt.Sprintf("denom %s already exists", bankDenom)), nil, nil
		}

		params := govk.GetParams(sdkCtx)
		deposit, skip, err := utils.RandomDeposit(r, sdkCtx, ak, bk, params.MinDeposit, authority)
		switch {
		case skip:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddAllowedDenom, "skip deposit"), nil, nil
		case err != nil:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddAllowedDenom, "unable to generate deposit"), nil, err
		}

		msg := types.MsgAddAllowedDenom{
			Authority:    authority.String(),
			BankDenom:    bankDenom,
			DisplayDenom: bankDenom,
			Exponent:     6,
		}

		anyMsg, err := codectypes.NewAnyWithValue(&msg)
		if err != nil {
			return simtypes.NoOpMsg(TypeMsgAddAllowedDenom, TypeMsgAddAllowedDenom, err.Error()), nil, err
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
			R:               r,
			App:             app,
			TxGen:           testutil.MakeTestEncodingConfig().TxConfig,
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

func isDenomExists(allowedDenom []*types.AllowedDenom, bankDenom string) bool {
	for _, denom := range allowedDenom {
		if denom.BankDenom == bankDenom {
			return true
		}
	}

	return false
}
