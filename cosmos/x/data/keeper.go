package data

import (
	"encoding/base64"
	"github.com/cosmos/cosmos-sdk/codec"
	"golang.org/x/crypto/blake2b"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	//schemaStoreKey  sdk.StoreKey
	dataStoreKey sdk.StoreKey

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

type DataRecord struct {
	Data        []byte
	BlockHeight int64
}

// NewKeeper creates new instances of the nameservice Keeper
func NewKeeper(
//schemaStoreKey sdk.StoreKey,
	dataStoreKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		//schemaStoreKey: schemaStoreKey,
		dataStoreKey: dataStoreKey,
		cdc:          cdc,
	}
}

//func (k Keeper) GetSchema(ctx sdk.Context, id string) (*gojsonschema.Schema, error) {
//	store := ctx.KVStore(k.schemaStoreKey)
//	bz := store.Get([]byte(id))
//	loader := gojsonschema.NewStringLoader(string(bz))
//	sl := gojsonschema.NewSchemaLoader()
//	schema, err := sl.Compile(loader)
//
//	if err != nil {
//		return nil, err
//	}
//	return schema, nil
//}
//
//func (k Keeper) RegisterSchema(ctx sdk.Context, schema string) string {
//	store := ctx.KVStore(k.schemaStoreKey)
//	hash := blake2b.Sum256([]byte(schema))
//	id := base64.URLEncoding.EncodeToString(hash[:])
//	store.Set([]byte(id), []byte(schema))
//	return id
//}

func (k Keeper) GetData(ctx sdk.Context, id string) []byte {
	return k.getDataRecord(ctx, id).Data
}

func (k Keeper) GetDataBlockHeight(ctx sdk.Context, id string) int64 {
	bh := k.getDataRecord(ctx, id).BlockHeight
	if bh <= 0 {
		return math.MaxInt64
	}
	return bh
}

func (k Keeper) getDataRecord(ctx sdk.Context, id string) (data DataRecord) {
	store := ctx.KVStore(k.dataStoreKey)
	bz := store.Get([]byte(id))
	if bz == nil {
		return
	}
	return k.decodeDataRecord(bz)
}

const (
	gasForHashAndLookup = 100
	gasPerByteStorage   = 100
)

func (k Keeper) StoreData(ctx sdk.Context, data []byte) string {
	ctx.GasMeter().ConsumeGas(gasForHashAndLookup, "hash data")
	store := ctx.KVStore(k.dataStoreKey)
	hash := blake2b.Sum256([]byte(data))
	id := base64.URLEncoding.EncodeToString(hash[:])
	existing := k.getDataRecord(ctx, id)
	if existing.BlockHeight != 0 {
		return id
	}
	bytes := len(data)
	ctx.GasMeter().ConsumeGas(gasPerByteStorage*uint64(bytes), "store data")
	bz := k.encodeDataRecord(DataRecord{
		Data:data,
		BlockHeight:ctx.BlockHeight(),
	})
	store.Set([]byte(id), bz)
	return id
}

//func (k Keeper) GetDataPointer(ctx sdk.Context, id string) string {
//	store := ctx.KVStore(k.dataStoreKey)
//	bz := store.Get([]byte(id))
//	return string(bz)
//}
//
//func (k Keeper) PutDataPointer(ctx sdk.Context, data string) string {
//	store := ctx.KVStore(k.dataStoreKey)
//	hash := blake2b.Sum256([]byte(data))
//	id := base64.URLEncoding.EncodeToString(hash[:])
//	store.Set([]byte(id), []byte(data))
//	return id
//}

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
