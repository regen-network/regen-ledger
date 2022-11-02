package simulation

import (
	"math/rand"
	"strings"

	"github.com/cosmos/cosmos-sdk/baseapp"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	regenerrors "github.com/regen-network/regen-ledger/errors"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation/utils"
)

const OpWeightMsgAddCreditType = "op_weight_msg_add_credit_type" //nolint:gosec

var TypeMsgAddCreditType = sdk.MsgTypeURL(&types.MsgAddCreditType{})

const WeightAddCreditType = 33

// SimulateMsgAddCreditType generates a MsgAddCreditType with random values.
func SimulateMsgAddCreditType(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, govk ecocredit.GovKeeper,
	qryClient types.QueryServer, authority sdk.AccAddress) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		proposer, _ := simtypes.RandomAcc(r, accs)
		proposerAddr := proposer.Address.String()

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, proposerAddr, TypeMsgAddCreditType)
		if spendable == nil {
			return op, nil, err
		}

		params := govk.GetDepositParams(sdkCtx)
		deposit, skip, err := utils.RandomDeposit(r, sdkCtx, ak, bk, params, proposer.Address)
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

		any, err := codectypes.NewAnyWithValue(&proposalMsg)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddCreditType, err.Error()), nil, err
		}

		msg := &govtypes.MsgSubmitProposal{
			Messages:       []*codectypes.Any{any},
			InitialDeposit: deposit,
			Proposer:       proposerAddr,
			Metadata:       simtypes.RandStringOfLength(r, 10),
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

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}
