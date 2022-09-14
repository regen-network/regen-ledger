package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation/utils"
)

const OpWeightMsgUpdateClassIssuers = "op_weight_msg_update_class_issuers" //nolint:gosec

var TypeMsgUpdateClassIssuers = sdk.MsgTypeURL(&types.MsgUpdateClassIssuers{})

const MsgUpdateClassIssuers = 33

// SimulateMsgUpdateClassIssuers generates a MsgUpdateClassMetaData with random values
func SimulateMsgUpdateClassIssuers(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgUpdateClassIssuers)
		if class == nil {
			return op, nil, err
		}

		admin := sdk.AccAddress(class.Admin)
		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, admin.String(), TypeMsgUpdateClassIssuers)
		if spendable == nil {
			return op, nil, err
		}

		issuersRes, err := qryClient.ClassIssuers(sdk.WrapSDKContext(sdkCtx), &types.QueryClassIssuersRequest{ClassId: class.Id})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateClassIssuers, err.Error()), nil, err
		}
		classIssuers := issuersRes.Issuers

		var addIssuers []string
		var removeIssuers []string

		issuers := randomIssuers(r, accs)
		if len(issuers) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateClassIssuers, "empty issuers"), nil, nil
		}

		for _, i := range issuers {
			if utils.Contains(classIssuers, i) {
				removeIssuers = append(removeIssuers, i)
			} else {
				addIssuers = append(addIssuers, i)
			}
		}

		msg := &types.MsgUpdateClassIssuers{
			Admin:         admin.String(),
			ClassId:       class.Id,
			AddIssuers:    addIssuers,
			RemoveIssuers: removeIssuers,
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
