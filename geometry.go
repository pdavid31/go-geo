package geo

type Geometry interface {
	// TODO: add complete simple feature model
	Lat() float64
	Lon() float64
	Z() float64

	/* MISC */
	// GeometryType returns subtype of geometry
	GeometryType() string
	// SRID returns Spatial Reference System ID of the geometric object
	SRID() int
	// Envelope returns the minimum bounding box for this geometric object
	Envelope() Geometry
	// AsText exports this geometric object to a specific Well-known Text Representation
	AsText() string
	// ToString returns WKT representation without prefix
	ToString() string
	// IsEmpty returns true if the geometric object is an empty Geometry
	IsEmpty() bool
	// Is3D returns true if the geometric object has a z coordinate
	Is3D() bool

	/* QUERY */
	// Equals returns true if this geometric object is "spatially equal" to another Geometry
	Equals(another Geometry) bool
	// Disjoint returns true if this geometric object is "spatially disjoint" from another Geometry
	Disjoint(another Geometry) bool
	// Intersects returns true if this geometric object is "spatially intersects" another Geometry
	Intersects(another Geometry) bool
	// Touches returns true if this geometric object is "spatially touches" another Geometry
	Touches(another Geometry) bool
	// Crosses returns true if this geometric object is "spatially crosses" another Geometry
	Crosses(another Geometry) bool
	// Within returns true if this geometric object is "spatially within" another Geometry
	Within(another Geometry) bool
	// Contains returns true if this geometric object is "spatially contains" another Geometry
	Contains(another Geometry) bool
	// Overlaps returns true if this geometric object is "spatially overlaps" another Geometry
	Overlaps(another Geometry) bool

	/* ANALYSIS */
	// Distance returns the shortest distance between any two Points in the two geometric objects as calculated in the spatial reference system of this geometric object
	// TODO: replace float64 distance with Distance struct
	Distance(another Geometry) float64
	// Buffer returns a geometric object that represents all Points whose distance from this geometric object is less than or equal to distance
	Buffer(distance float64) Geometry
	// ConvexHull returns a geometric object that represents the convex hull of this geometric object
	ConvexHull() Geometry
	// Intersection returns a geometric object that represents the Point set intersection of this geometric object with another Geometry
	Intersection(another Geometry) Geometry
	// Union returns a geometric object that represents the Point set union of this geometric object with another Geometry
	Union(another Geometry) Geometry
	// Difference returns a geometric object that represents the Point set difference of this geometric object with another Geometry
	Difference(another Geometry) Geometry
}
