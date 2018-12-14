package esp

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
    // "github.com/twpayne/go-geom/encoding/ewkb"
)

type Keeper struct {
	espStoreKey sdk.StoreKey
	espVersionStoreKey sdk.StoreKey
	espResultStoreKey sdk.StoreKey

	cdc *codec.Codec
}

func NewKeeper(
	espStoreKey sdk.StoreKey,
	espVersionStoreKey sdk.StoreKey,
	espResultStoreKey sdk.StoreKey,
	cdc *codec.Codec) Keeper {
	return Keeper{
		espStoreKey:espStoreKey,
		espVersionStoreKey:espVersionStoreKey,
		espResultStoreKey:espResultStoreKey,
		cdc:          cdc,
	}
}

func (espk Keeper) CreateESP(ctx sdk.Context)  {

}
