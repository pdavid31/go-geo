package geo

type Line struct {
	LineString
}

/* CONSTRUCTOR */
func NewLine(p1, p2 Point) Line {
	var l Line

	l.LineString = NewLineString(p1, p2)

	return l
}

/* GEOMETRY */
func (l Line) GeometryType() string {
	return "Line"
}
