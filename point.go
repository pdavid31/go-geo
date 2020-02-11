package geo

import "math"

const (
	degToRad = math.Pi / 180.0
	radToDeg = 180.0 / math.Pi
)

type Point struct {
	lat float64
	lon float64
}

// NewPoint creates a new Point object
func NewPoint(Lat, Lon float64) *Point {
	p := new(Point)
	p.lat = Lat
	p.lon = Lon

	return p
}

// Lat getter
func (p *Point) Lat() float64 {
	return p.lat
}

// Lon getter
func (p *Point) Lon() float64 {
	return p.lon
}

// BearingTo calculates the azimuth
func (p *Point) AzimuthTo(p2 *Point) float64 {
	dLon := (p2.Lon() - p.Lon()) * degToRad

	lat1 := p.Lat() * degToRad
	lat2 := p2.Lat() * degToRad

	y := math.Sin(dLon) * math.Cos(lat2)
	x := math.Cos(lat1)*math.Sin(lat2) - math.Sin(lat1)*math.Cos(lat2)*math.Cos(dLon)

	azimuth := math.Atan2(y, x) * radToDeg

	return azimuth
}

// Distance calculates the distance between p and p2
func (p *Point) Distance(p2 *Point) float64 {
	distance := math.Sqrt(math.Pow(p.lat-p2.lat, 2) + math.Pow(p.lon-p2.lon, 2))

	return distance
}

// Near checks if p2 is in the given distance of p
func (p *Point) Near(p2 *Point, distance float64) bool {
	d := p.Distance(p2)

	if d > distance {
		return false
	}

	return true
}
