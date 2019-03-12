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
	if len(msg.Data.EWKB) <= 0 {
		return sdk.ErrUnknownRequest("GeometryEWKB cannot be empty")
	}

	g, err := ewkb.Unmarshal(msg.Data.EWKB)

	if err != nil {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Geometry is not in EWKB format: %+v", err))
	}

	if g.SRID() != WGS84_SRID {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Geometry does not use WGS84 SRID, got %d", g.SRID()))
	}

	featureType := msg.Data.Type
	actual, err := GetFeatureType(g)

	if err != nil {
		return sdk.ErrUnknownRequest(err.Error())
	}

	if actual != featureType {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Wrong FeatureType %d, expected %d", featureType, actual))
	}

	return nil
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
