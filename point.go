package geo

import (
	"math"
)

const (
	GeometryType = "point"
)

type Point struct {
	x float64
	y float64
	z float64
	m float64
}

// NewPoint creates a new Point object
func NewPoint(params ...float64) *Point {
	x := params[0]
	y := params[1]

	z := -5000.0
	if len(params) > 2 {
		z = params[2]
	}

	m := -5000.0
	if len(params) > 3 {
		m = params[3]
	}

	return &Point{x: x, y: y, z: z, m: m}
}

func (p *Point) GeometryType() string {
	return "Point"
}

func (p *Point) SRID() int {
	// TODO: implement
	return 4326
}

func (p *Point) Envelope() Geometry {
	// TODO: implement
	return nil
}

func (p *Point) AsText() string {
	// TODO: implement
	return ""
}

func (p *Point) IsEmpty() bool {
	// TODO: implement
	return false
}

func (p *Point) Is3D() bool {
	// TODO: implement
	return false
}

func (p *Point) Equals(another Geometry) bool {
	panic("implement me")
}

func (p *Point) Disjoint(another Geometry) bool {
	panic("implement me")
}

func (p *Point) Intersects(another Geometry) bool {
	panic("implement me")
}

func (p *Point) Touches(another Geometry) bool {
	panic("implement me")
}

func (p *Point) Crosses(another Geometry) bool {
	panic("implement me")
}

func (p *Point) Within(another Geometry) bool {
	panic("implement me")
}

func (p *Point) Contains(another Geometry) bool {
	panic("implement me")
}

func (p *Point) Overlaps(another Geometry) bool {
	panic("implement me")
}

// Distance calculates the distance between p and p2
func (p *Point) Distance(p2 *Point) float64 {
	distance := math.Sqrt(math.Pow(p.x-p2.x, 2) + math.Pow(p.y-p2.y, 2))

	return distance
}

func (p *Point) Buffer(distance float64) Geometry {
	panic("implement me")
}

func (p *Point) ConvexHull() Geometry {
	panic("implement me")
}

func (p *Point) Intersection(another Geometry) Geometry {
	panic("implement me")
}

func (p *Point) Union(another Geometry) Geometry {
	panic("implement me")
}

func (p *Point) Difference(another Geometry) Geometry {
	panic("implement me")
}
