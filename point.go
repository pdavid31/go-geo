package geo

import "math"

type Point struct {
	Lat float64
	Lon float64
}

// NewPoint creates a new Point object
func NewPoint(lat float64, lon float64) *Point {
	return &Point{Lat: lat, Lon: lon}
}

// GeometryType returns subtype of geometry
func (p *Point) GeometryType() string {
	return "Point"
}

// SRID returns Spatial Reference System ID of the geometric object
func (p *Point) SRID() int {
	// TODO: implement
	return 4326
}

// Envelope returns the minimum bounding box for this geometric object
func (p *Point) Envelope() Geometry {
	// TODO: implement
	return nil
}

// AsText exports this geometric object to a specific Well-known Text Representation
func (p *Point) AsText() string {
	// TODO: implement
	return ""
}

// IsEmpty returns true if the geometric object is an empty Geometry
func (p *Point) IsEmpty() bool {
	return false
}

// Distance calculates the distance between p and p2
func (p *Point) Distance(p2 *Point) float64 {
	distance := math.Sqrt(math.Pow(p.Lat-p2.Lat, 2) + math.Pow(p.Lon-p2.Lon, 2))

	return distance
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

func (p *Point) Within(another Geometry) bool {
	panic("implement me")
}

func (p *Point) Contains(another Geometry) bool {
	panic("implement me")
}

func (p *Point) Overlaps(another Geometry) bool {
	panic("implement me")
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
