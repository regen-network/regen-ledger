package geo

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/regen-network/regen-ledger/utils"
	"golang.org/x/crypto/blake2b"
)

type Keeper struct {
	storeKey sdk.StoreKey

	cdc *codec.Codec
}

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{storeKey: storeKey, cdc: cdc}
}

func (keeper Keeper) GetGeometry(ctx sdk.Context, hash []byte) []byte {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(hash)
	if bz == nil {
		return nil
	}
	var geom []byte
	keeper.cdc.MustUnmarshalBinaryBare(bz, &geom)
	return geom
}

func (keeper Keeper) StoreGeometry(ctx sdk.Context, geometry Geometry) sdk.Result {
	// TODO consume gas
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryBare(geometry)
	hash, err := blake2b.New256(nil)
	if err != nil {
		panic(err)
	}
	hash.Write(geometry.EWKB)
	hashBz := hash.Sum(nil)
	existing := store.Get(hashBz)
	if existing != nil {
		return sdk.Result{
			Code: sdk.CodeUnknownRequest,
			Log:  "already exists",
		}
	}
	store.Set(hashBz, bz)
	tags := sdk.EmptyTags()
	tags = tags.AppendTag("geo.id", []byte(utils.MustEncodeBech32("xrngeo", hashBz)))
	return sdk.Result{Tags: tags}
}
