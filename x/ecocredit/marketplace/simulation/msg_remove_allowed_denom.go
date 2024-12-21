package simulation

import (
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

const OpWeightMsgRemoveAllowedDenom = "op_weight_msg_remove_allowed_denom" //nolint:gosec

const WeightRemoveAllowedDenom = 100

var TypeMsgRemoveAllowedDenom = types.MsgRemoveAllowedDenom{}.Route()

func SimulateMsgRemoveAllowedDenom(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	mktClient types.QueryServer, govk ecocredit.GovKeeper, authority sdk.AccAddress) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		response, err := mktClient.AllowedDenoms(sdkCtx, &types.QueryAllowedDenomsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRemoveAllowedDenom, err.Error()), nil, err
		}

		if len(response.AllowedDenoms) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRemoveAllowedDenom, "no allowed denom present"), nil, nil
		}

		proposer, _ := simtypes.RandomAcc(r, accs)
		proposerAddr := proposer.Address.String()

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, proposerAddr, TypeMsgRemoveAllowedDenom)
		if spendable == nil {
			return op, nil, err
		}

		params := govk.GetParams(sdkCtx)
		deposit, skip, err := utils.RandomDeposit(r, sdkCtx, ak, bk, params.MinDeposit, authority)
		switch {
		case skip:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRemoveAllowedDenom, "skip deposit"), nil, nil
		case err != nil:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRemoveAllowedDenom, "unable to generate deposit"), nil, err
		}

		msg := types.MsgRemoveAllowedDenom{
			Authority: authority.String(),
			Denom:     response.AllowedDenoms[r.Intn(len(response.AllowedDenoms))].BankDenom,
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
