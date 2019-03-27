package schema

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
)
import sdk "github.com/cosmos/cosmos-sdk/types"

// Keeper is the schema module keeper
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

var (
	nextPropertyIDKey = []byte("_/next-property-id")
)

// NewKeeper creates a new Keeper
func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{storeKey: storeKey, cdc: cdc}
}

// PropertyKey returns the store key for the given PropertyID that points to its PropertyDefinition
func PropertyKey(id PropertyID) []byte {
	return []byte(fmt.Sprintf("p/%d", id))
}

// PropertyURLKey returns the store key for the given property URL that points to its PropertyID
func PropertyURLKey(url string) []byte {
	return []byte(fmt.Sprintf("u/%s", url))
}

// GetProperty returns a PropertyDefinition given a PropertyID if one exists
func (keeper Keeper) GetProperty(ctx sdk.Context, id PropertyID) (prop PropertyDefinition, found bool) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(PropertyKey(id))
	if bz == nil {
		return prop, false
	}
	keeper.cdc.MustUnmarshalBinaryBare(bz, &prop)
	return prop, true
}

// GetPropertyID returns the ID for a property URL if one exists, the PropertyID 0 indicates no property with
// this URL is defined
func (keeper Keeper) GetPropertyID(ctx sdk.Context, propertyURL string) PropertyID {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(PropertyURLKey(propertyURL))
	if bz == nil {
		return 0
	}
	var id PropertyID
	keeper.cdc.MustUnmarshalBinaryBare(bz, &id)
	return id
}

// DefineProperty defines a property within the state store return the property's id and URL if it was defined
// successfully or else an error
func (keeper Keeper) DefineProperty(ctx sdk.Context, def PropertyDefinition) (id PropertyID, url string, err sdk.Error) {
	err = def.ValidateBasic()
	if err != nil {
		return 0, "", err
	}
	url = def.URL()
	store := ctx.KVStore(keeper.storeKey)
	if store.Has(PropertyURLKey(url)) {
		return id, url, sdk.ErrUnknownRequest("property already defined")
	}
	id = keeper.nextPropertyID(ctx)
	bz := keeper.cdc.MustMarshalBinaryBare(def)
	store.Set(PropertyKey(id), bz)
	store.Set(PropertyURLKey(url), keeper.cdc.MustMarshalBinaryBare(id))
	return id, url, nil
}

func (keeper Keeper) nextPropertyID(ctx sdk.Context) PropertyID {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(nextPropertyIDKey)
	var id PropertyID = 1
	if bz != nil {
		keeper.cdc.MustUnmarshalBinaryBare(bz, &id)
	}
	bz = keeper.cdc.MustMarshalBinaryBare(id + 1)
	store.Set(nextPropertyIDKey, bz)
	return id
}
