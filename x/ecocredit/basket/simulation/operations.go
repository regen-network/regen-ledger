package basketsims

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	gogotypes "github.com/gogo/protobuf/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	types "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation/utils"
)

// Simulation operation weights constants
const (
	OpWeightMsgCreate = "op_weight_msg_create_basket"   //nolint:gosec
	OpWeightMsgPut    = "op_weight_msg_put_into_basket" //nolint:gosec
	OpWeightMsgTake   = "op_weight_take_from_basket"    //nolint:gosec
)

// basket operations weights
const (
	WeightCreate = 100
	WeightPut    = 100
	WeightTake   = 100
)

// ecocredit message types
var (
	TypeMsgCreate = types.MsgCreate{}.Route()
	TypeMsgPut    = types.MsgPut{}.Route()
	TypeMsgTake   = types.MsgTake{}.Route()
)

func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryServer, basketQryClient types.QueryServer) simulation.WeightedOperations {

	var (
		weightMsgCreate int
		weightMsgPut    int
		weightMsgTake   int
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCreate, &weightMsgCreate, nil,
		func(_ *rand.Rand) {
			weightMsgCreate = WeightCreate
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgPut, &weightMsgPut, nil,
		func(_ *rand.Rand) {
			weightMsgPut = WeightPut
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgTake, &weightMsgCreate, nil,
		func(_ *rand.Rand) {
			weightMsgTake = WeightTake
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgCreate,
			SimulateMsgCreate(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgPut,
			SimulateMsgPut(ak, bk, qryClient, basketQryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgTake,
			SimulateMsgTake(ak, bk, qryClient, basketQryClient),
		),
	}
}

// SimulateMsgCreate generates a Basket/MsgCreate with random values.
func SimulateMsgCreate(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		curator, _ := simtypes.RandomAcc(r, accs)

		ctx := sdk.WrapSDKContext(sdkCtx)
		res, err := qryClient.Params(ctx, &core.QueryParamsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, err.Error()), nil, err
		}

		params := res.Params
		spendable := bk.SpendableCoins(sdkCtx, curator.Address)
		if !spendable.IsAllGTE(params.BasketFee) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, "not enough balance"), nil, nil
		}

		creditType, err := randomCreditType(ctx, r, qryClient)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, err.Error()), nil, err
		}

		if creditType == nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreate, "credit type not found"), nil, nil
		}

		classIDs, op, err := randomClassIds(r, sdkCtx, qryClient, creditType.Abbreviation, TypeMsgPut)
		if len(classIDs) == 0 {
			return op, nil, err
		}

		precision := creditType.Precision
		dateCriteria := randomDateCriteria(r, sdkCtx)
		msg := &types.MsgCreate{
			Name:              simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 3, 8)),
			Description:       simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 3, 256)),
			Fee:               params.BasketFee,
			DisableAutoRetire: r.Float32() < 0.5,
			Curator:           curator.Address.String(),
			Exponent:          utils.RandomExponent(r, precision),
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

