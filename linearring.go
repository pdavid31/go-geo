package geo

import "reflect"

type Linearring struct {
	Linestring
}

/* CONSTRUCTOR */
func NewLinearring(points ...*Point) Linearring {
	firstPoint := points[0]
	lastPoint := points[len(points)-1]

	// check if ring is closed, else append first point to it
	if !reflect.DeepEqual(firstPoint, lastPoint) {
		points = append(points, firstPoint)
	}

	l := Linearring{NewLinestring(points...)}

	return l
}

/* GEOMETRY */
func (l Linearring) GeometryType() string {
	return "Linearring"
}
