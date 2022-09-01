package utils

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/base"
)

// GetCreditTypeFromBatchDenom extracts the classID from a batch denom string, then retrieves it from the params.
func GetCreditTypeFromBatchDenom(ctx context.Context, store api.StateStore, denom string) (*api.CreditType, error) {
	classID := base.GetClassIDFromBatchDenom(denom)
	classInfo, err := store.ClassTable().GetById(ctx, classID)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get class with ID %s: %s", classID, err.Error())
	}
	return store.CreditTypeTable().Get(ctx, classInfo.CreditTypeAbbrev)
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
