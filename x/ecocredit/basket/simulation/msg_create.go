package simulation

import (
	"context"
	"fmt"
	"math/rand"
	"strings"

	gogotypes "github.com/gogo/protobuf/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	types "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation/utils"
)

const WeightCreate = 100

var TypeMsgCreate = types.MsgCreate{}.Route()

const OpWeightMsgCreate = "op_weight_msg_create_basket" //nolint:gosec

// SimulateMsgCreate generates a Basket/MsgCreate with random values.
func SimulateMsgCreate(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	baseClient basetypes.QueryServer, client types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		curator, _ := simtypes.RandomAcc(r, accs)

		ctx := sdk.WrapSDKContext(sdkCtx)
		res, err := baseClient.Params(ctx, &basetypes.QueryParamsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, err.Error()), nil, err
		}

		params := res.Params
		spendable := bk.SpendableCoins(sdkCtx, curator.Address)
		if !spendable.IsAllGTE(params.BasketFee) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, "not enough balance"), nil, nil
		}

		creditType, err := randomCreditType(ctx, r, baseClient)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, err.Error()), nil, err
		}

		if creditType == nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, "credit type not found"), nil, nil
		}

		classIDs, op, err := randomClassIds(r, sdkCtx, baseClient, creditType.Abbreviation, TypeMsgPut)
		if len(classIDs) == 0 {
			return op, nil, err
		}

		precision := creditType.Precision
		exponent := utils.RandomExponent(r, precision)
		basketName := simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 3, 8))
		denom, _, err := basket.FormatBasketDenom(basketName, creditType.Abbreviation, exponent)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, "failed to generate basket denom"), nil, err
		}

		result, err := client.Basket(sdkCtx, &types.QueryBasketRequest{
			BasketDenom: denom,
		})
		if err != nil && !ormerrors.NotFound.Is(err) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, err.Error()), nil, err
		}

		if result != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, fmt.Sprintf("basket with name %s already exists", basketName)), nil, nil
		}

		dateCriteria := randomDateCriteria(r, sdkCtx)
		msg := &types.MsgCreate{
			Name:              basketName,
			Description:       simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 3, 256)),
			Fee:               params.BasketFee,
			DisableAutoRetire: r.Float32() < 0.5,
			Curator:           curator.Address.String(),
			Exponent:          exponent,
			AllowedClasses:    classIDs,
			CreditTypeAbbrev:  creditType.Abbreviation,
			DateCriteria:      dateCriteria,
		}

		fees, err := simtypes.RandomFees(r, sdkCtx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, "fee error"), nil, err
		}

		account := ak.GetAccount(sdkCtx, curator.Address)
		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenSignedMockTx(
			r,
			txGen,
			[]sdk.Msg{msg},
			fees,
			2000000,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			curator.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
		if err != nil {
			if strings.Contains(err.Error(), "basket specified credit type") {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, err.Error()), nil, nil
			}

			if strings.Contains(err.Error(), "insufficient funds") {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, err.Error()), nil, nil
			}
			return simtypes.NoOpMsg(ecocredit.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil

	}
}

func randomCreditType(ctx context.Context, r *rand.Rand, qryClient basetypes.QueryServer) (*basetypes.CreditType, error) {
	res, err := qryClient.CreditTypes(ctx, &basetypes.QueryCreditTypesRequest{})
	if err != nil {
		return nil, err
	}

	creditTypes := res.CreditTypes
	if len(creditTypes) == 0 {
		return nil, nil
	}

	return creditTypes[r.Intn(len(creditTypes))], nil
}

func randomDateCriteria(r *rand.Rand, ctx sdk.Context) *types.DateCriteria {
	// 30% chance of date-criteria being enable
	includeCriteria := r.Int63n(101) <= 30
	if includeCriteria {
		seconds := ctx.BlockTime().AddDate(0, -1, 0).Unix()
		if r.Float32() < 0.5 {
			return &types.DateCriteria{
				MinStartDate: &gogotypes.Timestamp{
					Seconds: seconds,
				},
			}
		}
		return &types.DateCriteria{
			StartDateWindow: &gogotypes.Duration{Seconds: seconds},
		}
	}
	return nil
}

func randomClassIds(r *rand.Rand, ctx sdk.Context, qryClient basetypes.QueryServer,
	creditTypeAbbrev string, msgType string) ([]string, simtypes.OperationMsg, error) {
	classes, op, err := utils.GetClasses(ctx, r, qryClient, msgType)
	if len(classes) == 0 {
		return []string{}, op, err
	}

	if len(classes) == 1 {
		return []string{classes[0].Id}, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
	}

	max := simtypes.RandIntBetween(r, 1, min(5, len(classes)))
	var classIDs []string
	for i := 0; i < max; i++ {
		class := classes[i]
		if class.CreditTypeAbbrev == creditTypeAbbrev {
			classIDs = append(classIDs, class.Id)
		}
	}

	return classIDs, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
