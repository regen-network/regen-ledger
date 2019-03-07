package geo

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/regen-network/regen-ledger/index/postgresql"
	"gitlab.com/regen-network/regen-ledger/util"
	"golang.org/x/crypto/blake2b"
)

type Keeper struct {
	storeKey sdk.StoreKey

	cdc *codec.Codec

	pgIndexer postgresql.Indexer
}

const (
	Bech32Prefix = "xrn:geo/"
	WGS84_SRID   = 4326
)

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec, pgIndexer postgresql.Indexer) Keeper {
	return Keeper{storeKey, cdc, pgIndexer}
}

func (keeper Keeper) GetGeometry(ctx sdk.Context, addr GeoAddress) []byte {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(addr)
	if bz == nil {
		return nil
	}
	var geom []byte
	keeper.cdc.MustUnmarshalBinaryBare(bz, &geom)
	return geom
}

func GeoURL(addr GeoAddress) string {
	return util.MustEncodeBech32(Bech32Prefix, addr)
}

func (keeper Keeper) StoreGeometry(ctx sdk.Context, geometry Geometry) (addr GeoAddress, err sdk.Error) {
	// TODO consume gas
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryBare(geometry)
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
	store.Set(hashBz, bz)

	// Do Indexing
	if keeper.pgIndexer != nil {
		keeper.pgIndexer.Exec(
			"INSERT INTO geo (url, geog, geom) VALUES ($1, st_geogfromwkb($2), st_geomfromewkb($3))",
			GeoURL(hashBz), ewkb, ewkb)
	}

	return hashBz, nil
}

func MustDecodeBech32GeoID(bech string) []byte {
	hrp, bz := util.MustDecodeBech32(bech)
	if hrp != Bech32Prefix {
		panic(fmt.Sprintf("Bech32 GeoID must start with %s", Bech32Prefix))
	}
	return bz
}
