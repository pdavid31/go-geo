package geo

import (
	"math"
	"testing"
)

const lat float64 = 53.559073
const lon float64 = 13.264618

var p = NewPoint(lat, lon)

const offset float64 = 1.0

const lat2 = lat + offset
const lon2 = lon + offset

var p2 = NewPoint(lat2, lon2)

func TestNewPoint(t *testing.T) {
	tempPoint := NewPoint(lat, lon)
	if tempPoint.Lat() != lat || tempPoint.Lon() != lon {
		t.Error("Point creation failed")
	}
}

func TestPoint_AzimuthTo(t *testing.T) {
	// TODO: calculate correct azimuth
	correctAzimuth := 30.00696084711074

	azimuth := p.AzimuthTo(p2)
	if azimuth != correctAzimuth {
		t.Errorf("Azimuth calculation failed - expected: %f, got: %f", azimuth, correctAzimuth)
	}
}

func TestPoint_Distance(t *testing.T) {
	correctDistance := math.Sqrt(offset * 2)

	distance := p.Distance(p2)
	if distance != correctDistance {
		t.Errorf("Distance calculation failed - expected: %f, got: %f", correctDistance, distance)
	}
}

func TestPoint_Near(t *testing.T) {
	if !p.Near(p2, 5) || p.Near(p2, 1) {
		t.Error("Distance check failed")
	}
}
