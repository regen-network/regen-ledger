package simulation

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/regen-network/regen-ledger/types/v2/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/basket/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/simulation/utils"
)

const WeightPut = 100

const OpWeightMsgPut = "op_weight_msg_put_into_basket" //nolint:gosec

var TypeMsgPut = types.MsgPut{}.Route()

// SimulateMsgPut generates a Basket/MsgPut with random values.
func SimulateMsgPut(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient basetypes.QueryServer, bsktQryClient types.QueryServer) simtypes.Operation {
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
		var classInfoList []basetypes.ClassInfo
		max := 0

		var ownerAddr string
		var owner simtypes.Account
		for _, class := range classes {
			if class.CreditTypeAbbrev == rBasket.CreditTypeAbbrev {
				issuersRes, err := qryClient.ClassIssuers(sdk.WrapSDKContext(sdkCtx), &basetypes.QueryClassIssuersRequest{
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

			resProjects, err := qryClient.ProjectsByClass(ctx, &basetypes.QueryProjectsByClassRequest{ClassId: classInfo.Id})
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, err.Error()), nil, err
			}

			for _, projectInfo := range resProjects.GetProjects() {

				batchesRes, err := qryClient.BatchesByProject(ctx, &basetypes.QueryBatchesByProjectRequest{
					ProjectId: projectInfo.Id,
				})
				if err != nil {
					return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgPut, err.Error()), nil, err
				}

				batches := batchesRes.Batches
				if len(batches) != 0 {
					count := 0
					for _, batch := range batches {
						balanceRes, err := qryClient.Balance(ctx, &basetypes.QueryBalanceRequest{
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
