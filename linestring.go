package geo

import (
	"fmt"
	"reflect"
	"strings"
)

type Linestring []Point

/* CONSTRUCTOR */
func NewLinestring(points ...*Point) Linestring {
	ls := Linestring{}

	for _, p := range points {
		ls = append(ls, *p)
	}

	return ls
}

/* GEOMETRY */
func (l Linestring) GeometryType() string {
	return "Linestring"
}

func (l Linestring) SRID() int {
	return 4326
}

func (l Linestring) Envelope() Geometry {
	panic("implement me")
}

func (l Linestring) AsText() string {
	is3D := l.Is3D()
	lastIndex := len(l) - 2
	rep := strings.ToUpper(l.GeometryType()) + " ("
	for i, p := range l {
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

func (l Linestring) IsEmpty() bool {
	return len(l) == 0
}

func (l Linestring) Is3D() bool {
	for _, p := range l {
		if !p.Is3D() {
			return false
		}
	}

	return true
}

func (l Linestring) Equals(another Geometry) bool {
	panic("implement me")
}

func (l Linestring) Disjoint(another Geometry) bool {
	panic("implement me")
}

func (l Linestring) Intersects(another Geometry) bool {
	panic("implement me")
}

func (l Linestring) Touches(another Geometry) bool {
	panic("implement me")
}

func (l Linestring) Crosses(another Geometry) bool {
	panic("implement me")
}

func (l Linestring) Within(another Geometry) bool {
	panic("implement me")
}

func (l Linestring) Contains(another Geometry) bool {
	panic("implement me")
}

func (l Linestring) Overlaps(another Geometry) bool {
	panic("implement me")
}

func (l Linestring) Distance(another Geometry) float64 {
	panic("implement me")
}

func (l Linestring) Buffer(distance float64) Geometry {
	panic("implement me")
}

func (l Linestring) ConvexHull() Geometry {
	panic("implement me")
}

func (l Linestring) Intersection(another Geometry) Geometry {
	panic("implement me")
}

func (l Linestring) Union(another Geometry) Geometry {
	panic("implement me")
}

func (l Linestring) Difference(another Geometry) Geometry {
	panic("implement me")
}

/* CURVE */
func (l Linestring) Length() float64 {
	length := 0.
	lastIndex := len(l) - 1

	for i := range l {
		if i >= lastIndex {
			break
		}

		current := l[i]
		successor := l[i+1]
		length += current.Distance(&successor)
	}

	return length
}

func (l Linestring) StartPoint() Point {
	return l[0]
}

func (l Linestring) EndPoint() Point {
	return l[len(l)-1]
}

func (l Linestring) IsClosed() bool {
	start := l.StartPoint()
	end := l.EndPoint()

	return reflect.DeepEqual(start, end)
}
