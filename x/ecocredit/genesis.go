package ecocredit

import (
	apd "github.com/cockroachdb/apd/v2"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/types/math"
)

// ValidateGenesis check the given genesis state has no integrity issues.
func (s *GenesisState) Validate() error {
	decimalPlaces := make(map[string]uint32)
	calTradableSupply := make(map[string]*apd.Decimal)
	calRetiredSupply := make(map[string]*apd.Decimal)
	tradableSupply := make(map[string]*apd.Decimal)
	retiredSupply := make(map[string]*apd.Decimal)

	for _, precision := range s.Precisions {
		decimalPlaces[precision.BatchDenom] = precision.MaxDecimalPlaces
	}

	for _, tSupply := range s.TradableSupplies {
		supply, err := math.ParsePositiveFixedDecimal(tSupply.Supply, decimalPlaces[tSupply.BatchDenom])
		if err != nil {
			return err
		}
		tradableSupply[tSupply.BatchDenom] = supply
	}

	for _, rSupply := range s.RetiredSupplies {
		supply, err := math.ParsePositiveFixedDecimal(rSupply.Supply, decimalPlaces[rSupply.BatchDenom])
		if err != nil {
			return err
		}
		retiredSupply[rSupply.BatchDenom] = supply
	}

	for _, tBalance := range s.TradableBalances {
		balance, err := math.ParsePositiveFixedDecimal(tBalance.Balance, decimalPlaces[tBalance.BatchDenom])
		if err != nil {
			return err
		}

		if supply, ok := calTradableSupply[tBalance.BatchDenom]; ok {
			if err := math.Add(supply, supply, balance); err != nil {
				return err
			}
			calTradableSupply[tBalance.BatchDenom] = supply
		} else {
			calTradableSupply[tBalance.BatchDenom] = balance
		}
	}

	for _, rBalance := range s.RetiredBalances {
		balance, err := math.ParsePositiveFixedDecimal(rBalance.Balance, decimalPlaces[rBalance.BatchDenom])
		if err != nil {
			return err
		}

		if supply, ok := calRetiredSupply[rBalance.BatchDenom]; ok {
			if err := math.Add(supply, supply, balance); err != nil {
				return err
			}
			calRetiredSupply[rBalance.BatchDenom] = supply
		} else {
			calRetiredSupply[rBalance.BatchDenom] = balance
		}
	}

	for denom, calSupply := range calTradableSupply {
		if supply, ok := tradableSupply[denom]; ok {
			if supply.Cmp(calSupply) != 0 {
				return sdkerrors.ErrInvalidCoins.Wrapf("tradable: supply is incorrect for %s credit batch, expected %v, got %v", denom, supply, calSupply)
			}
		} else {
			return sdkerrors.ErrNotFound.Wrapf("tradable: supply is not found for %s credit batch", denom)
		}
	}

	for denom, calSupply := range calRetiredSupply {
		if supply, ok := retiredSupply[denom]; ok {
			if supply.Cmp(calSupply) != 0 {
				return sdkerrors.ErrInvalidCoins.Wrapf("retired: supply is incorrect for %s credit batch, expected %v, got %v", denom, supply, calSupply)
			}
		} else {
			return sdkerrors.ErrNotFound.Wrapf("retired: supply is not found for %s credit batch", denom)
		}
	}

	return nil
}

// DefaultGenesisState returns a default ecocredit module genesis state.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:           DefaultParams(),
		ClassInfo:        []*ClassInfo{},
		BatchInfo:        []*BatchInfo{},
		IdSeq:            0,
		TradableBalances: []*Balance{},
		RetiredBalances:  []*Balance{},
		TradableSupplies: []*Supply{},
		RetiredSupplies:  []*Supply{},
		Precisions:       []*Precision{},
	}
}
