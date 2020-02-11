package geo

type Geometry interface {
	// misc
	Dimension() int
	GeometryType() string
	SRID() int
	AsText() string
	IsEmpty() bool
	Boundary() Geometry

	// query
	Equals(another Geometry) bool
	Disjoint(another Geometry) bool
	Intersects(another Geometry) bool
	Touches(another Geometry) bool
	Within(another Geometry) bool
	Contains(another Geometry) bool
	Overlaps(another Geometry) bool

	// analysis
	// TODO: replace float64 distance with Distance struct
	Distance(another Geometry) float64
	Buffer(distance float64) Geometry
	ConvexHull() Geometry
	Intersection(another Geometry) Geometry
	Union(another Geometry) Geometry
	Difference(another Geometry) Geometry
}
