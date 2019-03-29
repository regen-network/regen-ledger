package data

import (
	bytes2 "bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/regen-network/regen-ledger/graph"
	"github.com/regen-network/regen-ledger/graph/binary"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/schema"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	//schemaStoreKey  sdk.StoreKey
	dataStoreKey sdk.StoreKey
	schemaKeeper schema.Keeper
	cdc          *codec.Codec // The wire codec for binary encoding/decoding.
}

type DataRecord struct {
	Data        []byte `json:"data"`
	BlockHeight int64  `json:"block_height"`
}

// NewKeeper creates new instances of the nameservice Keeper
func NewKeeper(dataStoreKey sdk.StoreKey, schemaKeeper schema.Keeper, cdc *codec.Codec) Keeper {
	return Keeper{
		dataStoreKey,
		schemaKeeper,
		cdc,
	}
}

func (k Keeper) GetData(ctx sdk.Context, hash []byte) []byte {
	return k.getDataRecord(ctx, hash).Data
}

func (k Keeper) GetDataBlockHeight(ctx sdk.Context, hash []byte) int64 {
	bh := k.getDataRecord(ctx, hash).BlockHeight
	if bh <= 0 {
		return math.MaxInt64
	}
	return bh
}

func (k Keeper) getDataRecord(ctx sdk.Context, hash []byte) (data DataRecord) {
	store := ctx.KVStore(k.dataStoreKey)
	bz := store.Get(hash)
	if bz == nil {
		return
	}
	return k.decodeDataRecord(bz)
}

const (
	gasForHashAndLookup = 100
	gasPerByteStorage   = 100
)

func (k Keeper) StoreGraph(ctx sdk.Context, hash []byte, data []byte) (types.DataAddress, sdk.Error) {
	ctx.GasMeter().ConsumeGas(gasForHashAndLookup, "hash data")
	g, err := binary.DeserializeGraph(binary.NewOnChainSchemaResolver(k.schemaKeeper, ctx), bytes2.NewBuffer(data))
	if err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("error deserializing graph %s", err.Error()))
	}
	hash2 := graph.Hash(g)
	if !bytes2.Equal(hash, hash2) {
		return nil, sdk.ErrUnknownRequest("incorrect graph hash")
	}
	store := ctx.KVStore(k.dataStoreKey)
	existing := k.getDataRecord(ctx, hash)
	if existing.BlockHeight != 0 {
		return nil, sdk.ErrUnknownRequest("already exists")
	}
	bytes := len(data)
	ctx.GasMeter().ConsumeGas(gasPerByteStorage*uint64(bytes), "store data")
	bz := k.encodeDataRecord(DataRecord{
		Data:        data,
		BlockHeight: ctx.BlockHeight(),
	})
	store.Set(hash, bz)
	return hash, nil
}

func (k Keeper) encodeDataRecord(data DataRecord) []byte {
	bz, err := k.cdc.MarshalBinaryBare(data)
	if err != nil {
		panic(err)
	}
	return bz
}

func (k Keeper) decodeDataRecord(bz []byte) (data DataRecord) {
	err := k.cdc.UnmarshalBinaryBare(bz, &data)
	if err != nil {
		panic(err)
	}
	return
}
