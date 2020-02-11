package geo

type MultiPoint []Point

func (m MultiPoint) GeometryType() string {
	return "Multipoint"
}

func (m MultiPoint) SRID() int {
	return 4326
}

func (m MultiPoint) Envelope() Geometry {
	panic("implement me")
}

func (m MultiPoint) AsText() string {
	panic("implement me")
}

func (m MultiPoint) IsEmpty() bool {
	return len(m) == 0
}

func (m MultiPoint) Is3D() bool {
	panic("implement me")
}

func (m MultiPoint) Equals(another Geometry) bool {
	panic("implement me")
}

func (m MultiPoint) Disjoint(another Geometry) bool {
	panic("implement me")
}

func (m MultiPoint) Intersects(another Geometry) bool {
	panic("implement me")
}

func (m MultiPoint) Touches(another Geometry) bool {
	panic("implement me")
}

func (m MultiPoint) Crosses(another Geometry) bool {
	panic("implement me")
}

func (m MultiPoint) Within(another Geometry) bool {
	panic("implement me")
}

func (m MultiPoint) Contains(another Geometry) bool {
	panic("implement me")
}

func (m MultiPoint) Overlaps(another Geometry) bool {
	panic("implement me")
}

func (m MultiPoint) Distance(another Geometry) float64 {
	panic("implement me")
}

func (m MultiPoint) Buffer(distance float64) Geometry {
	panic("implement me")
}

func (m MultiPoint) ConvexHull() Geometry {
	panic("implement me")
}

func (m MultiPoint) Intersection(another Geometry) Geometry {
	panic("implement me")
}

func (m MultiPoint) Union(another Geometry) Geometry {
	panic("implement me")
}

func (m MultiPoint) Difference(another Geometry) Geometry {
	panic("implement me")
}
