package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/x/ecocredit/v4"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/simulation/utils"
)

const OpWeightMsgSetClassCreatorAllowlist = "op_weight_msg_set_class_creator_allowlist" //nolint:gosec

var TypeMsgSetClassCreatorAllowlist = sdk.MsgTypeURL(&types.MsgSetClassCreatorAllowlist{})

const WeightSetClassCreatorAllowlist = 33

// SimulateMsgSetClassCreatorAllowlist generates a MsgSetClassCreatorAllowlist with random values.
func SimulateMsgSetClassCreatorAllowlist(txCfg client.TxConfig, ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, govk govkeeper.Keeper,
	_ types.QueryServer, authority sdk.AccAddress,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		proposer, _ := simtypes.RandomAcc(r, accs)
		proposerAddr := proposer.Address.String()

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, proposerAddr, TypeMsgSetClassCreatorAllowlist)
		if spendable == nil {
			return op, nil, err
		}

		params, err := govk.Params.Get(sdkCtx)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddCreditType, err.Error()), nil, err
		}

		deposit, skip, err := utils.RandomDeposit(r, sdkCtx, ak, bk, params.MinDeposit, proposer.Address)
		switch {
		case skip:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSetClassCreatorAllowlist, "skip deposit"), nil, nil
		case err != nil:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSetClassCreatorAllowlist, "unable to generate deposit"), nil, err
		}

		proposalMsg := types.MsgSetClassCreatorAllowlist{
			Authority: authority.String(),
			Enabled:   r.Float32() < 0.3, // 30% chance of allowlist being enabled,
		}

		anyMsg, err := codectypes.NewAnyWithValue(&proposalMsg)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSetClassCreatorAllowlist, err.Error()), nil, err
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
			TxGen:           txCfg,
			Cdc:             nil,
			Msg:             msg,
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCfg, txCtx)
	}
}
