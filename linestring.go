package geo

type Linestring []Point

func NewLinestring(points ...*Point) Linestring {
	ls := Linestring{}

	for _, p := range points {
		ls = append(ls, *p)
	}

	return ls
}

// TODO: implement Geometry interface

func (l Linestring) NumGeometries() int {
	return len(l)
}

func (l Linestring) GeometryN(n int) Geometry {
	return &l[n]
}

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

	// TODO: update comparison
	if start == end {
		return true
	}

	return false
}
