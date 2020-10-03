package geo

import "math"

const (
	// NullValue is the float64 value describing not set coordinates
	NullValue float64 = -math.MaxFloat64

	lat float64 = 0
	lon float64 = 0
	z   float64 = 0

	offset float64 = 1.0

	lat2 = lat + offset
	lon2 = lon + offset
	z2   = z + offset
)

var p = NewPoint(lat, lon, z)
var p2 = NewPoint(lat2, lon2, z2)
