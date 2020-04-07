package geo

import (
	"fmt"
	"math"
	"strings"
)

type Point struct {
	x float64
	y float64
	z float64
	m float64
}

/* CONSTRUCTOR */
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
func (p Point) Lat() float64 {
	return p.x
}

func (p Point) Lon() float64 {
	return p.y
}

func (p Point) Z() float64 {
	return p.z
}

func (p Point) GeometryType() string {
	return "Point"
}

func (p Point) SRID() int {
	return 4326
}

func (p Point) Envelope() Geometry {
	// TODO: implement
	return nil
}

func (p Point) AsText() string {
	// example POINT ZM (1 1 5 60)
	rep := strings.ToUpper(p.GeometryType()) + " "

	if p.Is3D() {
		rep += "Z"
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

func (p Point) IsEmpty() bool {
	return p.x == NullValue || p.y == NullValue
}

func (p Point) Is3D() bool {
	return p.z != NullValue
}

func (p Point) Equals(another Geometry) bool {
	// check if given geometry is a point
	switch another.(type) {
	case Point:
		return p.Lat() == another.Lat() && p.Lon() == another.Lon() && p.Z() == another.Z()
	default:
		return false
	}
}

func (p Point) Disjoint(another Geometry) bool {
	return !p.Intersects(another)
}

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

func (p Point) Touches(another Geometry) bool {
	panic("implement me")
}

func (p Point) Crosses(another Geometry) bool {
	panic("implement me")
}

func (p Point) Within(another Geometry) bool {
	return another.Contains(p)
}

func (p Point) Contains(another Geometry) bool {
	panic("implement me")
}

func (p Point) Overlaps(another Geometry) bool {
	panic("implement me")
}

func (p Point) Distance(another Geometry) float64 {
	switch another.(type) {
	case Point:
		p2 := another.(Point)

		// euclidean distance
		diffX := p2.Lat() - p.Lat()
		diffY := p2.Lon() - p.Lon()
		diffZ := p2.Z() - p.Z()

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

func (p Point) Buffer(distance float64) Geometry {
	var points []Point
	angle := 20

	lat := p.Lat()
	lon := p.Lon()
	z := p.Z()

	for i := 0; i < 360; i += angle {
		rad := degToRad(float64(i))
		x := lat + distance*math.Cos(rad)
		y := lon + distance*math.Sin(rad)

		points = append(points, NewPoint(x, y, z))
	}

	return NewPolygon(NewLinearRing(points...))
}

func (p Point) ConvexHull() Geometry {
	panic("implement me")
}

func (p Point) Intersection(another Geometry) Geometry {
	panic("implement me")
}

func (p Point) Union(another Geometry) Geometry {
	panic("implement me")
}

func (p Point) Difference(another Geometry) Geometry {
	panic("implement me")
}
