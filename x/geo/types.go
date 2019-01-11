package geo

type GeometryType int

//const (
//	Point   GeometryType = 0
//	Polygon GeometryType = 1
//)

type Geometry struct {
	//Type GeometryType `json:"type"`
	EWKB []byte `json:"ewkb,omitempty"`
}
