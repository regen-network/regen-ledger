package ecocredit

import (
	apd "github.com/cockroachdb/apd/v2"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/types/math"
)

// Validate performs basic validation for each credit-batch,
// it returns an error if credit-batch tradable or retired supply
// does not match the sum of all tradable or retired balances
func (s *GenesisState) Validate() error {
	decimalPlaces := make(map[string]uint32)
	calSupplies := make(map[string]*apd.Decimal)
	supplies := make(map[string]*apd.Decimal)

	for _, precision := range s.Precisions {
		decimalPlaces[precision.BatchDenom] = precision.MaxDecimalPlaces
	}

	for _, s := range s.Supplies {
		supply, err := math.ParsePositiveFixedDecimal(s.Supply, decimalPlaces[s.BatchDenom])
		if err != nil {
			return err
		}
		supplies[s.BatchDenom] = supply
	}

	if err := calculateSupply(decimalPlaces, s.Balances, calSupplies); err != nil {
		return err
	}

	if err := validateSupply(calSupplies, supplies); err != nil {
		return err
	}

	return nil
}

func validateSupply(calSupply, supply map[string]*apd.Decimal) error {
	for denom, cs := range calSupply {
		if s, ok := supply[denom]; ok {
			if s.Cmp(cs) != 0 {
				return sdkerrors.ErrInvalidCoins.Wrapf("supply is incorrect for %s credit batch, expected %v, got %v", denom, s, cs)
			}
		} else {
			return sdkerrors.ErrNotFound.Wrapf("supply is not found for %s credit batch", denom)
		}
	}
	return nil
}

func calculateSupply(decimalPlaces map[string]uint32, balances []*Balance, calSupply map[string]*apd.Decimal) error {
	for _, b := range balances {
		if b.Type == Balance_TYPE_UNSPECIFIED {
			return sdkerrors.ErrInvalidType.Wrapf("expecting %v or %v, got %v", Balance_TYPE_TRADABLE, Balance_TYPE_RETIRED, Balance_TYPE_UNSPECIFIED)
		}
		balance, err := math.ParsePositiveFixedDecimal(b.Balance, decimalPlaces[b.BatchDenom])
		if err != nil {
			return err
		}

		if supply, ok := calSupply[b.BatchDenom]; ok {
			if err := math.Add(supply, supply, balance); err != nil {
				return err
			}
			calSupply[b.BatchDenom] = supply
		} else {
			calSupply[b.BatchDenom] = balance
		}
	}
	return nil
}

// DefaultGenesisState returns a default ecocredit module genesis state.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:     DefaultParams(),
		ClassInfo:  []*ClassInfo{},
		BatchInfo:  []*BatchInfo{},
		IdSeq:      0,
		Balances:   []*Balance{},
		Supplies:   []*Supply{},
		Precisions: []*Precision{},
	}
}
