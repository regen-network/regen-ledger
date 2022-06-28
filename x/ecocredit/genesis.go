package ecocredit

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	dbm "github.com/tendermint/tm-db"

	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types/math"
)

// Validate performs basic validation for each credit-batch,
// it returns an error if credit-batch tradable or retired supply
// does not match the sum of all tradable or retired balances
func (s *GenesisState) Validate(bz json.RawMessage) error {
	decimalPlaces := make(map[string]uint32)
	calSupplies := make(map[string]math.Dec)
	supplies := make(map[string]math.Dec)

	for _, batch := range s.BatchInfo {
		if _, exists := decimalPlaces[batch.BatchDenom]; exists {
			continue
		}

		for _, class := range s.ClassInfo {
			if batch.ClassId == class.ClassId {
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

	db := dbm.NewMemDB()
	backend := ormtable.NewBackend(ormtable.BackendOptions{
		CommitmentStore: db,
		IndexStore:      db,
	})

	db1, err := ormdb.NewModuleDB(ModuleSchema, ormdb.ModuleDBOptions{})
	if err != nil {
		return err
	}

	ormCtx := ormtable.WrapContextDefault(backend)
	basketStore, err := basketapi.NewStateStore(db1)
	if err != nil {
		return err
	}

	jsonSource, err := ormjson.NewRawMessageSource(bz)
	if err != nil {
		return err
	}

	err = db1.ImportJSON(ormCtx, jsonSource)
	if err != nil {
		return err
	}

	if err := db1.ValidateJSON(jsonSource); err != nil {
		return err
	}

	// calculate credit batch supply from genesis tradable and retired balances
	if err := calculateSupply(decimalPlaces, s.Balances, calSupplies); err != nil {
		return err
	}

	// add basket balances to credit batch supply
	bBalanceItr, err := basketStore.BasketBalanceStore().List(ormCtx, basketapi.BasketBalancePrimaryKey{})
	if err != nil {
		return err
	}
	defer bBalanceItr.Close()

	for bBalanceItr.Next() {
		bb, err := bBalanceItr.Value()
		if err != nil {
			return err
		}

		batchSupply, ok := calSupplies[bb.BatchDenom]
		if !ok {
			return fmt.Errorf("unknown credit batch %s in basket", bb.BatchDenom)
		}

		balance, err := math.NewDecFromString(bb.Balance)
		if err != nil {
			return err
		}

		result, err := math.Add(batchSupply, balance)
		if err != nil {
			return err
		}

		calSupplies[bb.BatchDenom] = result
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
		Params:    DefaultParams(),
		ClassInfo: []*ClassInfo{},
		BatchInfo: []*BatchInfo{},
		Sequences: []*CreditTypeSeq{},
		Balances:  []*Balance{},
		Supplies:  []*Supply{},
	}
}
