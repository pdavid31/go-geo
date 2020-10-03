package geo

// GeometryCollection interface describes functions on the GeometryCollection object type
type GeometryCollection interface {
	NumGeometries() int
	GeometryN(n int) Geometry
}
