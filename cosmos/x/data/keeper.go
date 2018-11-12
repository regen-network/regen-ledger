package data

import (
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/gojsonschema"
	"golang.org/x/crypto/blake2b"
	"encoding/base64"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	schemaStoreKey  sdk.StoreKey
	dataStoreKey  sdk.StoreKey

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the nameservice Keeper
func NewKeeper(schemaStoreKey sdk.StoreKey, dataStoreKey  sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		schemaStoreKey: schemaStoreKey,
		dataStoreKey: dataStoreKey,
		cdc:            cdc,
	}
}

func (k Keeper) GetSchema(ctx sdk.Context, id string) (*gojsonschema.Schema, error) {
	store := ctx.KVStore(k.schemaStoreKey)
	bz := store.Get([]byte(id))
	loader := gojsonschema.NewStringLoader(string(bz))
	sl := gojsonschema.NewSchemaLoader()
	schema, err := sl.Compile(loader)

	if err != nil {
		return nil, err
	}
	return schema, nil
}

func (k Keeper) RegisterSchema(ctx sdk.Context, schema string) string {
	store := ctx.KVStore(k.schemaStoreKey)
	hash := blake2b.Sum256([]byte(schema))
	id := base64.URLEncoding.EncodeToString(hash[:])
	store.Set([]byte(id), []byte(schema))
	return id
}
