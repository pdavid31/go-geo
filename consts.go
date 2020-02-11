package geo

const (
	lat float64 = 53.559073
	lon float64 = 13.264618
	z   float64 = 20

	offset float64 = 1.0

	lat2 = lat + offset
	lon2 = lon + offset
	z2   = z + offset
)

var p = NewPoint(lat, lon, z)
var p2 = NewPoint(lat2, lon2, z2)
