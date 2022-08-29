package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type totalBalances struct {
	batchDenoms    []string
	tradableAmount math.Dec
	retiredAmount  math.Dec
	escrowedAmount math.Dec
}

func (k Keeper) AllBalances(ctx context.Context, req *core.QueryAllBalancesRequest) (*core.QueryAllBalancesResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err
	}
	it, err := k.stateStore.BatchBalanceTable().List(ctx, api.BatchBalancePrimaryKey{}, ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	balanceMap := make(map[string]totalBalances)
	for it.Next() {
		balance, err := it.Value()
		if err != nil {
			return nil, err
		}

		batch, err := k.stateStore.BatchTable().Get(ctx, balance.BatchKey)
		if err != nil {
			return nil, err
		}
		owner := sdk.AccAddress(balance.Address).String()
		bal := balanceMap[owner]
		newBal, err := bal.Add(balance, batch.Denom)
		if err != nil {
			return nil, sdkerrors.ErrLogic.Wrapf("unable to add balance for %s in %s denom: %s", owner, batch.Denom, err.Error())
		}
		balanceMap[owner] = newBal
	}
	var res core.QueryAllBalancesResponse
	for a, b := range balanceMap {
		res.Balances = append(res.Balances, &core.CreditBalances{
			Address: a,
			Balances: &core.CreditBalance{
				BatchDenoms:    b.batchDenoms,
				TradableAmount: b.tradableAmount.String(),
				RetiredAmount:  b.retiredAmount.String(),
				EscrowedAmount: b.escrowedAmount.String(),
			},
		})
	}
	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}
	res.Pagination = pr
	return &res, nil
}

func (b totalBalances) Add(balance *api.BatchBalance, denom string) (totalBalances, error) {
	tradable, err := math.NewNonNegativeDecFromString(balance.TradableAmount)
	if err != nil {
		return totalBalances{}, err
	}
	retired, err := math.NewNonNegativeDecFromString(balance.RetiredAmount)
	if err != nil {
		return totalBalances{}, err
	}
	escrowed, err := math.NewNonNegativeDecFromString(balance.EscrowedAmount)
	if err != nil {
		return totalBalances{}, err
	}

	sumTradable, err := b.tradableAmount.Add(tradable)
	if err != nil {
		return totalBalances{}, err
	}
	sumRetired, err := b.retiredAmount.Add(retired)
	if err != nil {
		return totalBalances{}, err
	}
	sumEscrowed, err := b.escrowedAmount.Add(escrowed)
	if err != nil {
		return totalBalances{}, err
	}
	b.batchDenoms = append(b.batchDenoms, denom)
	b.tradableAmount = sumTradable
	b.retiredAmount = sumRetired
	b.escrowedAmount = sumEscrowed
	return b, nil
}
