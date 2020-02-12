package geo

type Line struct {
	Linestring
}

/* CONSTRUCTOR */
func NewLine(p1, p2 *Point) Line {
	l := Line{NewLinestring(p1, p2)}

	return l
}
