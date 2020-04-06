package geo

import (
	"math"
	"reflect"
	"strings"
)

type LineString []Point

/* CONSTRUCTOR */
func NewLineString(points ...Point) LineString {
	var ls LineString

	for _, p := range points {
		ls = append(ls, p)
	}

	return ls
}

/* GEOMETRY */
func (l LineString) Lat() float64 {
	// see: https://www.ae.msstate.edu/vlsm/shape/centroid_of_a_line/example_1.htm
	topSum := 0.0
	bottomSum := 0.0

	for i := 0; i < len(l)-1; i++ {
		firstPoint := l[i]
		successorPoint := l[i+1]

		length := firstPoint.Distance(successorPoint)
		mid := (firstPoint.Lat() + successorPoint.Lat()) / 2

		topSum += length * mid
		bottomSum += length
	}

	x := topSum / bottomSum
	return x
}

func (l LineString) Lon() float64 {
	topSum := 0.0
	bottomSum := 0.0

	for i := 0; i < len(l)-1; i++ {
		firstPoint := l[i]
		successorPoint := l[i+1]

		length := firstPoint.Distance(successorPoint)
		mid := (firstPoint.Lon() + successorPoint.Lon()) / 2

		topSum += length * mid
		bottomSum += length
	}

	y := topSum / bottomSum
	return y
}

func (l LineString) Z() float64 {
	topSum := 0.0
	bottomSum := 0.0

	for i := 0; i < len(l)-1; i++ {
		firstPoint := l[i]
		successorPoint := l[i+1]

		length := firstPoint.Distance(successorPoint)
		mid := (firstPoint.Z() + successorPoint.Z()) / 2

		topSum += length * mid
		bottomSum += length
	}

	z := topSum / bottomSum
	return z
}

func (l LineString) GeometryType() string {
	return "LineString"
}

func (l LineString) SRID() int {
	return 4326
}

func (l LineString) Envelope() Geometry {
	panic("implement me")
}

func (l LineString) AsText() string {
	rep := strings.ToUpper(l.GeometryType()) + " "

	rep += l.ToString()

	return rep
}

func (l LineString) ToString() string {
	rep := "("

	lastIndex := len(l) - 2
	for i, p := range l {
		rep += p.ToString()

		if i <= lastIndex {
			rep += ", "
		}
	}

	rep += ")"
	return rep
}

func (l LineString) IsEmpty() bool {
	return len(l) == 0
}

func (l LineString) Is3D() bool {
	for _, p := range l {
		if !p.Is3D() {
			return false
		}
	}

	return true
}

func (l LineString) Equals(another Geometry) bool {
	// check if given geometry is a point
	switch another.(type) {
	case LineString, Line, LinearRing:
		// TODO: get rid of reflect
		// TODO: add independence of order
		return reflect.DeepEqual(l, another)
	default:
		return false
	}
}

func (l LineString) Disjoint(another Geometry) bool {
	panic("implement me")
}

func (l LineString) Intersects(another Geometry) bool {
	panic("implement me")
}

func (l LineString) Touches(another Geometry) bool {
	panic("implement me")
}

func (l LineString) Crosses(another Geometry) bool {
	panic("implement me")
}

func (l LineString) Within(another Geometry) bool {
	panic("implement me")
}

func (l LineString) Contains(another Geometry) bool {
	panic("implement me")
}

func (l LineString) Overlaps(another Geometry) bool {
	panic("implement me")
}

func (l LineString) Distance(another Geometry) float64 {
	switch another.(type) {
	case Point:
		p := another.(Point)

		return p.Distance(l)
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

		numPoints := l.NumPoints()
		numPoints2 := ls.NumPoints()

		for i := 0; i < numPoints-1; i++ {
			for j := 0; j < numPoints2-1; j++ {
				dist := distanceSegmentToSegment(l.PointN(i), l.PointN(i+1), ls.PointN(i), ls.PointN(i+1))

				if dist < minDist {
					minDist = dist
				}
			}
		}

		return minDist
	case Polygon:
		pg := another.(Polygon)

		minDist := math.MaxFloat64

		for _, lr := range pg {
			dist := l.Distance(lr)

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
			dist := l.Distance(gm.GeometryN(i))

			if dist < minDist {
				minDist = dist
			}
		}

		return minDist
	default:
		return 0
	}
}

func (l LineString) Buffer(distance float64) Geometry {
	panic("implement me")
}

func (l LineString) ConvexHull() Geometry {
	panic("implement me")
}

func (l LineString) Intersection(another Geometry) Geometry {
	panic("implement me")
}

func (l LineString) Union(another Geometry) Geometry {
	panic("implement me")
}

func (l LineString) Difference(another Geometry) Geometry {
	panic("implement me")
}

/* CURVE */
func (l LineString) Length() float64 {
	length := 0.0
	lastIndex := len(l) - 1

	for i := range l {
		if i >= lastIndex {
			break
		}

		current := l[i]
		successor := l[i+1]
		length += current.Distance(successor)
	}

	return length
}

func (l LineString) StartPoint() Point {
	return l[0]
}

func (l LineString) EndPoint() Point {
	return l[len(l)-1]
}

func (l LineString) IsClosed() bool {
	start := l.StartPoint()
	end := l.EndPoint()

	return start.Equals(end)
}

/* LINESTRING */
func (l LineString) NumPoints() int {
	return len(l)
}

func (l LineString) PointN(n int) Point {
	return l[n]
}
