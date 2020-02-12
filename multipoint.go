package geo

type Multipoint []Point

func NewMultipoint(points ...*Point) Multipoint {
	mp := Multipoint{}

	for _, p := range points {
		mp = append(mp, *p)
	}

	return mp
}

func (m Multipoint) NumGeometries() int {
	return len(m)
}

func (m Multipoint) GeometryN(n int) Geometry {
	return &m[n]
}
