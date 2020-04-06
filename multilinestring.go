package geo

import (
	"strings"
)

type MultiLineString []LineString

/* CONSTRUCTOR */
func NewMultiLineString(linestrings ...LineString) MultiLineString {
	var mls MultiLineString

	for _, ls := range linestrings {
		mls = append(mls, ls)
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
	return "MultiLineString"
}

func (m MultiLineString) SRID() int {
	return 4326
}

func (m MultiLineString) Envelope() Geometry {
	panic("implement me")
}

func (m MultiLineString) AsText() string {
	rep := strings.ToUpper(m.GeometryType()) + " "

	rep += m.ToString()

	return rep
}

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

func (m MultiLineString) IsEmpty() bool {
	return len(m) == 0
}

func (m MultiLineString) Is3D() bool {
	for _, ls := range m {
		if !ls.Is3D() {
			return false
		}
	}

	return true
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
	return len(m)
}

func (m MultiLineString) GeometryN(n int) Geometry {
	return m[n]
}

/* MULTICURVE */
func (m MultiLineString) Length() float64 {
	sum := 0.0

	for _, ls := range m {
		sum += ls.Length()
	}

	return sum
}

func (m MultiLineString) IsClosed() bool {
	for _, ls := range m {
		if !ls.IsClosed() {
			return false
		}
	}

	return true
}
