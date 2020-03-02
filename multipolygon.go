package geo

type MultiPolygon []Polygon

/* CONSTRUCTOR */
func NewMultiPolygon(polygons ...Polygon) MultiPolygon {
	var mpoly MultiPolygon

	for _, poly := range polygons {
		mpoly = append(mpoly, poly)
	}

	return mpoly
}

/* GEOMETRY */
func (m MultiPolygon) Lat() float64 {
	panic("implement me")
}

func (m MultiPolygon) Lon() float64 {
	panic("implement me")
}

func (m MultiPolygon) Z() float64 {
	panic("implement me")
}

func (m MultiPolygon) GeometryType() string {
	return "MultiPolygon"
}

func (m MultiPolygon) SRID() int {
	return 4326
}

func (m MultiPolygon) Envelope() Geometry {
	panic("implement me")
}

func (m MultiPolygon) AsText() string {
	panic("implement me")
}

func (m MultiPolygon) IsEmpty() bool {
	return len(m) == 0
}

func (m MultiPolygon) Is3D() bool {
	for _, p := range m {
		if !p.Is3D() {
			return false
		}
	}

	return true
}

func (m MultiPolygon) Equals(another Geometry) bool {
	panic("implement me")
}

func (m MultiPolygon) Disjoint(another Geometry) bool {
	panic("implement me")
}

func (m MultiPolygon) Intersects(another Geometry) bool {
	panic("implement me")
}

func (m MultiPolygon) Touches(another Geometry) bool {
	panic("implement me")
}

func (m MultiPolygon) Crosses(another Geometry) bool {
	panic("implement me")
}

func (m MultiPolygon) Within(another Geometry) bool {
	panic("implement me")
}

func (m MultiPolygon) Contains(another Geometry) bool {
	panic("implement me")
}

func (m MultiPolygon) Overlaps(another Geometry) bool {
	panic("implement me")
}

func (m MultiPolygon) Distance(another Geometry) float64 {
	panic("implement me")
}

func (m MultiPolygon) Buffer(distance float64) Geometry {
	panic("implement me")
}

func (m MultiPolygon) ConvexHull() Geometry {
	panic("implement me")
}

func (m MultiPolygon) Intersection(another Geometry) Geometry {
	panic("implement me")
}

func (m MultiPolygon) Union(another Geometry) Geometry {
	panic("implement me")
}

func (m MultiPolygon) Difference(another Geometry) Geometry {
	panic("implement me")
}

/* GEOMETRYCOLLECTION */
func (m MultiPolygon) NumGeometries() int {
	return len(m)
}

func (m MultiPolygon) GeometryN(n int) Geometry {
	return m[n]
}

/* MULTISURFACE */
func (m MultiPolygon) Area() float64 {
	panic("implement me")
}

func (m MultiPolygon) Centroid() Point {
	panic("implement me")
}

func (m MultiPolygon) PointOnSurface() Point {
	panic("implement me")
}
