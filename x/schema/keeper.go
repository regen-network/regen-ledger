package schema

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/regen-network/regen-ledger/graph"
	"net/url"
)
import sdk "github.com/cosmos/cosmos-sdk/types"

// Keeper is the schema module keeper
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

var (
	LastPropertyIDKey = []byte("_/next-property-id")
)

// NewKeeper creates a new Keeper
func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{storeKey: storeKey, cdc: cdc}
}

// PropertyKey returns the store key for the given graph.PropertyID that points to its PropertyDefinition
func PropertyKey(id graph.PropertyID) []byte {
	return []byte(fmt.Sprintf("p/%d", id))
}

// PropertyURLKey returns the store key for the given property URL that points to its graph.PropertyID
func PropertyURLKey(url string) []byte {
	return []byte(fmt.Sprintf("u/%s", url))
}

// GetProperty returns a PropertyDefinition given a graph.PropertyID if one exists
func (keeper Keeper) GetPropertyDefinition(ctx sdk.Context, id graph.PropertyID) (prop PropertyDefinition, found bool) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(PropertyKey(id))
	if bz == nil {
		return prop, false
	}
	keeper.cdc.MustUnmarshalBinaryBare(bz, &prop)
	return prop, true
}

// GetPropertyID returns the ID for a property URL if one exists, the graph.PropertyID 0 indicates no property with
// this URL is defined
func (keeper Keeper) GetPropertyID(ctx sdk.Context, propertyURL string) graph.PropertyID {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(PropertyURLKey(propertyURL))
	if bz == nil {
		return 0
	}
	var id graph.PropertyID
	keeper.cdc.MustUnmarshalBinaryBare(bz, &id)
	return id
}

// DefineProperty defines a property within the state store return the property's id and URL if it was defined
// successfully or else an error
func (keeper Keeper) DefineProperty(ctx sdk.Context, def PropertyDefinition) (id graph.PropertyID, uri *url.URL, err sdk.Error) {
	err = def.ValidateBasic()
	if err != nil {
		return 0, nil, err
	}
	uri = def.URI()
	uriKey := PropertyURLKey(uri.String())
	store := ctx.KVStore(keeper.storeKey)
	if store.Has(uriKey) {
		return id, uri, sdk.ErrUnknownRequest("property already defined")
	}
	id = keeper.newPropertyID(ctx)
	bz := keeper.cdc.MustMarshalBinaryBare(def)
	store.Set(PropertyKey(id), bz)
	store.Set(uriKey, keeper.cdc.MustMarshalBinaryBare(id))
	return id, uri, nil
}

func (keeper Keeper) newPropertyID(ctx sdk.Context) graph.PropertyID {
	lastID := keeper.GetLastPropertyID(ctx)
	id := lastID + 1
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryBare(id)
	store.Set(LastPropertyIDKey, bz)
	return id
}

func (keeper Keeper) GetLastPropertyID(ctx sdk.Context) graph.PropertyID {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(LastPropertyIDKey)
	var id graph.PropertyID
	if bz == nil {
		return 0
	}
	keeper.cdc.MustUnmarshalBinaryBare(bz, &id)
	return id
}
