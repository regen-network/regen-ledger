package utils

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/v2/math"
)

// GetCreditTypeFromBatch extracts the credit type from a batch by querying the class table for the class key, and then
// querying the credit type table for the credit type abbreviation. This is a convenience function for use in message
// handlers, where the credit type is needed to perform further operations on the batch.
func GetCreditTypeFromBatch(ctx context.Context, store api.StateStore, batch *api.Batch) (*api.CreditType, error) {
	cls, err := store.ClassTable().Get(ctx, batch.ClassKey)
	if err != nil {
		return nil, err
	}

	return store.CreditTypeTable().Get(ctx, cls.CreditTypeAbbrev)
}

// GetNonNegativeFixedDecs takes an arbitrary amount of decimal strings, and returns their corresponding fixed decimals
// in a slice.
func GetNonNegativeFixedDecs(precision uint32, decimals ...string) ([]math.Dec, error) {
	decs := make([]math.Dec, len(decimals))
	for i, decimal := range decimals {
		dec, err := math.NewNonNegativeFixedDecFromString(decimal, precision)
		if err != nil {
			return nil, err
		}
		decs[i] = dec
	}
	return decs, nil
}

// GetBalance gets the balance from the account, returning a default, zero value balance if no balance is found.
// NOTE: the default value is not inserted into the balance table in the `not found` case. Calling Update when the default
// value is returned will cause an error. The `Save` method should be used when dealing with balances from this function.
func GetBalance(ctx context.Context, table api.BatchBalanceTable, addr types.AccAddress, key uint64) (*api.BatchBalance, error) {
	bal, err := table.Get(ctx, addr, key)
	if err != nil {
		if !ormerrors.IsNotFound(err) {
			return nil, err
		}
		bal = &api.BatchBalance{
			BatchKey:       key,
			Address:        addr,
			TradableAmount: "0",
			RetiredAmount:  "0",
			EscrowedAmount: "0",
		}
	}
	return bal, nil
}
