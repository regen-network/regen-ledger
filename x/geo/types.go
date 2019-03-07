package geo

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgStoreGeo struct {
	Data   Geometry       `json:"data"`
	Signer sdk.AccAddress `json:"signer"`
}

type FeatureType int

const (
	Point FeatureType = iota
	LineString
	Polygon
	MultiPoint
	MultiLineString
	MultiPolygon
)

type Geometry struct {
	Type FeatureType `json:"type"`
	// EWKB representation of the geo feature. Must be in the WGS84 coordinate
	// system and represent a Point, LineString, Polygon, MultiPoint, MultiLineString or MultiPolygon
	EWKB []byte `json:"ewkb,omitempty"`
}

type GeoAddress []byte

const PostgresSchema = `
CREATE TABLE geo (
  url text NOT NULL PRIMARY KEY,
  -- Both the Postgis geography and geometry representations are stored
  geog geography NOT NULL,
  geom geometry NOT NULL
);

CREATE INDEX geo_geom_gist ON geo USING GIST ( geom );
`
