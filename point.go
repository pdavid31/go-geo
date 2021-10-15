package geo

import (
	"fmt"
	"math"
	"strings"
)

// Point type
type Point struct {
	x float64
	y float64
	z float64
	m float64
}

/* CONSTRUCTOR */

// NewPoint Point constructor
func NewPoint(params ...float64) Point {
	var p Point

	p.x = NullValue
	p.y = NullValue
	p.z = NullValue
	p.m = NullValue

	if len(params) < 2 {
		return p
	}

	p.x = params[0]
	p.y = params[1]

	if len(params) > 2 {
		p.z = params[2]
	}

	if len(params) > 3 {
		p.m = params[3]
	}

	return p
}

/* GEOMETRY */

// Lat function returns the latitude of the Geometry
func (p Point) Lat() float64 {
	return p.x
}

// Lon function returns the longitude of the Geometry
func (p Point) Lon() float64 {
	return p.y
}

// Alt function returns the altitude of the Geometry
func (p Point) Alt() float64 {
	return p.z
}

// GeometryType returns GeometryType as a string
func (p Point) GeometryType() string {
	return "Point"
}

// SRID returns the SRID as an integer
func (p Point) SRID() int {
	return 4326
}

// Envelope returns the Envelope of the geometry
func (p Point) Envelope() Geometry {
	// TODO: implement
	return nil
}

// AsText returns the string representation of the geometry
func (p Point) AsText() string {
	// example POINT ZM (1 1 5 60)
	rep := strings.ToUpper(p.GeometryType()) + " "

	if p.Is3D() {
		rep += "Alt"
	}

	if p.m != NullValue {
		rep += "M"
	}

	lastCharacter := rep[len(rep)-1:]
	if lastCharacter != " " {
		rep += " "
	}

	rep += "("

	rep += p.ToString()

	rep += ")"

	return rep
}

// ToString returns the internal string representation of the geometry
func (p Point) ToString() string {
	x := fmt.Sprintf("%f", p.x)
	y := fmt.Sprintf("%f", p.y)

	rep := x + " " + y

	if p.Is3D() {
		z := fmt.Sprintf("%f", p.z)
		rep += " " + z
	}

	if p.m != NullValue {
		m := fmt.Sprintf("%f", p.m)
		rep += " " + m
	}

	return rep
}

// IsEmpty returns a bool indicating if the geometry is empty
func (p Point) IsEmpty() bool {
	return p.x == NullValue || p.y == NullValue
}

// Is3D returns a bool indicating if the geometry is three dimensional
func (p Point) Is3D() bool {
	return p.z != NullValue
}

// Equals returns true if the geometry is equal to a given other
func (p Point) Equals(another Geometry) bool {
	// check if given geometry is a point
	switch another.(type) {
	case Point:
		return p.Lat() == another.Lat() && p.Lon() == another.Lon() && p.Alt() == another.Alt()
	default:
		return false
	}
}

// Disjoint returns true if the geometry is disjoint to a given other
func (p Point) Disjoint(another Geometry) bool {
	return !p.Intersects(another)
}

// Intersects returns true if the geometry intersects another given geometry
func (p Point) Intersects(another Geometry) bool {
	switch another.(type) {
	case Point, Curve:
		return p.Distance(another) == 0
	case Polygon:
		return p.Within(another)
	case GeometryCollection:
		geom := another.(GeometryCollection)

		numGeoms := geom.NumGeometries()
		for i := 0; i < numGeoms; i++ {
			if p.Intersects(geom.GeometryN(i)) {
				return true
			}
		}

		return false
	default:
		return false
	}
}

// Touches returns true if the geometry touches another given geometry
func (p Point) Touches(another Geometry) bool {
	panic("implement me")
}

// Crosses returns true if the geometry crosses another given geometry
func (p Point) Crosses(another Geometry) bool {
	panic("implement me")
}

// Within returns true if the geometry is within another given geometry
func (p Point) Within(another Geometry) bool {
	return another.Contains(p)
}

// Contains returns true if the geometry contains another given geometry
func (p Point) Contains(another Geometry) bool {
	panic("implement me")
}

// Overlaps returns true if the geometry overlaps another given geometry
func (p Point) Overlaps(another Geometry) bool {
	panic("implement me")
}

// Distance calculates the distane to another given geometry
func (p Point) Distance(another Geometry) float64 {
	switch another.(type) {
	case Point:
		p2 := another.(Point)

		// euclidean distance
		diffX := p2.Lat() - p.Lat()
		diffY := p2.Lon() - p.Lon()
		diffZ := p2.Alt() - p.Alt()

		return math.Sqrt(math.Pow(diffX, 2) + math.Pow(diffY, 2) + math.Pow(diffZ, 2))
	case Curve:
		var ls LineString

		switch another.(type) {
		case Line:
			l := another.(Line)
			ls = l.LineString
		case LinearRing:
			lr := another.(LinearRing)
			ls = lr.LineString
		default:
			ls = another.(LineString)
		}

		minDist := math.MaxFloat64

		numPoints := ls.NumPoints()
		for i := 0; i < numPoints-1; i++ {
			dist := distancePointToSegment(p, ls.PointN(i), ls.PointN(i+1))

			if dist < minDist {
				minDist = dist
			}
		}

		return minDist
	case Polygon:
		pg := another.(Polygon)

		// TODO: return 0 if point within polygon?

		minDist := math.MaxFloat64

		for _, lr := range pg {
			dist := p.Distance(lr)

			if dist < minDist {
				minDist = dist
			}
		}

		return minDist
	case GeometryCollection:
		gm := another.(GeometryCollection)

		minDist := math.MaxFloat64

		numGeoms := gm.NumGeometries()
		for i := 0; i < numGeoms; i++ {
			dist := p.Distance(gm.GeometryN(i))

			if dist < minDist {
				minDist = dist
			}
		}

		return minDist
	default:
		return 0
	}
}

// Buffer returns the Buffer with the given distance around the geometry
func (p Point) Buffer(distance float64) Geometry {
	var points []Point
	angle := 10

	lat := p.Lat()
	lon := p.Lon()
	z := p.Alt()

	for i := 0; i < 360; i += angle {
		rad := degToRad(float64(i))
		x := lat + distance*math.Cos(rad)
		y := lon + distance*math.Sin(rad)

		points = append(points, NewPoint(x, y, z))
	}

	return NewPolygon(NewLinearRing(points...))
}

// ConvexHull returns the ConvexHull containing the geometry
func (p Point) ConvexHull() Geometry {
	panic("implement me")
}

// Intersection returns the Intersection of the geometry and another given geometry
func (p Point) Intersection(another Geometry) Geometry {
	panic("implement me")
}

// Union returns the Union of the geometry and another given geometry
func (p Point) Union(another Geometry) Geometry {
	panic("implement me")
}

// Difference returns the Difference of the geometry and another given geometry
func (p Point) Difference(another Geometry) Geometry {
	panic("implement me")
}

/* INTERFACE GUARD */
var (
	_ Geometry = (*Point)(nil)
)
