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

// This file gives a walk-through tutorial of how to create property-based unit tests
// using gopter. It also shows the generally method for doing unit tests of Cosmos SDK
// module keepers
//
// Property-based testing is a methodology of using randomly generated values in testing.
// It allows us to easily test a lot more edge cases than we could create manually.
// Here's a good intro article: https://hypothesis.works/articles/what-is-property-based-testing/

// This struct defines a context that any test in this module
// can use to access a Keeper and an SDK context. When testing a
// keeper we would usually define something like this at minimum
// and we will often want to add other values here - like keepers
// for other modules we're using. See the many examples in the Cosmos SDK
// itself for more info
type testCtx struct {
	keeper Keeper
	ctx    sdk.Context
}

// This function actually sets up our testCtx and can be used by
// any test in this module to get an instance of the keeper
func setupTestCtx() testCtx {
	// Create a memory DB and a CommitMultiStore. In production
	// these things are done in sdk.BaseApp using a disk-backed DB
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)

	// Create a key for our store. In more complex cases we'll need keys for other module stores
	key := sdk.NewKVStoreKey("geo")

	// Create an amino codec
	cdc := codec.New()
	// and register it with this module. In more complex cases, we'll have to register other module codecs
	RegisterCodec(cdc)

	// Create a Keeper for this module. In more complex cases, we'll need to create keepers for other modules and
	// link them together
	keeper := NewKeeper(key, cdc, nil)

	// Mount the key on our store. For more complex keeper's, we'll have to mount the keys for other module stores.
	// The equivalent of app.MountStores used in production
	cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)

	// Load the latest version of the CommitMultiStore and
	// create an sdk.Context for our CommitMultiStore. Done inside sdk.BaseApp in production
	_ = cms.LoadLatestVersion()
	ctx := sdk.NewContext(cms, abci.Header{}, false, log.NewNopLogger())

	// Return a testCtx
	return testCtx{keeper, ctx}
}

// This example test is going to generate some random geometries (geom.T) using the property-based testing framework
// provided by gopter, store them on the blockchain and test that we can retrieve the same geometries back.
//
func TestKeeper_StoreGeometry(t *testing.T) {
	// Setup the test context
	setup := setupTestCtx()

	// Create new properties for us to define gopter property tests.
	// We can pass in some parameters here, for example to
	// change the number of tests to run for each property (the default is 100).
	params := gopter.DefaultTestParameters()
	params.MaxDiscardRatio = 10 // TODO: fix this, we need to set it because our generator produces so many nil values
	properties := gopter.NewProperties(params)

	// This adds a gopter.Prop to test
	properties.Property("can store and retrieve geo shapes",
		// prop.ForAll creates a gopter.Prop that tests a "for all"
		// type assertion. In this case "for all geometries that we can create, we can store and retrieve them"
		prop.ForAll(
			// The first argument to ForAll is a function that takes a generated value and
			// tests whether our property holds for that generated value. In this case we
			// generate a geometry (our test is "for all geometries") and test whether the property
			// applies. We return true if the property holds or false and an error value if it fails
			func(g geom.T) (bool, error) {
				return CanSaveAndRetrieveAGeometry(setup, g)
			},

			// The second argument of ForAll is the generator that generates the values
			// we use the above function to test a property for
			GenGeomT(),
			// GenGeomT() returns a custom generator (an instance of gopter.Gen) that we have created in this file
			//
			// Gopter's gen package provides basic generators that you should familiarize yourself with first:
			// https://godoc.org/github.com/leanovate/gopter/gen.
			//
			// The rest of this file gives a walkthrough of how to create customer generators using the built-in
			// generators and the Map, DeriveGen and FlatMap methods. This approach should enable you to create custom
			// generators for most use cases. Occasionally, you may need other functionality like the SuchThat()
			// method. Refer to the gopter documentation for more details.
		))

	// Run the property tests
	properties.TestingRun(t)
}

