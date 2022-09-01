package keeper

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types/math"
)

// PruneSellOrders is a BeginBlock function that moves escrowed credits back into their tradable balance and deletes orders
// that have expired.
func (k Keeper) PruneSellOrders(ctx context.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// we set the min to 1 ns because nil expirations are encoded as the 0 value timestamp,
	// and we DO NOT want those to be deleted/unescrowed.
	// https://github.com/cosmos/cosmos-sdk/issues/11980
	min, blockTime := timestamppb.New(time.Unix(0, 1)), timestamppb.New(sdkCtx.BlockTime())
	fromKey, toKey := api.SellOrderExpirationIndexKey{}.WithExpiration(min), api.SellOrderExpirationIndexKey{}.WithExpiration(blockTime)

	it, err := k.stateStore.SellOrderTable().ListRange(ctx, fromKey, toKey)
	if err != nil {
		return err
	}
	for it.Next() {
		sellOrder, err := it.Value()
		if err != nil {
			return err
		}
		if err = k.unescrowCredits(ctx, sellOrder.Seller, sellOrder.BatchKey, sellOrder.Quantity); err != nil {
			return err
		}

	}
	it.Close()

	return k.stateStore.SellOrderTable().DeleteRange(ctx, fromKey, toKey)
}

// unescrowCredits moves `amount` of credits from the sellerAddr's escrowed balance, into their tradable balance.
func (k Keeper) unescrowCredits(ctx context.Context, sellerAddr sdk.AccAddress, batchKey uint64, amount string) error {

	creditAmt, err := math.NewDecFromString(amount)
	if err != nil {
		return err
	}

	moveCredits := func(escrowed, tradable string) (string, string, error) {
		escrowedDec, err := math.NewDecFromString(escrowed)
		if err != nil {
			return "", "", err
		}
		tradableDec, err := math.NewDecFromString(tradable)
		if err != nil {
			return "", "", err
		}

		escrowedDec, err = math.SafeSubBalance(escrowedDec, creditAmt)
		if err != nil {
			return "", "", err
		}
		tradableDec, err = math.SafeAddBalance(tradableDec, creditAmt)
		if err != nil {
			return "", "", err
		}
		return escrowedDec.String(), tradableDec.String(), nil
	}

	bal, err := k.coreStore.BatchBalanceTable().Get(ctx, sellerAddr, batchKey)
	if err != nil {
		return err
	}
	bal.EscrowedAmount, bal.TradableAmount, err = moveCredits(bal.EscrowedAmount, bal.TradableAmount)
	if err != nil {
		return err
	}

	return k.coreStore.BatchBalanceTable().Update(ctx, bal)
}
