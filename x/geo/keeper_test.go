package geo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/ewkb"
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

func GenWGS84XYCoord() gopter.Gen {
	return gopter.DeriveGen(
		func(x float64, y float64) geom.Coord {
			return geom.Coord([]float64{x, y})
		},
		func(coord geom.Coord) (x float64, y float64) {
			return coord.X(), coord.Y()
		},
		gen.Float64Range(-90.0, 90.0),
		gen.Float64Range(-180.0, 180.0),
	)
}

func GenCoords() gopter.Gen {
	return gen.IntRange(3, 50).FlatMap(
		func(n interface{}) gopter.Gen {
			return gen.SliceOfN(n.(int), GenWGS84XYCoord())
		},
		reflect.TypeOf([]geom.Coord{}),
	)
}

func GenGeomT() gopter.Gen {
	return FeatureTypeGen().FlatMap(
		func(ft interface{}) gopter.Gen {
			switch ft.(FeatureType) {
			case Point:
				return gopter.DeriveGen(
					func(coord geom.Coord) *geom.Point {
						pt := geom.NewPoint(geom.XY)
						pt.SetSRID(WGS84_SRID)
						return pt.MustSetCoords(coord)
					},
					func(pt *geom.Point) geom.Coord { return pt.Coords() },
					GenWGS84XYCoord(),
				)
			case LineString:
				return gopter.DeriveGen(
					func(coords []geom.Coord) *geom.LineString {
						ls := geom.NewLineString(geom.XY)
						ls.SetSRID(WGS84_SRID)
						return ls.MustSetCoords(coords)
					},
					func(poly *geom.LineString) []geom.Coord { return poly.Coords() },
					GenCoords(),
				)
			case Polygon:
				return gopter.DeriveGen(
					func(ring []geom.Coord) *geom.Polygon {
						poly := geom.NewPolygon(geom.XY)
						poly.SetSRID(WGS84_SRID)
						return poly.MustSetCoords([][]geom.Coord{ring})
					},
					func(poly *geom.Polygon) []geom.Coord { return poly.Coords()[0] },
					GenCoords(),
				)
			case MultiPoint:
				return gopter.DeriveGen(
					func(coords []geom.Coord) *geom.MultiPoint {
						pt := geom.NewMultiPoint(geom.XY)
						pt.SetSRID(WGS84_SRID)
						return pt.MustSetCoords(coords)
					},
					func(pt *geom.MultiPoint) []geom.Coord {
						return pt.Coords()
					},
					GenCoords(),
				)
			case MultiLineString:
				return gopter.DeriveGen(
					func(coords []geom.Coord) *geom.MultiLineString {
						ls := geom.NewMultiLineString(geom.XY)
						ls.SetSRID(WGS84_SRID)
						return ls.MustSetCoords([][]geom.Coord{coords})
					},
					func(poly *geom.MultiLineString) []geom.Coord { return poly.Coords()[0] },
					GenCoords(),
				)
			case MultiPolygon:
				return gopter.DeriveGen(
					func(ring []geom.Coord) *geom.MultiPolygon {
						poly := geom.NewMultiPolygon(geom.XY)
						poly.SetSRID(WGS84_SRID)
						return poly.MustSetCoords([][][]geom.Coord{{ring}})
					},
					func(poly *geom.MultiPolygon) []geom.Coord { return poly.Coords()[0][0] },
					GenCoords(),
				)
			default:
				panic(fmt.Errorf("unexpected FeatureType"))
			}
		},
		reflect.TypeOf((*geom.T)(nil)))
}

func TestKeeper_StoreGeometry(t *testing.T) {
	setup := setupTestInput()
	properties := gopter.NewProperties(nil)

	properties.Property("can store and retrieve geo shapes", prop.ForAll(
		func(g geom.T) (bool, error) {
			bz, err := ewkb.Marshal(g, binary.LittleEndian)
			if err != nil {
				return false, err
			}
			typ, err := GetFeatureType(g)
			if err != nil {
				return false, err
			}
			addr, err := setup.keeper.StoreGeometry(setup.ctx, Geometry{EWKB: bz, Type: typ})
			if err != nil {
				return false, fmt.Errorf(err.Error())
			}
			bzCopy := setup.keeper.GetGeometry(setup.ctx, addr)
			if !bytes.Equal(bz, bzCopy) {
				return false, fmt.Errorf("EWKB doesn't match")
			}
			gCopy, err := ewkb.Unmarshal(bzCopy)
			if err != nil {
				return false, err
			}
			if g.Layout() != gCopy.Layout() || g.SRID() != gCopy.SRID() { // TODO compare coords
				return false, fmt.Errorf("geometries don't match")
			}
			return true, nil
		},
		GenGeomT(),
	))

	properties.TestingRun(t)
}
