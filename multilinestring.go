package geo

import (
	"strings"
)

// MultiLineString type
type MultiLineString []LineString

/* CONSTRUCTOR */

// NewMultiLineString MultiLineString constructor
func NewMultiLineString(linestrings ...LineString) MultiLineString {
	var mls MultiLineString

	for _, ls := range linestrings {
		mls = append(mls, ls)
	}

	return mls
}

/* GEOMETRY */

// Lat function returns the latitude of the Geometry
func (m MultiLineString) Lat() float64 {
	panic("implement me")
}

// Lon function returns the longitude of the Geometry
func (m MultiLineString) Lon() float64 {
	panic("implement me")
}

// Alt function returns the altitude of the Geometry
func (m MultiLineString) Alt() float64 {
	panic("implement me")
}

// GeometryType returns GeometryType as a string
func (m MultiLineString) GeometryType() string {
	return "MultiLineString"
}

// SRID returns the SRID as an integer
func (m MultiLineString) SRID() int {
	return 4326
}

// Envelope returns the Envelope of the geometry
func (m MultiLineString) Envelope() Geometry {
	panic("implement me")
}

// AsText returns the string representation of the geometry
func (m MultiLineString) AsText() string {
	rep := strings.ToUpper(m.GeometryType()) + " "

	rep += m.ToString()

	return rep
}

// ToString returns the internal string representation of the geometry
func (m MultiLineString) ToString() string {
	rep := "("

	lastIndex := len(m) - 2
	for i, ls := range m {
		rep += ls.ToString()

		if i <= lastIndex {
			rep += ", "
		}
	}

	rep += ")"
	return rep
}

// IsEmpty returns a bool indicating if the geometry is empty
func (m MultiLineString) IsEmpty() bool {
	return len(m) == 0
}

// Is3D returns a bool indicating if the geometry is three dimensional
func (m MultiLineString) Is3D() bool {
	for _, ls := range m {
		if !ls.Is3D() {
			return false
		}
	}

	return true
}

// Equals returns true if the geometry is equal to a given other
func (m MultiLineString) Equals(another Geometry) bool {
	panic("implement me")
}

// Disjoint returns true if the geometry is disjoint to a given other
func (m MultiLineString) Disjoint(another Geometry) bool {
	panic("implement me")
}

// Intersects returns true if the geometry intersects another given geometry
func (m MultiLineString) Intersects(another Geometry) bool {
	panic("implement me")
}

// Touches returns true if the geometry touches another given geometry
func (m MultiLineString) Touches(another Geometry) bool {
	panic("implement me")
}

// Crosses returns true if the geometry crosses another given geometry
func (m MultiLineString) Crosses(another Geometry) bool {
	panic("implement me")
}

// Within returns true if the geometry is within another given geometry
func (m MultiLineString) Within(another Geometry) bool {
	panic("implement me")
}

// Contains returns true if the geometry contains another given geometry
func (m MultiLineString) Contains(another Geometry) bool {
	panic("implement me")
}

// Overlaps returns true if the geometry overlaps another given geometry
func (m MultiLineString) Overlaps(another Geometry) bool {
	panic("implement me")
}

// Distance calculates the distane to another given geometry
func (m MultiLineString) Distance(another Geometry) float64 {
	panic("implement me")
}

// Buffer returns the Buffer with the given distance around the geometry
func (m MultiLineString) Buffer(distance float64) Geometry {
	panic("implement me")
}

// ConvexHull returns the ConvexHull containing the geometry
func (m MultiLineString) ConvexHull() Geometry {
	panic("implement me")
}

// Intersection returns the Intersection of the geometry and another given geometry
func (m MultiLineString) Intersection(another Geometry) Geometry {
	panic("implement me")
}

// Union returns the Union of the geometry and another given geometry
func (m MultiLineString) Union(another Geometry) Geometry {
	panic("implement me")
}

// Difference returns the Difference of the geometry and another given geometry
func (m MultiLineString) Difference(another Geometry) Geometry {
	panic("implement me")
}

/* GEOMETRYCOLLECTION */

// NumGeometries returns the amount of geometries in the collection
func (m MultiLineString) NumGeometries() int {
	return len(m)
}

// GeometryN returns the Geometry a position n
func (m MultiLineString) GeometryN(n int) Geometry {
	return m[n]
}

/* MULTICURVE */

// Length returns the total length of the multicurve
func (m MultiLineString) Length() float64 {
	sum := 0.0

	for _, ls := range m {
		sum += ls.Length()
	}

	return sum
}

// IsClosed return whether each curve is closed or not
func (m MultiLineString) IsClosed() bool {
	for _, ls := range m {
		if !ls.IsClosed() {
			return false
		}
	}

	return true
}

/* INTERFACE GUARD */
var (
	_ Geometry           = (*MultiLineString)(nil)
	_ GeometryCollection = (*MultiLineString)(nil)
	_ MultiCurve         = (*MultiLineString)(nil)
)
