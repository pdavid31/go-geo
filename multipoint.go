package geo

import (
	"fmt"
	"strings"
)

type MultiPoint []Point

/* CONSTRUCTOR */
func NewMultiPoint(points ...Point) MultiPoint {
	var mp MultiPoint

	for _, p := range points {
		mp = append(mp, p)
	}

	return mp
}

/* GEOMETRY */
func (m MultiPoint) Lat() float64 {
	panic("implement me")
}

func (m MultiPoint) Lon() float64 {
	panic("implement me")
}

func (m MultiPoint) Z() float64 {
	panic("implement me")
}

func (m MultiPoint) GeometryType() string {
	return "MultiPoint"
}

func (m MultiPoint) SRID() int {
	return 4326
}

func (m MultiPoint) Envelope() Geometry {
	panic("implement me")
}

func (m MultiPoint) AsText() string {
	is3D := m.Is3D()
	lastIndex := len(m) - 2
	rep := strings.ToUpper(m.GeometryType()) + " ("
	for i, p := range m {
		x := fmt.Sprintf("%f", p.x)
		y := fmt.Sprintf("%f", p.y)

		rep += x + " " + y

		if is3D {
			z := fmt.Sprintf("%f", p.z)
			rep += " " + z
		}

		if i <= lastIndex {
			rep += ", "
		}
	}

	rep += ")"
	return rep
}

func (m MultiPoint) IsEmpty() bool {
	return len(m) == 0
}

func (m MultiPoint) Is3D() bool {
	for _, p := range m {
		if !p.Is3D() {
			return false
		}
	}

	return true
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

/* GEOMETRYCOLLECTION */
func (m MultiPoint) NumGeometries() int {
	return len(m)
}

func (m MultiPoint) GeometryN(n int) Geometry {
	return m[n]
}
