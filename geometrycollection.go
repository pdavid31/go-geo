package geo

type GeometryCollection interface {
	NumGeometries() int
	GeometryN(n int) Geometry
}
