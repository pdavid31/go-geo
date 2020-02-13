package geo

import (
	"fmt"
	"reflect"
	"strings"
)

type LineString []Point

/* CONSTRUCTOR */
func NewLineString(points ...*Point) LineString {
	ls := LineString{}

	for _, p := range points {
		ls = append(ls, *p)
	}

	return ls
}

/* GEOMETRY */
func (l LineString) GeometryType() string {
	return "LineString"
}

func (l LineString) SRID() int {
	return 4326
}

func (l LineString) Envelope() Geometry {
	panic("implement me")
}

func (l LineString) AsText() string {
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

func (l LineString) IsEmpty() bool {
	return len(l) == 0
}

func (l LineString) Is3D() bool {
	for _, p := range l {
		if !p.Is3D() {
			return false
		}
	}

	return true
}

func (l LineString) Equals(another Geometry) bool {
	panic("implement me")
}

func (l LineString) Disjoint(another Geometry) bool {
	panic("implement me")
}

func (l LineString) Intersects(another Geometry) bool {
	panic("implement me")
}

func (l LineString) Touches(another Geometry) bool {
	panic("implement me")
}

func (l LineString) Crosses(another Geometry) bool {
	panic("implement me")
}

func (l LineString) Within(another Geometry) bool {
	panic("implement me")
}

func (l LineString) Contains(another Geometry) bool {
	panic("implement me")
}

func (l LineString) Overlaps(another Geometry) bool {
	panic("implement me")
}

func (l LineString) Distance(another Geometry) float64 {
	panic("implement me")
}

func (l LineString) Buffer(distance float64) Geometry {
	panic("implement me")
}

func (l LineString) ConvexHull() Geometry {
	panic("implement me")
}

func (l LineString) Intersection(another Geometry) Geometry {
	panic("implement me")
}

func (l LineString) Union(another Geometry) Geometry {
	panic("implement me")
}

func (l LineString) Difference(another Geometry) Geometry {
	panic("implement me")
}

/* CURVE */
func (l LineString) Length() float64 {
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

func (l LineString) StartPoint() Point {
	return l[0]
}

func (l LineString) EndPoint() Point {
	return l[len(l)-1]
}

func (l LineString) IsClosed() bool {
	start := l.StartPoint()
	end := l.EndPoint()

	return reflect.DeepEqual(start, end)
}
