package geo

type MultiLineString []LineString

/* CONSTRUCTOR */
func NewMultiLineString(linestrings ...*LineString) MultiLineString {
	mls := MultiLineString{}

	for _, ls := range linestrings {
		mls = append(mls, *ls)
	}

	return mls
}

/* GEOMETRY */
func (m MultiLineString) Lat() float64 {
	panic("implement me")
}

func (m MultiLineString) Lon() float64 {
	panic("implement me")
}

func (m MultiLineString) Z() float64 {
	panic("implement me")
}

func (m MultiLineString) GeometryType() string {
	panic("implement me")
}

func (m MultiLineString) SRID() int {
	panic("implement me")
}

func (m MultiLineString) Envelope() Geometry {
	panic("implement me")
}

func (m MultiLineString) AsText() string {
	panic("implement me")
}

func (m MultiLineString) IsEmpty() bool {
	panic("implement me")
}

func (m MultiLineString) Is3D() bool {
	panic("implement me")
}

func (m MultiLineString) Equals(another Geometry) bool {
	panic("implement me")
}

func (m MultiLineString) Disjoint(another Geometry) bool {
	panic("implement me")
}

func (m MultiLineString) Intersects(another Geometry) bool {
	panic("implement me")
}

func (m MultiLineString) Touches(another Geometry) bool {
	panic("implement me")
}

func (m MultiLineString) Crosses(another Geometry) bool {
	panic("implement me")
}

func (m MultiLineString) Within(another Geometry) bool {
	panic("implement me")
}

func (m MultiLineString) Contains(another Geometry) bool {
	panic("implement me")
}

func (m MultiLineString) Overlaps(another Geometry) bool {
	panic("implement me")
}

func (m MultiLineString) Distance(another Geometry) float64 {
	panic("implement me")
}

func (m MultiLineString) Buffer(distance float64) Geometry {
	panic("implement me")
}

func (m MultiLineString) ConvexHull() Geometry {
	panic("implement me")
}

func (m MultiLineString) Intersection(another Geometry) Geometry {
	panic("implement me")
}

func (m MultiLineString) Union(another Geometry) Geometry {
	panic("implement me")
}

func (m MultiLineString) Difference(another Geometry) Geometry {
	panic("implement me")
}

/* GEOMETRYCOLLECTION */
func (m MultiLineString) NumGeometries() int {
	panic("implement me")
}

func (m MultiLineString) GeometryN(n int) Geometry {
	panic("implement me")
}

/* CURVE */
func (m MultiLineString) Length() float64 {
	panic("implement me")
}

func (m MultiLineString) IsClosed() {
	panic("implement me")
}
