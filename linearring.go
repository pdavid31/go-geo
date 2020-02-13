package geo

import "reflect"

type LinearRing struct {
	LineString
}

/* CONSTRUCTOR */
func NewLinearRing(points ...Point) LinearRing {
	firstPoint := points[0]
	lastPoint := points[len(points)-1]

	// check if ring is closed, else append first point to it
	if !reflect.DeepEqual(firstPoint, lastPoint) {
		points = append(points, firstPoint)
	}

	l := LinearRing{NewLineString(points...)}

	return l
}

/* GEOMETRY */
func (l LinearRing) GeometryType() string {
	return "LinearRing"
}
