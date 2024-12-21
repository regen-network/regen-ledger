package simulation

import (
	"math/rand"
	"strings"

	"github.com/cosmos/cosmos-sdk/baseapp"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/simulation/utils"
)

const OpWeightMsgAddCreditType = "op_weight_msg_add_credit_type" //nolint:gosec

var TypeMsgAddCreditType = sdk.MsgTypeURL(&types.MsgAddCreditType{})

const WeightAddCreditType = 33

// SimulateMsgAddCreditType generates a MsgAddCreditType with random values.
func SimulateMsgAddCreditType(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, govk ecocredit.GovKeeper,
	qryClient types.QueryServer, authority sdk.AccAddress) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		proposer, _ := simtypes.RandomAcc(r, accs)
		proposerAddr := proposer.Address.String()

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, proposerAddr, TypeMsgAddCreditType)
		if spendable == nil {
			return op, nil, err
		}

		params := govk.GetParams(sdkCtx)
		deposit, skip, err := utils.RandomDeposit(r, sdkCtx, ak, bk, params.MinDeposit, proposer.Address)
		switch {
		case skip:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddCreditType, "skip deposit"), nil, nil
		case err != nil:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddCreditType, "unable to generate deposit"), nil, err
		}

		abbrev := simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 1, 3))
		abbrev = strings.ToUpper(abbrev)
		name := simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 1, 10))

		_, err = qryClient.CreditType(sdkCtx, &types.QueryCreditTypeRequest{
			Abbreviation: abbrev,
		})
		if err != nil {
			if !regenerrors.ErrNotFound.Is(err) {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddCreditType, err.Error()), nil, err
			}
		}

		proposalMsg := types.MsgAddCreditType{
			Authority: authority.String(),
			CreditType: &types.CreditType{
				Abbreviation: abbrev,
				Name:         name,
				Unit:         "kg",
				Precision:    6,
			},
		}

		anyMsg, err := codectypes.NewAnyWithValue(&proposalMsg)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddCreditType, err.Error()), nil, err
		}

		msg := &govtypes.MsgSubmitProposal{
			Title:          simtypes.RandStringOfLength(r, 10),
			Messages:       []*codectypes.Any{anyMsg},
			InitialDeposit: deposit,
			Proposer:       proposerAddr,
			Metadata:       simtypes.RandStringOfLength(r, 10),
			Summary:        simtypes.RandStringOfLength(r, 10),
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           moduletestutil.MakeTestEncodingConfig().TxConfig,
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

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}
