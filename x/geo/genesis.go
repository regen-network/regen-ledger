package geo

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/twpayne/go-geom/encoding/ewkb"
)

// GenesisState - geo genesis state
type GenesisState struct {
	Geometries []Geometry `json:"geometries"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState() GenesisState {
	return GenesisState{}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

// ValidateGenesis
func ValidateGenesis(data GenesisState) error {
	for _, geo := range data.Geometries {
		_, err := ValidateGeometry(geo)
		if err != nil {
			return err
		}
	}
	return nil
}

// InitGenesis
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	for _, geo := range data.Geometries {
		_, _ = keeper.StoreGeometry(ctx, geo)
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	store := ctx.KVStore(keeper.storeKey)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	var geometries []Geometry
	for ; iterator.Valid(); iterator.Next() {
		bz := iterator.Value()

		g, err := ewkb.Unmarshal(bz)
		if err != nil {
			panic(err)
		}

		ty, err := GetFeatureType(g)
		if err != nil {
			panic(err)
		}
		geometries = append(geometries, Geometry{EWKB: bz, Type: ty})
	}
	return GenesisState{Geometries: geometries}
}
