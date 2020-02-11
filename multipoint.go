package geo

type Multipoint []Point

func (m Multipoint) NumGeometries() int {
	return len(m)
}

func (m Multipoint) GeometryN(n int) Geometry {
	return &m[n]
}
