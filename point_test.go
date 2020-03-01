package geo

import (
	"fmt"
	"math"
	"testing"
)

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
	if p.AsText() != fmt.Sprintf("POINT Z (%f %f %f)", p.Lat(), p.Lon(), p.Z()) {
		t.Error("Point AsText failed")
	}
}

func TestPoint_IsEmpty(t *testing.T) {
	if p.IsEmpty() {
		t.Error("Point IsEmpty failed")
	}

	p3 := NewPoint()
	if !p3.IsEmpty() {
		t.Error("Point IsEmpty failed")
	}
}

func TestPoint_Is3D(t *testing.T) {
	if !p.Is3D() {
		t.Error("Point Is3D failed")
	}
}

func TestPoint_Equals(t *testing.T) {
	if !p.Equals(p) {
		t.Error("Point Equals failed")
	}

	if p.Equals(p2) {
		t.Error("Point Equals failed")
	}
}

// TODO: test disjoint
// TODO: test intersects
// TODO: test touches
// TODO: test crosses
// TODO: test within
// TODO: test contains
// TODO: test overlaps

func TestPoint_Distance(t *testing.T) {
	correctDistance := math.Sqrt(offset * 3)

	if dToP := p.Distance(p2); dToP != correctDistance {
		t.Errorf("Point Distance (to Point) failed - expected: %f, got: %f", correctDistance, dToP)
	}

	lineString := NewLineString(NewPoint(p.Lat()-offset, p.Lon()+offset, 0), NewPoint(p.Lat()+offset, p.Lon()+offset, 0))
	if dToL := p.Distance(lineString); dToL != 1 {
		t.Errorf("Point Distance (to LineString) failed - expected: %f, got: %f", offset, dToL)
	}
}

// TODO: test buffer
// TODO: test convexhull
// TODO: test intersection
// TODO: test union
// TODO: test difference
