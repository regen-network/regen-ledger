package geo

type GeometryType int

const (
	Polygon GeometryType = 1
	Point GeometryType = 2
)

type Geometry struct {
	EWKB []byte
	Type GeometryType
}
