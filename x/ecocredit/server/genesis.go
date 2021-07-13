package server

import (
	"bytes"
	"encoding/binary"
	"encoding/json"

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

	if err := setBalanceAndSupply(store, genesisState.TradableBalances,
		func(addr sdk.AccAddress, batchDenom batchDenomT) []byte {
			return TradableBalanceKey(addr, batchDenom)
		},
		func(batchDenom batchDenomT) []byte {
			return TradableSupplyKey(batchDenom)
		}); err != nil {
		return nil, err
	}

	if err := validateSupply(store, genesisState.TradableSupplies, func(denom batchDenomT) []byte {
		return TradableSupplyKey(denom)
	}); err != nil {
		return nil, err
	}

	if err := setBalanceAndSupply(store, genesisState.RetiredBalances,
		func(addr sdk.AccAddress, batchDenom batchDenomT) []byte {
			return RetiredBalanceKey(addr, batchDenom)
		},
		func(batchDenom batchDenomT) []byte {
			return RetiredSupplyKey(batchDenom)
		}); err != nil {
		return nil, err
	}

	if err := validateSupply(store, genesisState.RetiredSupplies, func(denom batchDenomT) []byte {
		return RetiredSupplyKey(denom)
	}); err != nil {
		return nil, err
	}

	for _, precision := range genesisState.Precisions {
		key := MaxDecimalPlacesKey(batchDenomT(precision.BatchDenom))
		setUInt32(store, key, precision.MaxDecimalPlaces)
	}

	return []abci.ValidatorUpdate{}, nil
}

// validateSupply returns an error if credit batch genesis tradable or retired supply does not equals to calculated supply.
func validateSupply(store sdk.KVStore, supplies []*ecocredit.Supply, keyFunc supplyKey) error {
	for _, supply := range supplies {
		parsedSupply, err := math.ParseNonNegativeDecimal(supply.Supply)
		if err != nil {
			return err
		}

		s, err := getDecimal(store, keyFunc(batchDenomT(supply.BatchDenom)))
		if err != nil {
			return err
		}

		if s.Cmp(parsedSupply) != 0 {
			return sdkerrors.ErrInvalidCoins.Wrapf("supply is incorrect for %s credit batch, expected %v, got %v", supply.BatchDenom, s, parsedSupply)
		}
	}

	return nil
}

// setBalanceAndSupply sets the tradable or retired balance for an account and update tradable or retired supply for batch denom.
func setBalanceAndSupply(store sdk.KVStore, balances []*ecocredit.Balance, balanceKeyFunc balanceKey,
	supplyKeyFunc supplyKey) error {
	for _, balance := range balances {
		addr, err := sdk.AccAddressFromBech32(balance.Address)
		if err != nil {
			return err
		}
		d, err := math.ParseNonNegativeDecimal(balance.Balance)
		if err != nil {
			return err
		}
		setDecimal(store, balanceKeyFunc(addr, batchDenomT(balance.BatchDenom)), d)

		getAddAndSetDecimal(store, supplyKeyFunc(batchDenomT(balance.BatchDenom)), d)
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