func CanSaveAndRetrieveAGeometry(setup testCtx, g geom.T) (bool, error) {
	// Convert the geometry to EWKB format
	bz, err := ewkb.Marshal(g, binary.LittleEndian)
	if err != nil {
		return false, err
	}

	// Get the FeatureType of the geom
	typ, err := GetFeatureType(g)
	if err != nil {
		return false, err
	}

	// Try to store the geometry
	addr, err := setup.keeper.StoreGeometry(setup.ctx, Geometry{EWKB: bz, Type: typ})
	if err != nil {
		return false, fmt.Errorf(err.Error())
	}

	// Try to retrieve the geometry
	bzCopy := setup.keeper.GetGeometry(setup.ctx, addr)
	if !bytes.Equal(bz, bzCopy) {
		return false, fmt.Errorf("EWKB doesn't match")
	}

	// Convert it back from EWKB format
	gCopy, err := ewkb.Unmarshal(bzCopy)
	if err != nil {
		return false, err
	}

	// Check that things match
	err = GeomEqual(g, gCopy)
	if err != nil {
		return false, err
	}

	// If we get here the property held successfully
	return true, nil
}

func GenGeomsTest(geomGen gopter.Gen) {
	geoJsonGen := geomGen.Map(func(g geom.T) string {
		geoJsonBz, _ := geojson.Marshal(g.(geom.T))
		geoJsonStr := string(geoJsonBz)
		return geoJsonStr
	})
	samples, wasGenerated := gen.SliceOfN(10, geoJsonGen).Sample()
	fmt.Println("Random geometries", samples, "were generated", wasGenerated)
}

// Before we get into creating custom generators, this test shows us how to play around with generators using their Sample method
func TestGenerators(t *testing.T) {
	// Sample() generates a sample value of a Gen. We can use it to see if generator is working as we expect it.
	// It returns two values - first the generated value (if one was generated) and second, a bool
	// indicating whether a value was or wasn't generated. Sometimes due to the state of a generator,
	// gopter can't generate a value and you'll either need to do some tweaking of your generator or call Sample()
	// a few more times
	theFeatureType, wasGenerated := FeatureTypeGen().Sample()
	fmt.Println("A random FeatureType", theFeatureType, "was generated", wasGenerated)

	wgs84coord, wasGenerated := GenWGS84XYCoord().Sample()
	fmt.Println("A random WGS84 Coord", wgs84coord, "was generated", wasGenerated)

	coords, wasGenerated := GenCoords().Sample()
	fmt.Println("A random slice of WGS84 Coord's", coords, "was generated", wasGenerated)

	GenGeomsTest(GenPoint())
	GenGeomsTest(GenLineString())
	GenGeomsTest(GenPolygon())
	GenGeomsTest(GenMultiPoint())
	GenGeomsTest(GenMultiLineString())
	GenGeomsTest(GenPolygon())
	GenGeomsTest(GenGeomT())
}

// This is our first custom generator.
// It generates a geom.Coord with a latitude/longitude (XY) coordinate in the WGS84 projection.
// It uses DeriveGen to take two float64's and turn them into a Coord.
// DeriveGen is the recommended method to create new generators when possible.
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

// This function generates a slice of coordinates using FlatMap. FlatMap is only to be used instead of DeriveGen
// when we know we need to create a new generator from a generated value. If you are just mapping values from some
// existing values into a new type, you should probably use DeriveGen. Here we are randomly generated a number,
// and then creating a new generator based on that number
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

// This generates an int using gen.IntRange and converts it to a FeatureType using gopter.DeriveGen
func FeatureTypeGen() gopter.Gen {
	// along with a lot of other helpful generators that you should take a look at
	//
	return gopter.DeriveGen(
		// cast int to FeatureType
		func(x int) FeatureType {
			return FeatureType(x)
		},
		// cast FeatureType back to int
		func(ft FeatureType) int {
			return int(ft)
		},
		// gen.IntRange generates int's in the range of min and max values of the FeatureType enumeration
		gen.IntRange(int(Point), int(MultiPolygon)),
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
				return GenPoint()
			case LineString:
				// Here DeriveGen is used to convert a []Coord to LineString
				return GenLineString()
			case Polygon:
				// Here DeriveGen is used to convert a []Coord to Polygon
				return GenPolygon()
			case MultiPoint:
				// Here DeriveGen is used to convert a []Coord to MultiPoint
				return GenMultiPoint()
			case MultiLineString:
				// Here DeriveGen is used to convert a []Coord to MultiLineString
				return GenMultiLineString()
			case MultiPolygon:
				// Here DeriveGen is used to convert a []Coord to MultiPolygon
				return GenMultiPolygon()
			default:
				// Our generator shouldn't have produced a FeatureType we didn't expect
				// so if we get here we panic
				panic(fmt.Errorf("unexpected FeatureType"))
			}
		},
		// The interface type for geom.T (go requires some strange syntax to get this)
		reflect.TypeOf((*geom.T)(nil)))
}

