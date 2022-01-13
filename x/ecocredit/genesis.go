package ecocredit

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/types/math"
)

// Validate performs basic validation for each credit-batch,
// it returns an error if credit-batch tradable or retired supply
// does not match the sum of all tradable or retired balances
func (s *GenesisState) Validate() error {
	decimalPlaces := make(map[string]uint32)
	calSupplies := make(map[string]math.Dec)
	supplies := make(map[string]math.Dec)
	classIds := make(map[string]string)

	for _, project := range s.ProjectInfo {
		if _, exists := classIds[project.ProjectId]; exists {
			continue
		}
		classIds[project.ProjectId] = project.ClassId
	}

	for _, batch := range s.BatchInfo {
		if _, exists := decimalPlaces[batch.BatchDenom]; exists {
			continue
		}

		for _, class := range s.ClassInfo {
			if classIds[batch.ProjectId] == class.ClassId {
				decimalPlaces[batch.BatchDenom] = class.CreditType.GetPrecision()
				break
			}
		}
	}

	for _, s := range s.Supplies {
		tSupply := math.NewDecFromInt64(0)
		rSupply := math.NewDecFromInt64(0)
		var err error
		if s.TradableSupply != "" {
			tSupply, err = math.NewNonNegativeFixedDecFromString(s.TradableSupply, decimalPlaces[s.BatchDenom])
			if err != nil {
				return err
			}
		}
		if s.RetiredSupply != "" {
			rSupply, err = math.NewNonNegativeFixedDecFromString(s.RetiredSupply, decimalPlaces[s.BatchDenom])
			if err != nil {
				return err
			}
		}

		total, err := math.SafeAddBalance(tSupply, rSupply)
		if err != nil {
			return err
		}

		supplies[s.BatchDenom] = total
	}

	// calculate credit batch supply from genesis tradable and retired balances
	if err := calculateSupply(decimalPlaces, s.Balances, calSupplies); err != nil {
		return err
	}

	if err := validateSupply(calSupplies, supplies); err != nil {
		return err
	}

	// run each params validation method
	if err := s.Params.Validate(); err != nil {
		return err
	}

	// check that the CreditTypes in the ClassInfo slice all exist in params.CreditTypes
	if err := validateClassInfoTypes(s.Params.CreditTypes, s.ClassInfo); err != nil {
		return err
	}

	return nil
}

func validateClassInfoTypes(creditTypes []*CreditType, classInfos []*ClassInfo) error {
	typeMap := make(map[string]CreditType, len(creditTypes))

	// convert to map for easier lookups
	for _, cType := range creditTypes {
		typeMap[cType.Abbreviation] = *cType
	}

	for _, cInfo := range classInfos {
		// fetch param via abbreviation
		cType, ok := typeMap[cInfo.CreditType.Abbreviation]

		// if it's not found, it's an invalid credit type
		if !ok {
			return sdkerrors.ErrNotFound.Wrapf("unknown credit type abbreviation: %s", cInfo.CreditType.Abbreviation)
		}

		// check that the credit types are equal
		if cType != *cInfo.CreditType {
			return sdkerrors.ErrInvalidType.Wrapf("credit type %+v does not match param type %+v", *cInfo.CreditType, cType)
		}
	}
	return nil
}

func validateSupply(calSupply, supply map[string]math.Dec) error {
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

func calculateSupply(decimalPlaces map[string]uint32, balances []*Balance, calSupply map[string]math.Dec) error {
	for _, b := range balances {
		tBalance := math.NewDecFromInt64(0)
		rBalance := math.NewDecFromInt64(0)
		var err error

		if b.TradableBalance != "" {
			tBalance, err = math.NewNonNegativeFixedDecFromString(b.TradableBalance, decimalPlaces[b.BatchDenom])
			if err != nil {
				return err
			}
		}

		if b.RetiredBalance != "" {
			rBalance, err = math.NewNonNegativeFixedDecFromString(b.RetiredBalance, decimalPlaces[b.BatchDenom])
			if err != nil {
				return err
			}
		}

		total, err := math.SafeAddBalance(tBalance, rBalance)
		if err != nil {
			return err
		}

		if supply, ok := calSupply[b.BatchDenom]; ok {
			result, err := math.SafeAddBalance(supply, total)
			if err != nil {
				return err
			}
			calSupply[b.BatchDenom] = result
		} else {
			calSupply[b.BatchDenom] = total
		}
	}

	return nil
}

// DefaultGenesisState returns a default ecocredit module genesis state.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:      DefaultParams(),
		ClassInfo:   []*ClassInfo{},
		BatchInfo:   []*BatchInfo{},
		Sequences:   []*CreditTypeSeq{},
		Balances:    []*Balance{},
		Supplies:    []*Supply{},
		ProjectInfo: []*ProjectInfo{},
	}
}
