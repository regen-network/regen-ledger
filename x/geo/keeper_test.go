package geo

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/twpayne/go-geom"
	"reflect"
	"testing"
)

type testInput struct {
	keeper Keeper
	ctx    sdk.Context
}

func setupTestInput() testInput {
	db := dbm.NewMemDB()
	key := sdk.NewKVStoreKey("geo")

	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)
	_ = cms.LoadLatestVersion()

	cdc := codec.New()
	RegisterCodec(cdc)

	ctx := sdk.NewContext(cms, abci.Header{}, false, log.NewNopLogger())

	return testInput{
		NewKeeper(key, cdc, nil),
		ctx,
	}
}

func FeatureTypeGen() gopter.Gen {
	return gen.IntRange(int(Point), int(MultiPolygon)).Map(func(x int) FeatureType {
		return FeatureType(x)
	})
}

//func GenWGS84XYCoords() gopter.Gen {
//	return gopter.DeriveGen()
//}

func GeomTGen() gopter.Gen {
	return FeatureTypeGen().FlatMap(
		func(ft interface{}) gopter.Gen {
			switch ft.(FeatureType) {
			case Point:
				return gopter.DeriveGen(
					func(x float64, y float64) *geom.Point {
						pt := geom.NewPointFlat(geom.XY, []float64{x, y})
						pt.SetSRID(4326) //WGS84
						return pt
					},
					func(pt *geom.Point) (x float64, y float64) {
						return pt.X(), pt.Y()
					},
					gen.Float64Range(-90.0, 90.0),
					gen.Float64Range(-180.0, 180.0),
				)
			case Polygon:
				return gopter.DeriveGen(
					func(x float64, y float64) *geom.Polygon {
						poly := geom.NewPolygon(geom.XY)
						return poly
					},
					func(pt *geom.Point) (x float64, y float64) {
						return pt.X(), pt.Y()
					},
					gen.Float64Range(-90.0, 90.0),
					gen.Float64Range(-180.0, 180.0),
				)
			default:
				panic(fmt.Errorf("unexpected FeatureType"))
			}
		},
		reflect.TypeOf((*geom.T)(nil)))
}

//func EWKBGen() gopter.Gen {
//	return GeomTGen().Map(func(g) {})
//}

func TestKeeper_StoreGeometry(t *testing.T) {
	//input := setupTestInput()

}
