package geo

import "reflect"

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
// TODO: implement Geometry interface

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
