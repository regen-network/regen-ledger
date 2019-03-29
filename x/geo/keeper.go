package geo

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/index/postgresql"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/util"
	"golang.org/x/crypto/blake2b"
)

type Keeper struct {
	storeKey sdk.StoreKey

	cdc *codec.Codec

	pgIndexer postgresql.Indexer
}

const (
	WGS84_SRID = 4326
)

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec, pgIndexer postgresql.Indexer) Keeper {
	return Keeper{storeKey, cdc, pgIndexer}
}

func (keeper Keeper) GetGeometry(ctx sdk.Context, addr types.GeoAddress) []byte {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(addr)
	if bz == nil {
		return nil
	}
	return bz
}

func (keeper Keeper) StoreGeometry(ctx sdk.Context, geometry Geometry) (addr types.GeoAddress, err sdk.Error) {
	// TODO consume gas
	store := ctx.KVStore(keeper.storeKey)
	hash, e := blake2b.New256(nil)
	if e != nil {
		return nil, sdk.ErrUnknownRequest(e.Error())
	}
	ewkb := geometry.EWKB
	hash.Write(ewkb)
	hashBz := hash.Sum(nil)
	existing := store.Get(hashBz)
	if existing != nil {
		return nil, sdk.ErrUnknownRequest("already exists")
	}
	store.Set(hashBz, ewkb)

	addr = hashBz

	// Do Indexing
	if keeper.pgIndexer != nil {
		keeper.pgIndexer.Exec(
			"INSERT INTO geo (url, geog, geom) VALUES ($1, st_geogfromwkb($2), st_geomfromewkb($3))",
			addr.String(), ewkb, ewkb)
	}

	return addr, nil
}

func MustDecodeBech32GeoID(bech string) []byte {
	hrp, bz := util.MustDecodeBech32(bech)
	if hrp != types.Bech32GeoAddressPrefix {
		panic(fmt.Sprintf("Bech32 GeoID must start with %s", types.Bech32GeoAddressPrefix))
	}
	return bz
}
