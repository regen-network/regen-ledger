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
	"github.com/twpayne/go-geom/encoding/geojson"
	"reflect"
	"testing"
)

// This struct defines a context that any test in this module
// can use to access a Keeper and an SDK context
type testCtx struct {
	keeper Keeper
	ctx    sdk.Context
}

// This is can be used by any test to configure a testCtx
func setupTestCtx() testCtx {
	db := dbm.NewMemDB()
	key := sdk.NewKVStoreKey("geo")

	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)
	_ = cms.LoadLatestVersion()

	cdc := codec.New()
	RegisterCodec(cdc)

	ctx := sdk.NewContext(cms, abci.Header{}, false, log.NewNopLogger())

	return testCtx{
		NewKeeper(key, cdc, nil),
		ctx,
	}
}

// This example test is going to generate some random geometries (geom.T),
// store them on the blockchain and test that we can retrieve the same geometries back.
//
// In order to do that we need to create a generator for Geometry that will generate
// random geometry values that correspond to what our geo module accepts.
//
// The main methods to use for creating new generators using gopter are
// are gopter.DeriveGen, gopter.Gen.Map and gopter.Gen.FlatMap. See FeatureTypeGen(),
// GenWGS84XYCoord(), and GenGeomT() for walkthroughs of how to create generators
// using these functions.
func TestKeeper_StoreGeometry(t *testing.T) {
	setup := setupTestCtx()
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

// This test is just here show how to generate sample
func TestGenerators(t *testing.T) {
	// We can use Sample to generate a sample value for a Gen to see if our generator is working correctly
	theFeatureType, wasGenerated := FeatureTypeGen().Sample()
	fmt.Println("A random FeatureType", theFeatureType, "was generated", wasGenerated)

	wgs84coord, wasGenerated := GenWGS84XYCoord().Sample()
	fmt.Println("A random WGS84 Coord", wgs84coord, "was generated", wasGenerated)

	coords, wasGenerated := GenCoords().Sample()
	fmt.Println("A random slice of WGS84 Coord's", coords, "was generated", wasGenerated)

	g, wasGenerated := GenGeomT().Sample()
	var geoJson []byte
	if wasGenerated {
		geoJson, _ = geojson.Marshal(g.(geom.T))
	}
	fmt.Println("A random geometry", string(geoJson), "was generated", wasGenerated)
}

// This generates an int using gen.IntRange and converts it to a FeatureType using gopter.Map
func FeatureTypeGen() gopter.Gen {
	// IntRange generates int's in a specified range (in this example we use the min and max values of
	// the FeatureType enumeration). IntRange is included in the gen package of gopter
	// along with a lot of other helpful generators that you should take a look at
	// (https://godoc.org/github.com/leanovate/gopter/gen).
	return gen.IntRange(int(Point), int(MultiPolygon)).
		// Map is a method on a Gen that takes a value produced by that generator
		// and lets the user provide a function that converts that value to another type.
		// It is used here to cast an int as a FeatureType
		Map(func(x int) FeatureType { return FeatureType(x) })
}

// This generates a geom.Coord with simple XY layout for the WGS84 projection
// It uses DeriveGen to take two float64's and turn them into a Coord
func GenWGS84XYCoord() gopter.Gen {
	return gopter.DeriveGen(
		// The first argument to DeriveGen is a function that takes the generated
		// parameters and maps them to the result type, in this case two float64's to
		// a Coord
		func(x float64, y float64) geom.Coord {
			return geom.Coord([]float64{x, y})
		},
		// The second argument to DeriveGen is a function that pulls the float64's
		// back out of a Coord. gopter may use this internally during its shrinking process
		// which helps narrow down failing cases to the smallest failing case
		// the test framework can find
		func(coord geom.Coord) (x float64, y float64) {
			return coord.X(), coord.Y()
		},
		// Here we provide the generators whose results get passed to the
		// function we passed in as the first argument to DeriveGen -
		// these are passed in as vararg's so you can provide as many as
		// you need but make sure they match the arguments to that function
		gen.Float64Range(-180.0, 180.0), // WGS84 Longitude range
		gen.Float64Range(-90.0, 90.0),   // WGS84 Latitude range
	)
}

// This function generates a slice of coordinates using FlatMap
func GenCoords() gopter.Gen {
	// We use gen.IntRange to generate the length of the slice that we are going to create randomly.
	// In this case we've decided to make a slice with 3 to 50 elements
	return gen.IntRange(3, 50).
		// FlatMap is a method on a Gen that let's you map a value generated from that
		// generator and depending on that value, create another generator
		FlatMap(
			// The first argument to FlatMap is a function that takes
			// a value generated by our original generator and returns a new
			// generator
			func(n interface{}) gopter.Gen {
				// in this case we're using the SliceOfN generator
				// to create a Gen that generates a slice of Coord's
				return gen.SliceOfN(
					// n is the length of the slice that we generated with gen.IntRange above
					n.(int),
					// this is the Gen that will be used to generate each element of the slice
					GenWGS84XYCoord(),
				)
			},
			// the second argument of FlatMap is the Type of value that gets generated by
			// our new generator. reflect.TypeOf can be used to get a Type from any value
			reflect.TypeOf([]geom.Coord{}),
		)
}

// This function generates a geom.T using FlatMap. It looks more involved
// but that is just because there are a lot of different cases
func GenGeomT() gopter.Gen {
	// In this case we are using FeatureTypeGen() to randomly choose what type
	// of feature we want to generate (i.e. Point, Polygon, etc.)
	return FeatureTypeGen().FlatMap(
		func(ft interface{}) gopter.Gen {
			switch ft.(FeatureType) {
			case Point:
				// Here DeriveGen is used to convert a Coord to a Point with the WGS84 SRID
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
				// Here DeriveGen is used to convert a []Coord to LineString
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
				// Here DeriveGen is used to convert a []Coord to Polygon
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
				// Here DeriveGen is used to convert a []Coord to MultiPoint
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
				// Here DeriveGen is used to convert a []Coord to MultiLineString
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
				// Here DeriveGen is used to convert a []Coord to MultiPolygon
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
				// Our generator shouldn't have produced a FeatureType we didn't expect
				// so if we get here we panic
				panic(fmt.Errorf("unexpected FeatureType"))
			}
		},
		// The interface type for geom.T (go requires some strange syntax to get this)
		reflect.TypeOf((*geom.T)(nil)))
}
