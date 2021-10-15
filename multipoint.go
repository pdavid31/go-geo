package geo

import (
	"strings"
)

// MultiPoint type
type MultiPoint []Point

/* CONSTRUCTOR */

// NewMultiPoint MultiPoint constructor
func NewMultiPoint(points ...Point) MultiPoint {
	var mp MultiPoint

	for _, p := range points {
		mp = append(mp, p)
	}

	return mp
}

/* GEOMETRY */

// Lat function returns the latitude of the Geometry
func (m MultiPoint) Lat() float64 {
	panic("implement me")
}

// Lon function returns the longitude of the Geometry
func (m MultiPoint) Lon() float64 {
	panic("implement me")
}

// Alt function returns the altitude of the Geometry
func (m MultiPoint) Alt() float64 {
	panic("implement me")
}

// GeometryType returns GeometryType as a string
func (m MultiPoint) GeometryType() string {
	return "MultiPoint"
}

// SRID returns the SRID as an integer
func (m MultiPoint) SRID() int {
	return 4326
}

// Envelope returns the Envelope of the geometry
func (m MultiPoint) Envelope() Geometry {
	panic("implement me")
}

// AsText returns the string representation of the geometry
func (m MultiPoint) AsText() string {
	rep := strings.ToUpper(m.GeometryType()) + " "

	rep += m.ToString()

	return rep
}

// ToString returns the internal string representation of the geometry
func (m MultiPoint) ToString() string {
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
func (m MultiPoint) IsEmpty() bool {
	return len(m) == 0
}

// Is3D returns a bool indicating if the geometry is three dimensional
func (m MultiPoint) Is3D() bool {
	for _, p := range m {
		if !p.Is3D() {
			return false
		}
	}

	return true
}

// Equals returns true if the geometry is equal to a given other
func (m MultiPoint) Equals(another Geometry) bool {
	panic("implement me")
}

// Disjoint returns true if the geometry is disjoint to a given other
func (m MultiPoint) Disjoint(another Geometry) bool {
	panic("implement me")
}

// Intersects returns true if the geometry intersects another given geometry
func (m MultiPoint) Intersects(another Geometry) bool {
	panic("implement me")
}

// Touches returns true if the geometry touches another given geometry
func (m MultiPoint) Touches(another Geometry) bool {
	panic("implement me")
}

// Crosses returns true if the geometry crosses another given geometry
func (m MultiPoint) Crosses(another Geometry) bool {
	panic("implement me")
}

// Within returns true if the geometry is within another given geometry
func (m MultiPoint) Within(another Geometry) bool {
	panic("implement me")
}

// Contains returns true if the geometry contains another given geometry
func (m MultiPoint) Contains(another Geometry) bool {
	panic("implement me")
}

// Overlaps returns true if the geometry overlaps another given geometry
func (m MultiPoint) Overlaps(another Geometry) bool {
	panic("implement me")
}

// Distance calculates the distane to another given geometry
func (m MultiPoint) Distance(another Geometry) float64 {
	panic("implement me")
}

// Buffer returns the Buffer with the given distance around the geometry
func (m MultiPoint) Buffer(distance float64) Geometry {
	panic("implement me")
}

// ConvexHull returns the ConvexHull containing the geometry
func (m MultiPoint) ConvexHull() Geometry {
	panic("implement me")
}

// Intersection returns the Intersection of the geometry and another given geometry
func (m MultiPoint) Intersection(another Geometry) Geometry {
	panic("implement me")
}

// Union returns the Union of the geometry and another given geometry
func (m MultiPoint) Union(another Geometry) Geometry {
	panic("implement me")
}

// Difference returns the Difference of the geometry and another given geometry
func (m MultiPoint) Difference(another Geometry) Geometry {
	panic("implement me")
}

/* GEOMETRYCOLLECTION */

// NumGeometries returns the amount of geometries in the collection
func (m MultiPoint) NumGeometries() int {
	return len(m)
}

// GeometryN returns the Geometry a position n
func (m MultiPoint) GeometryN(n int) Geometry {
	return m[n]
}

/* INTERFACE GUARD */
var (
	_ Geometry           = (*MultiPoint)(nil)
	_ GeometryCollection = (*MultiPoint)(nil)
)
