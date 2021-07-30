package server

import (
	"bytes"
	"encoding/binary"
	"encoding/json"

	"github.com/cockroachdb/apd/v2"
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

	s.idSeq.InitVal(ctx, genesisState.IdSeq)

	s.paramSpace.SetParamSet(ctx.Context, &genesisState.Params)

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

	for _, precision := range genesisState.Precisions {
		key := MaxDecimalPlacesKey(batchDenomT(precision.BatchDenom))
		setUInt32(store, key, precision.MaxDecimalPlaces)
	}

	return []abci.ValidatorUpdate{}, nil
}

// validateSupplies returns an error if credit batch genesis supply does not equal to calculated supply.
func validateSupplies(store sdk.KVStore, supplies []*ecocredit.Supply) error {
	var denomT batchDenomT
	actual := apd.New(0, 0)
	for _, supply := range supplies {
		genesisSupply, err := math.ParseNonNegativeDecimal(supply.Supply)
		if err != nil {
			return err
		}

		denomT = batchDenomT(supply.BatchDenom)
		tradable, err := getDecimal(store, TradableSupplyKey(denomT))
		if err != nil {
			return err
		}
		retired, err := getDecimal(store, RetiredSupplyKey(denomT))
		if err != nil {
			return err
		}

		if err := math.Add(actual, tradable, retired); err != nil {
			return err
		}

		if actual.Cmp(genesisSupply) != 0 {
			return sdkerrors.ErrInvalidCoins.Wrapf("supply is incorrect for %s credit batch, expected %v, got %v", supply.BatchDenom, actual, genesisSupply)
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
		d, err := math.ParseNonNegativeDecimal(balance.Balance)
		if err != nil {
			return err
		}
		denomT := batchDenomT(balance.BatchDenom)
		balanceKey, err := getBalanceKey(balance.Type, addr, denomT)
		if err != nil {
			return err
		}
		setDecimal(store, balanceKey, d)

		supplyKey, err := getSupplyKey(balance.Type, denomT)
		if err != nil {
			return err
		}

		getAddAndSetDecimal(store, supplyKey, d)
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

	var balances []*ecocredit.Balance
	iterateBalances(store, TradableBalancePrefix, func(address, denom, balance string) bool {
		balances = append(balances, &ecocredit.Balance{
			Address:    address,
			BatchDenom: denom,
			Balance:    balance,
			Type:       ecocredit.Balance_TYPE_TRADABLE,
		})
		return false
	})

	iterateBalances(store, RetiredBalancePrefix, func(address, denom, balance string) bool {
		balances = append(balances, &ecocredit.Balance{
			Address:    address,
			BatchDenom: denom,
			Balance:    balance,
			Type:       ecocredit.Balance_TYPE_RETIRED,
		})
		return false
	})

	suppliesMap := make(map[string]*apd.Decimal)

	iterateSupplies(store, TradableSupplyPrefix, func(denom, value string) (bool, error) {
		supply, err := math.ParseNonNegativeDecimal(value)
		if err != nil {
			panic(err)
		}
		if _, exists := suppliesMap[denom]; exists {
			return true, sdkerrors.ErrConflict.Wrapf("duplicate tradable supply found: denom %s", denom)
		}
		suppliesMap[denom] = supply

		return false, nil
	})

	iterateSupplies(store, RetiredSupplyPrefix, func(denom, value string) (bool, error) {
		supply := apd.New(0, 0)
		rSupply, err := math.ParseNonNegativeDecimal(value)
		if err != nil {
			return true, err
		}
		if tSupply, exists := suppliesMap[denom]; exists {
			math.Add(supply, tSupply, rSupply)
			suppliesMap[denom] = supply
		} else {
			suppliesMap[denom] = rSupply
		}

		return false, nil
	})

	supplies := make([]*ecocredit.Supply, len(suppliesMap))
	i := 0
	for denom, supply := range suppliesMap {
		supplies[i] = &ecocredit.Supply{
			BatchDenom: denom,
			Supply:     supply.String(),
		}
		i++
	}

	precisions := s.getPrecisions(store, MaxDecimalPlacesPrefix)

	gs := &ecocredit.GenesisState{
		Params:     params,
		ClassInfo:  classInfo,
		BatchInfo:  batchInfo,
		IdSeq:      s.idSeq.CurVal(ctx),
		Balances:   balances,
		Supplies:   supplies,
		Precisions: precisions,
	}

	return cdc.MustMarshalJSON(gs), nil
}

func (s serverImpl) getPrecisions(store sdk.KVStore, storeKey byte) []*ecocredit.Precision {
	iter := sdk.KVStorePrefixIterator(store, []byte{storeKey})
	defer iter.Close()
	var precisions []*ecocredit.Precision
	for ; iter.Valid(); iter.Next() {
		denomMetaData := ParseMaxDecimalPlacesKey(iter.Key())

		buf := bytes.NewReader(iter.Value())
		var val uint32
		err := binary.Read(buf, binary.LittleEndian, &val)
		if err != nil {
			panic(err)
		}

		precisions = append(precisions, &ecocredit.Precision{
			BatchDenom:       string(denomMetaData),
			MaxDecimalPlaces: val,
		})
	}
	return precisions
}
