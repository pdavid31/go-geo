package geo

import (
	"fmt"
	"math"
	"testing"
)

func TestNewPoint(t *testing.T) {
	if p.x != lat || p.y != lon || p.z != z {
		t.Error("Point constructor failed")
	}
}

func TestPoint_GeometryType(t *testing.T) {
	if p.GeometryType() != "Point" {
		t.Error("Point GeometryType failed")
	}
}

func TestPoint_SRID(t *testing.T) {
	if p.SRID() != 4326 {
		t.Error("Point SRID failed")
	}
}

// TODO: test envelope

func TestPoint_AsText(t *testing.T) {
	x := fmt.Sprintf("%f", lat)
	y := fmt.Sprintf("%f", lon)
	zStr := fmt.Sprintf("%f", z)

	if p.AsText() != "POINT Z ("+x+" "+y+" "+zStr+")" {
		t.Error("Point AsText failed")
	}
}

// TODO: test isempty

func TestPoint_Is3D(t *testing.T) {
	if !p.Is3D() {
		t.Error("Point Is3D failed")
	}
}

// TODO: test equals
// TODO: test disjoint
// TODO: test intersects
// TODO: test touches
// TODO: test crosses
// TODO: test within
// TODO: test contains
// TODO: test overlaps

func TestPoint_Distance(t *testing.T) {
	correctDistance := math.Sqrt(offset * 2)

	distance := p.Distance(p2)
	if distance != correctDistance {
		t.Errorf("Point Distance failed - expected: %f, got: %f", correctDistance, distance)
	}
}

// TODO: test buffer
// TODO: test convexhull
// TODO: test intersection
// TODO: test union
// TODO: test difference
