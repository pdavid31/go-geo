package geo

import (
	"fmt"
	"strings"
)

type Multipoint []Point

/* CONSTRUCTOR */
func NewMultipoint(points ...*Point) Multipoint {
	mp := Multipoint{}

	for _, p := range points {
		mp = append(mp, *p)
	}

	return mp
}

/* GEOMETRY */
func (m Multipoint) GeometryType() string {
	return "Multipoint"
}

func (m Multipoint) SRID() int {
	return 4326
}

func (m Multipoint) Envelope() Geometry {
	panic("implement me")
}

func (m Multipoint) AsText() string {
	is3D := m.Is3D()
	lastIndex := len(m) - 2
	rep := strings.ToUpper(m.GeometryType()) + " ("
	for i, p := range m {
		x := fmt.Sprintf("%f", p.x)
		y := fmt.Sprintf("%f", p.y)

		rep += "(" + x + " " + y

		if is3D {
			z := fmt.Sprintf("%f", p.z)
			rep += " " + z
		}

		rep += ")"

		if i <= lastIndex {
			rep += " "
		}
	}

	rep += ")"
	return rep
}

func (m Multipoint) IsEmpty() bool {
	return len(m) == 0
}

func (m Multipoint) Is3D() bool {
	for _, p := range m {
		if !p.Is3D() {
			return false
		}
	}

	return true
}

func (m Multipoint) Equals(another Geometry) bool {
	panic("implement me")
}

func (m Multipoint) Disjoint(another Geometry) bool {
	panic("implement me")
}

func (m Multipoint) Intersects(another Geometry) bool {
	panic("implement me")
}

func (m Multipoint) Touches(another Geometry) bool {
	panic("implement me")
}

func (m Multipoint) Crosses(another Geometry) bool {
	panic("implement me")
}

func (m Multipoint) Within(another Geometry) bool {
	panic("implement me")
}

func (m Multipoint) Contains(another Geometry) bool {
	panic("implement me")
}

func (m Multipoint) Overlaps(another Geometry) bool {
	panic("implement me")
}

func (m Multipoint) Distance(another Geometry) float64 {
	panic("implement me")
}

func (m Multipoint) Buffer(distance float64) Geometry {
	panic("implement me")
}

func (m Multipoint) ConvexHull() Geometry {
	panic("implement me")
}

func (m Multipoint) Intersection(another Geometry) Geometry {
	panic("implement me")
}

func (m Multipoint) Union(another Geometry) Geometry {
	panic("implement me")
}

func (m Multipoint) Difference(another Geometry) Geometry {
	panic("implement me")
}

/* GEOMETRYCOLLECTION */
func (m Multipoint) NumGeometries() int {
	return len(m)
}

func (m Multipoint) GeometryN(n int) Geometry {
	return &m[n]
}