func GenPoint() gopter.Gen {
	return gopter.DeriveGen(
		func(coord geom.Coord) *geom.Point {
			pt := geom.NewPoint(geom.XY)
			pt.SetSRID(WGS84_SRID)
			return pt.MustSetCoords(coord)
		},
		func(pt *geom.Point) geom.Coord { return pt.Coords() },
		GenWGS84XYCoord(),
	)
}

func GenLineString() gopter.Gen {
	// Map() is a simpler way of creating a new generator from an existing generator
	// than DeriveGen(). It just takes one parameter which is exactly the same as the
	// first parameter to DeriveGen(). It appears that sometimes the two-way mapping
	// that DeriveGen() causes a lot of data to be thrown away (possibly if the underlying type does some
	// internal modifications to the generated data). So in many cases, it may be
	// better to just use the simpler Map() method if generators with DeriveGen() cause
	// a lot of data to be thrown away. Use the Sample() method to test this as demonstrated
	// in TestGenerators()
	return GenCoords().Map(
		func(coords []geom.Coord) *geom.LineString {
			ls := geom.NewLineString(geom.XY)
			ls.SetSRID(WGS84_SRID)
			ls = ls.MustSetCoords(coords)
			return ls
		},
	)
}

func GenPolygon() gopter.Gen {
	return GenCoords().Map(
		func(ring []geom.Coord) *geom.Polygon {
			poly := geom.NewPolygon(geom.XY)
			poly.SetSRID(WGS84_SRID)
			return poly.MustSetCoords([][]geom.Coord{ring})
		},
	)
}

func GenMultiPoint() gopter.Gen {
	return GenCoords().Map(
		func(coords []geom.Coord) *geom.MultiPoint {
			pt := geom.NewMultiPoint(geom.XY)
			pt.SetSRID(WGS84_SRID)
			return pt.MustSetCoords(coords)
		},
	)
}

func GenMultiLineString() gopter.Gen {
	return GenCoords().Map(
		func(coords []geom.Coord) *geom.MultiLineString {
			ls := geom.NewMultiLineString(geom.XY)
			ls.SetSRID(WGS84_SRID)
			return ls.MustSetCoords([][]geom.Coord{coords})
		},
	)
}

func GenMultiPolygon() gopter.Gen {
	return GenCoords().Map(
		func(ring []geom.Coord) *geom.MultiPolygon {
			poly := geom.NewMultiPolygon(geom.XY)
			poly.SetSRID(WGS84_SRID)
			return poly.MustSetCoords([][][]geom.Coord{{ring}})
		},
	)
}

func GeomEqual(a geom.T, b geom.T) error {
	if a.Layout() != b.Layout() {
		return fmt.Errorf("layouts not equal")
	}
	if a.SRID() != b.SRID() {
		return fmt.Errorf("SRID's not equal")
	}
	if !floatSlicesEqual(a.FlatCoords(), b.FlatCoords()) {
		return fmt.Errorf("flat coords not equal: %+v, %+v", a.FlatCoords(), b.FlatCoords())
	}
	if !intSlicesEqual(a.Ends(), b.Ends()) {
		return fmt.Errorf("ends not equal: %+v, %+v", a.Ends(), b.Ends())
	}
	return nil
}

func floatSlicesEqual(a []float64, b []float64) bool {

	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func intSlicesEqual(a []int, b []int) bool {

	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
