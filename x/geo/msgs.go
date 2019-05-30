package geo

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/ewkb"
)

type MsgStoreGeometry struct {
	Data   Geometry       `json:"data"`
	Signer sdk.AccAddress `json:"signer"`
}

func (MsgStoreGeometry) Route() string { return "geo" }

func (MsgStoreGeometry) Type() string { return "geo.store" }

func GetFeatureType(g geom.T) (FeatureType, error) {
	switch g.(type) {
	case *geom.Point:
		return Point, nil
	case *geom.LineString:
		return LineString, nil
	case *geom.Polygon:
		return Polygon, nil
	case *geom.MultiPoint:
		return MultiPoint, nil
	case *geom.MultiLineString:
		return MultiLineString, nil
	case *geom.MultiPolygon:
		return MultiPolygon, nil
	}
	return -1, fmt.Errorf("unsupported geometry type %T", g)
}

func (msg MsgStoreGeometry) ValidateBasic() sdk.Error {
	_, err := ValidateGeometry(msg.Data)
	if err != nil {
		return sdk.ErrUnknownRequest(err.Error())
	}
	return nil
}

func ValidateGeometry(geom Geometry) (geom.T, error) {
	if len(geom.EWKB) <= 0 {
		return nil, fmt.Errorf("EWKB bytes cannot be empty")
	}

	g, err := ewkb.Unmarshal(geom.EWKB)

	if err != nil {
		return nil, fmt.Errorf("geometry is not in EWKB format: %+v", err)
	}

	if g.SRID() != WGS84_SRID {
		return nil, fmt.Errorf("geometry does not use WGS84 SRID, got %d", g.SRID())
	}

	featureType := geom.Type
	actual, err := GetFeatureType(g)
	if actual != featureType {
		return nil, fmt.Errorf("wrong FeatureType %d, expected %d", featureType, actual)
	}

	return g, nil
}

func (msg MsgStoreGeometry) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgStoreGeometry) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}
