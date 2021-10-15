package geo

import (
	"strings"
)

// MultiPolygon type
type MultiPolygon []Polygon

/* CONSTRUCTOR */

// NewMultiPolygon MultiPolygon constructor
func NewMultiPolygon(polygons ...Polygon) MultiPolygon {
	var mpoly MultiPolygon

	for _, poly := range polygons {
		mpoly = append(mpoly, poly)
	}

	return mpoly
}

/* GEOMETRY */

// Lat function returns the latitude of the Geometry
func (m MultiPolygon) Lat() float64 {
	panic("implement me")
}

// Lon function returns the longitude of the Geometry
func (m MultiPolygon) Lon() float64 {
	panic("implement me")
}

// Alt function returns the altitude of the Geometry
func (m MultiPolygon) Alt() float64 {
	panic("implement me")
}

// GeometryType returns GeometryType as a string
func (m MultiPolygon) GeometryType() string {
	return "MultiPolygon"
}

// SRID returns the SRID as an integer
func (m MultiPolygon) SRID() int {
	return 4326
}

// Envelope returns the Envelope of the geometry
func (m MultiPolygon) Envelope() Geometry {
	panic("implement me")
}

// AsText returns the string representation of the geometry
func (m MultiPolygon) AsText() string {
	rep := strings.ToUpper(m.GeometryType()) + " "

	rep += m.ToString()

	return rep
}

// ToString returns the internal string representation of the geometry
func (m MultiPolygon) ToString() string {
	rep := "("

	lastIndex := len(m) - 2
	for i, p := range m {
		rep += p.ToString()

		if i <= lastIndex {
			rep += ", "
		}
	}

	rep += ")"
	return rep
}

// IsEmpty returns a bool indicating if the geometry is empty
func (m MultiPolygon) IsEmpty() bool {
	return len(m) == 0
}

// Is3D returns a bool indicating if the geometry is three dimensional
func (m MultiPolygon) Is3D() bool {
	for _, p := range m {
		if !p.Is3D() {
			return false
		}
	}

	return true
}

// Equals returns true if the geometry is equal to a given other
func (m MultiPolygon) Equals(another Geometry) bool {
	panic("implement me")
}

// Disjoint returns true if the geometry is disjoint to a given other
func (m MultiPolygon) Disjoint(another Geometry) bool {
	panic("implement me")
}

// Intersects returns true if the geometry intersects another given geometry
func (m MultiPolygon) Intersects(another Geometry) bool {
	panic("implement me")
}

// Touches returns true if the geometry touches another given geometry
func (m MultiPolygon) Touches(another Geometry) bool {
	panic("implement me")
}

// Crosses returns true if the geometry crosses another given geometry
func (m MultiPolygon) Crosses(another Geometry) bool {
	panic("implement me")
}

// Within returns true if the geometry is within another given geometry
func (m MultiPolygon) Within(another Geometry) bool {
	panic("implement me")
}

// Contains returns true if the geometry contains another given geometry
func (m MultiPolygon) Contains(another Geometry) bool {
	panic("implement me")
}

// Overlaps returns true if the geometry overlaps another given geometry
func (m MultiPolygon) Overlaps(another Geometry) bool {
	panic("implement me")
}

// Distance calculates the distane to another given geometry
func (m MultiPolygon) Distance(another Geometry) float64 {
	panic("implement me")
}

// Buffer returns the Buffer with the given distance around the geometry
func (m MultiPolygon) Buffer(distance float64) Geometry {
	panic("implement me")
}

// ConvexHull returns the ConvexHull containing the geometry
func (m MultiPolygon) ConvexHull() Geometry {
	panic("implement me")
}

// Intersection returns the Intersection of the geometry and another given geometry
func (m MultiPolygon) Intersection(another Geometry) Geometry {
	panic("implement me")
}

// Union returns the Union of the geometry and another given geometry
func (m MultiPolygon) Union(another Geometry) Geometry {
	panic("implement me")
}

// Difference returns the Difference of the geometry and another given geometry
func (m MultiPolygon) Difference(another Geometry) Geometry {
	panic("implement me")
}

/* GEOMETRYCOLLECTION */

// NumGeometries returns the amount of geometries in the collection
func (m MultiPolygon) NumGeometries() int {
	return len(m)
}

// GeometryN returns the Geometry a position n
func (m MultiPolygon) GeometryN(n int) Geometry {
	return m[n]
}

/* MULTISURFACE */

// Area returns the total area of all surfaces
func (m MultiPolygon) Area() float64 {
	panic("implement me")
}

// Centroid returns the centroid of all surfaces
func (m MultiPolygon) Centroid() Point {
	panic("implement me")
}

// PointOnSurface returns a random point on a surface
func (m MultiPolygon) PointOnSurface() Point {
	panic("implement me")
}

/* INTERFACE GUARD */
var (
	_ Geometry           = (*MultiPolygon)(nil)
	_ GeometryCollection = (*MultiPolygon)(nil)
	_ MultiSurface       = (*MultiPolygon)(nil)
)
