package server

import (
	"bytes"
	"encoding/binary"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types/math"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

type supplyKey func(batchDenom batchDenomT) []byte
type balanceKey func(addr sdk.AccAddress, batchDenom batchDenomT) []byte

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

	if err := setBalance(store, genesisState.TradableBalances, func(addr sdk.AccAddress, batchDenom batchDenomT) []byte {
		return TradableBalanceKey(addr, batchDenom)
	}); err != nil {
		return nil, err
	}

	if err := setBalance(store, genesisState.RetiredBalances, func(addr sdk.AccAddress, batchDenom batchDenomT) []byte {
		return RetiredBalanceKey(addr, batchDenom)
	}); err != nil {
		return nil, err
	}

	if err := setSupply(store, genesisState.TradableSupplies, func(bd batchDenomT) []byte {
		return TradableSupplyKey(bd)
	}); err != nil {
		return nil, err
	}

	if err := setSupply(store, genesisState.RetiredSupplies, func(bd batchDenomT) []byte {
		return RetiredSupplyKey(bd)
	}); err != nil {
		return nil, err
	}

	for _, precision := range genesisState.Precisions {
		key := MaxDecimalPlacesKey(batchDenomT(precision.BatchDenom))
		setUInt32(store, key, precision.MaxDecimalPlaces)
	}

	return []abci.ValidatorUpdate{}, nil
}

func setSupply(store sdk.KVStore, supplies []*ecocredit.Supply, keyFunc supplyKey) error {
	for _, supply := range supplies {
		d, err := math.ParseNonNegativeDecimal(supply.Supply)
		if err != nil {
			return err
		}
		setDecimal(store, keyFunc(batchDenomT(supply.BatchDenom)), d)
	}

	return nil
}

func setBalance(store sdk.KVStore, balances []*ecocredit.Balance, keyFunc balanceKey) error {
	for _, balance := range balances {
		addr, err := sdk.AccAddressFromBech32(balance.Address)
		if err != nil {
			return err
		}
		d, err := math.ParseNonNegativeDecimal(balance.Balance)
		if err != nil {
			return err
		}
		setDecimal(store, keyFunc(addr, batchDenomT(balance.BatchDenom)), d)
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

	tradableBalances := s.getBalances(store, TradableBalancePrefix)
	retiredBalances := s.getBalances(store, RetiredBalancePrefix)
	tradableSupplies := s.getSupplies(store, TradableSupplyPrefix)
	retiredSupplies := s.getSupplies(store, RetiredSupplyPrefix)
	precisions := s.getPrecisions(store, MaxDecimalPlacesPrefix)

	gs := &ecocredit.GenesisState{
		Params:           params,
		ClassInfo:        classInfo,
		BatchInfo:        batchInfo,
		IdSeq:            s.idSeq.CurVal(ctx),
		TradableBalances: tradableBalances,
		RetiredBalances:  retiredBalances,
		TradableSupplies: tradableSupplies,
		RetiredSupplies:  retiredSupplies,
		Precisions:       precisions,
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

func (s serverImpl) getBalances(store sdk.KVStore, storeKey byte) []*ecocredit.Balance {
	iter := sdk.KVStorePrefixIterator(store, []byte{storeKey})
	defer iter.Close()
	var tradableBalances []*ecocredit.Balance
	for ; iter.Valid(); iter.Next() {
		addr, denomMetaData := ParseBalanceKey(iter.Key())
		tradableBalances = append(tradableBalances, &ecocredit.Balance{
			Address:    addr.String(),
			BatchDenom: string(denomMetaData),
			Balance:    string(iter.Value()),
		})
	}
	return tradableBalances
}

func (s serverImpl) getSupplies(store sdk.KVStore, storeKey byte) []*ecocredit.Supply {
	iter := sdk.KVStorePrefixIterator(store, []byte{storeKey})
	defer iter.Close()
	var retiredSupplies []*ecocredit.Supply
	for ; iter.Valid(); iter.Next() {
		denomMetaData := ParseSupplyKey(iter.Key())
		retiredSupplies = append(retiredSupplies, &ecocredit.Supply{
			BatchDenom: string(denomMetaData),
			Supply:     string(iter.Value()),
		})
	}
	return retiredSupplies
}
