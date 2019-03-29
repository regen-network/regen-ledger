package data

import (
	bytes2 "bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/graph"
	"github.com/regen-network/regen-ledger/graph/binary"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/schema"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	//schemaStoreKey  sdk.StoreKey
	dataStoreKey sdk.StoreKey
	schemaKeeper schema.Keeper
	cdc          *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the nameservice Keeper
func NewKeeper(dataStoreKey sdk.StoreKey, schemaKeeper schema.Keeper, cdc *codec.Codec) Keeper {
	return Keeper{
		dataStoreKey,
		schemaKeeper,
		cdc,
	}
}

// GetData returns the data if it exists or returns an error
func (k Keeper) GetData(ctx sdk.Context, addr types.DataAddress) ([]byte, sdk.Error) {
	store := ctx.KVStore(k.dataStoreKey)
	bz := store.Get(addr)
	if bz == nil || len(bz) < 1 {
		return nil, sdk.ErrUnknownRequest("not found")
	}
	switch addr[0] {
	case types.DataAddressPrefixOnChainGraph:
		return bz, nil
	default:
		return nil, sdk.ErrUnknownRequest("bad address")
	}
}

const (
	gasForHashAndLookup = 100
	gasPerByteStorage   = 100
)

// StoreGraph stores a graph with the binary representation data and the provided hash
func (k Keeper) StoreGraph(ctx sdk.Context, hash []byte, data []byte) (types.DataAddress, sdk.Error) {
	ctx.GasMeter().ConsumeGas(gasForHashAndLookup, "hash data")
	g, err := binary.DeserializeGraph(schema.NewOnChainSchemaResolver(k.schemaKeeper, ctx), bytes2.NewBuffer(data))
	if err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("error deserializing graph %s", err.Error()))
	}
	hash2 := graph.Hash(g)
	if !bytes2.Equal(hash, hash2) {
		return nil, sdk.ErrUnknownRequest("incorrect graph hash")
	}
	store := ctx.KVStore(k.dataStoreKey)
	addr := types.GetDataAddressOnChainGraph(hash)
	existing, err := k.GetData(ctx, addr)
	if err == nil && existing != nil {
		return nil, sdk.ErrUnknownRequest("already exists")
	}
	bytes := len(data)
	ctx.GasMeter().ConsumeGas(gasPerByteStorage*uint64(bytes), "store data")
	store.Set(addr, data)
	return addr, nil
}
