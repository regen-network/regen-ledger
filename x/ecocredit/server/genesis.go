package server

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pkg/errors"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types/math"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// InitGenesis performs genesis initialization for the ecocredit module. It
// returns no validator updates.
func (s serverImpl) InitGenesis(ctx types.Context, cdc codec.Codec, data json.RawMessage) ([]abci.ValidatorUpdate, error) {
	var genesisState ecocredit.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)

	s.paramSpace.SetParamSet(ctx.Context, &genesisState.Params)

	if err := orm.ImportTableData(ctx, s.creditTypeSeqTable, genesisState.Sequences, 0); err != nil {
		return nil, errors.Wrap(err, "sequences")
	}

	if err := orm.ImportTableData(ctx, s.classInfoTable, genesisState.ClassInfo, 0); err != nil {
		return nil, errors.Wrap(err, "class-info")
	}

	if err := orm.ImportTableData(ctx, s.batchInfoTable, genesisState.BatchInfo, 0); err != nil {
		return nil, errors.Wrap(err, "batch-info")
	}

	store := ctx.KVStore(s.storeKey)
	if err := setBalanceAndSupply(store, genesisState.Balances); err != nil {
		return nil, err
	}

	if err := validateSupplies(store, genesisState.Supplies); err != nil {
		return nil, err
	}

	return []abci.ValidatorUpdate{}, nil
}

// validateSupplies returns an error if credit batch genesis supply does not equal to calculated supply.
func validateSupplies(store sdk.KVStore, supplies []*ecocredit.Supply) error {
	var denomT batchDenomT
	for _, supply := range supplies {
		denomT = batchDenomT(supply.BatchDenom)
		tradableSupply := math.NewDecFromInt64(0)
		retiredSupply := math.NewDecFromInt64(0)
		var err error
		if supply.TradableSupply != "" {
			tradableSupply, err = math.NewNonNegativeDecFromString(supply.TradableSupply)
			if err != nil {
				return err
			}
		}

		tradable, err := getDecimal(store, TradableSupplyKey(denomT))
		if err != nil {
			return err
		}

		if tradableSupply.Cmp(tradable) != 0 {
			return sdkerrors.ErrInvalidCoins.Wrapf("tradable supply is incorrect for %s credit batch, expected %v, got %v", supply.BatchDenom, tradable, tradableSupply)
		}

		if supply.RetiredSupply != "" {
			retiredSupply, err = math.NewNonNegativeDecFromString(supply.RetiredSupply)
			if err != nil {
				return err
			}
		}

		retired, err := getDecimal(store, RetiredSupplyKey(denomT))
		if err != nil {
			return err
		}

		if retiredSupply.Cmp(retired) != 0 {
			return sdkerrors.ErrInvalidCoins.Wrapf("retired supply is incorrect for %s credit batch, expected %v, got %v", supply.BatchDenom, retired, retiredSupply)
		}
	}

	return nil
}

// setBalanceAndSupply sets the tradable and retired balance for an account and update supply for batch denom.
func setBalanceAndSupply(store sdk.KVStore, balances []*ecocredit.Balance) error {
	for _, balance := range balances {
		addr, err := sdk.AccAddressFromBech32(balance.Address)
		if err != nil {
			return err
		}
		denomT := batchDenomT(balance.BatchDenom)

		// set tradable balance and update supply
		if balance.TradableBalance != "" {
			d, err := math.NewNonNegativeDecFromString(balance.TradableBalance)
			if err != nil {
				return err
			}
			key := TradableBalanceKey(addr, denomT)
			setDecimal(store, key, d)

			key = TradableSupplyKey(denomT)
			getAddAndSetDecimal(store, key, d)
		}

		// set retired balance and update supply
		if balance.RetiredBalance != "" {
			d, err := math.NewNonNegativeDecFromString(balance.RetiredBalance)
			if err != nil {
				return err
			}
			key := RetiredBalanceKey(addr, denomT)
			setDecimal(store, key, d)

			key = RetiredSupplyKey(denomT)
			getAddAndSetDecimal(store, key, d)
		}
	}

	return nil
}

// ExportGenesis will dump the ecocredit module state into a serializable GenesisState.
func (s serverImpl) ExportGenesis(ctx types.Context, cdc codec.Codec) (json.RawMessage, error) {
	// Get Params from the store and put them in the genesis state
	var params ecocredit.Params
	s.paramSpace.GetParamSet(ctx.Context, &params)

	store := ctx.KVStore(s.storeKey)
	var classInfo []*ecocredit.ClassInfo
	if _, err := orm.ExportTableData(ctx, s.classInfoTable, &classInfo); err != nil {
		return nil, errors.Wrap(err, "class-info")
	}

	var batchInfo []*ecocredit.BatchInfo
	if _, err := orm.ExportTableData(ctx, s.batchInfoTable, &batchInfo); err != nil {
		return nil, errors.Wrap(err, "batch-info")
	}

	var sequences []*ecocredit.CreditTypeSeq
	if _, err := orm.ExportTableData(ctx, s.creditTypeSeqTable, &sequences); err != nil {
		return nil, errors.Wrap(err, "batch-info")
	}

	suppliesMap := make(map[string]*ecocredit.Supply)
	iterateSupplies(store, TradableSupplyPrefix, func(denom, supply string) (bool, error) {
		suppliesMap[denom] = &ecocredit.Supply{
			BatchDenom:     denom,
			TradableSupply: supply,
		}

		return false, nil
	})

	iterateSupplies(store, RetiredSupplyPrefix, func(denom, supply string) (bool, error) {
		if _, exists := suppliesMap[denom]; exists {
			suppliesMap[denom].RetiredSupply = supply
		} else {
			suppliesMap[denom] = &ecocredit.Supply{
				BatchDenom:    denom,
				RetiredSupply: supply,
			}
		}

		return false, nil
	})

	supplies := make([]*ecocredit.Supply, len(suppliesMap))
	index := 0
	for _, supply := range suppliesMap {
		supplies[index] = supply
		index++
	}

	balancesMap := make(map[string]*ecocredit.Balance)
	iterateBalances(store, TradableBalancePrefix, func(address, denom, balance string) bool {
		balancesMap[fmt.Sprintf("%s%s", address, denom)] = &ecocredit.Balance{
			Address:         address,
			BatchDenom:      denom,
			TradableBalance: balance,
		}

		return false
	})

	iterateBalances(store, RetiredBalancePrefix, func(address, denom, balance string) bool {
		index := fmt.Sprintf("%s%s", address, denom)
		if _, exists := balancesMap[index]; exists {
			balancesMap[index].RetiredBalance = balance
		} else {
			balancesMap[index] = &ecocredit.Balance{
				Address:        address,
				BatchDenom:     denom,
				RetiredBalance: balance,
			}
		}

		return false
	})

	balances := make([]*ecocredit.Balance, len(balancesMap))
	index = 0
	for _, balance := range balancesMap {
		balances[index] = balance
		index++
	}

	gs := &ecocredit.GenesisState{
		Params:    params,
		ClassInfo: classInfo,
		BatchInfo: batchInfo,
		Sequences: sequences,
		Balances:  balances,
		Supplies:  supplies,
	}

	return cdc.MustMarshalJSON(gs), nil
}
