package core

import (
	"context"

	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// assertClassIssuer makes sure that the issuer is part of issuers of given classID.
// Returns ErrUnauthorized otherwise.
func (k Keeper) assertClassIssuer(goCtx context.Context, classID uint64, addr sdk.AccAddress) error {
	found, err := k.stateStore.ClassIssuerTable().Has(goCtx, classID, addr)
	if err != nil {
		return err
	}
	if !found {
		return sdkerrors.ErrUnauthorized.Wrapf("%s is not an issuer for the class", addr)
	}
	return nil
}

// AddAndSaveBalance adds 'amt' to the addr's tradable balance.
func AddAndSaveBalance(ctx context.Context, table ecoApi.BatchBalanceTable, addr sdk.AccAddress, batchKey uint64, amt math.Dec) error {
	bal, err := utils.GetBalance(ctx, table, addr, batchKey)
	if err != nil {
		return err
	}
	tradable, err := math.NewDecFromString(bal.Tradable)
	if err != nil {
		return err
	}
	newTradable, err := math.SafeAddBalance(tradable, amt)
	if err != nil {
		return err
	}
	bal.Tradable = newTradable.String()
	return table.Save(ctx, bal)
}

// RetireAndSaveBalance adds 'amt' to the addr's retired balance.
func RetireAndSaveBalance(ctx context.Context, table ecoApi.BatchBalanceTable, addr sdk.AccAddress, batchKey uint64, amount math.Dec) error {
	bal, err := table.Get(ctx, addr, batchKey)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			bal = &ecoApi.BatchBalance{
				BatchKey: batchKey,
				Address:  addr,
				Tradable: "0",
				Retired:  "0",
				Escrowed: "0",
			}
		} else {
			return err
		}
	}
	retired, err := math.NewDecFromString(bal.Retired)
	if err != nil {
		return err
	}
	newRetired, err := math.SafeAddBalance(retired, amount)
	if err != nil {
		return err
	}
	bal.Retired = newRetired.String()
	return table.Save(ctx, bal)
}

// RetireSupply moves `amount` of credits from the supply's tradable amount to its retired amount.
func RetireSupply(ctx context.Context, table ecoApi.BatchSupplyTable, batchKey uint64, amount math.Dec) error {
	supply, err := table.Get(ctx, batchKey)
	if err != nil {
		return err
	}
	tradable, err := math.NewDecFromString(supply.TradableAmount)
	if err != nil {
		return err
	}

	retired, err := math.NewDecFromString(supply.RetiredAmount)
	if err != nil {
		return err
	}

	tradable, err = math.SafeSubBalance(tradable, amount)
	if err != nil {
		return err
	}

	retired, err = math.SafeAddBalance(retired, amount)
	if err != nil {
		return err
	}

	supply.TradableAmount = tradable.String()
	supply.RetiredAmount = retired.String()

	return table.Update(ctx, supply)
}