// SimulateMsgPut generates a Basket/MsgPut with random values.
func SimulateMsgPut(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryServer, bsktQryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		ctx := sdk.WrapSDKContext(sdkCtx)
		res, err := bsktQryClient.Baskets(ctx, &types.QueryBasketsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, err.Error()), nil, err
		}

		baskets := res.Baskets
		if len(baskets) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, "no baskets"), nil, nil
		}

		classes, op, err := utils.GetClasses(sdkCtx, r, qryClient, TypeMsgPut)
		if len(classes) == 0 {
			return op, nil, err
		}

		rBasket := baskets[r.Intn(len(baskets))]
		var classInfoList []core.ClassInfo
		max := 0

		var ownerAddr string
		var owner simtypes.Account
		for _, class := range classes {
			if class.CreditTypeAbbrev == rBasket.CreditTypeAbbrev {
				issuersRes, err := qryClient.ClassIssuers(sdk.WrapSDKContext(sdkCtx), &core.QueryClassIssuersRequest{
					ClassId: class.Id,
				})
				if err != nil {
					return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, err.Error()), nil, err
				}
				issuers := issuersRes.Issuers
				if len(issuers) == 0 {
					return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, "no class issuers"), nil, nil
				}

				if ownerAddr == "" {
					bechAddr, err := sdk.AccAddressFromBech32(issuers[0])
					if err != nil {
						return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, err.Error()), nil, err
					}

					acc, found := simtypes.FindAccount(accs, bechAddr)
					if found {
						ownerAddr = issuers[0]
						owner = acc
						classInfoList = append(classInfoList, *class)
						max++
					}
				} else if utils.Contains(issuers, ownerAddr) {
					classInfoList = append(classInfoList, *class)
					max++
				}

				if max == 2 {
					break
				}
			}
		}
		if len(classInfoList) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, "no classes"), nil, nil
		}

		var credits []*types.BasketCredit
		for _, classInfo := range classInfoList {

			resProjects, err := qryClient.ProjectsByClass(ctx, &core.QueryProjectsByClassRequest{ClassId: classInfo.Id})
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, err.Error()), nil, err
			}

			for _, projectInfo := range resProjects.GetProjects() {

				batchesRes, err := qryClient.BatchesByProject(ctx, &core.QueryBatchesByProjectRequest{
					ProjectId: projectInfo.Id,
				})
				if err != nil {
					return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, err.Error()), nil, err
				}

				batches := batchesRes.Batches
				if len(batches) != 0 {
					count := 0
					for _, batch := range batches {
						balanceRes, err := qryClient.Balance(ctx, &core.QueryBalanceRequest{
							Address: ownerAddr, BatchDenom: batch.Denom,
						})
						if err != nil {
							return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, err.Error()), nil, err
						}

						tradableAmount := balanceRes.Balance.TradableAmount
						if tradableAmount != "0" {
							d, err := math.NewPositiveDecFromString(tradableAmount)
							if err != nil {
								return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, err.Error()), nil, nil
							}

							dInt, err := d.Int64()
							if err != nil {
								return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, err.Error()), nil, nil
							}

							if dInt == 1 {
								credits = append(credits, &types.BasketCredit{
									BatchDenom: batch.Denom,
									Amount:     "1",
								})
								count++
							} else {
								amt := simtypes.RandIntBetween(r, 1, int(dInt))
								credits = append(credits, &types.BasketCredit{
									BatchDenom: batch.Denom,
									Amount:     fmt.Sprintf("%d", amt),
								})
								count++
							}
						}

						if count == 3 {
							break
						}
					}
				}
			}
		}
		if len(credits) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, "no basket credits"), nil, nil
		}

		msg := &types.MsgPut{
			Owner:       owner.Address.String(),
			BasketDenom: rBasket.BasketDenom,
			Credits:     credits,
		}
		spendable := bk.SpendableCoins(sdkCtx, owner.Address)
		fees, err := simtypes.RandomFees(r, sdkCtx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, "fee error"), nil, err
		}

		account := ak.GetAccount(sdkCtx, owner.Address)
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
			owner.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
		if err != nil {
			if strings.Contains(err.Error(), "is not allowed in this basket") {
				return simtypes.NoOpMsg(ecocredit.ModuleName, msg.Type(), "class is not allowed"), nil, nil
			}

			return simtypes.NoOpMsg(ecocredit.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}

// SimulateMsgTake generates a Basket/MsgTake with random values.
func SimulateMsgTake(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryServer, bsktQryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		owner, _ := simtypes.RandomAcc(r, accs)
		ownerAddr := owner.Address.String()

		ctx := sdk.WrapSDKContext(sdkCtx)
		res, err := bsktQryClient.Baskets(ctx, &types.QueryBasketsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, err.Error()), nil, err
		}

		baskets := res.BasketsInfo
		if len(baskets) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, "no baskets"), nil, nil
		}

		var rBasket *types.BasketInfo
		var bBalances []*types.BasketBalanceInfo
		for _, b := range baskets {
			balancesRes, err := bsktQryClient.BasketBalances(ctx, &types.QueryBasketBalancesRequest{
				BasketDenom: b.BasketDenom,
			})
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, err.Error()), nil, err
			}
			balances := balancesRes.BalancesInfo
			if len(balances) != 0 {
				rBasket = b
				bBalances = balances
				break
			}
		}
		if rBasket == nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, "no basket"), nil, nil
		}

		var amt int
		for _, b := range bBalances {
			iAmount, err := strconv.Atoi(b.Balance)
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, err.Error()), nil, nil
			}

			switch {
			case iAmount == 0:
				continue
			case iAmount == 1:
				amt = iAmount
			default:
				amt = simtypes.RandIntBetween(r, 1, iAmount)
			}
		}
		if amt == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, "basket balance"), nil, nil
		}

		msg := &types.MsgTake{
			Owner:                  ownerAddr,
			BasketDenom:            rBasket.BasketDenom,
			Amount:                 fmt.Sprintf("%d", amt),
			RetirementJurisdiction: "AQ",
			RetireOnTake:           !rBasket.DisableAutoRetire,
		}

		spendable := bk.SpendableCoins(sdkCtx, owner.Address)
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      owner,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

func randomClassIds(r *rand.Rand, ctx sdk.Context, qryClient core.QueryServer,
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

func randomCreditType(ctx context.Context, r *rand.Rand, qryClient core.QueryServer) (*core.CreditType, error) {
	res, err := qryClient.CreditTypes(ctx, &core.QueryCreditTypesRequest{})
	if err != nil {
		return nil, err
	}

	creditTypes := res.CreditTypes
	if len(creditTypes) == 0 {
		return nil, nil
	}

	return creditTypes[r.Intn(len(creditTypes))], nil
}
