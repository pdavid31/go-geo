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

func BenchmarkPoint_Equals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p.Equals(p2)
	}
}

func TestPoint_Disjoint(t *testing.T) {
	if p.Disjoint(p) || !p.Disjoint(p2) {
		t.Error("Point Disjoint failed")
	}
}

func TestPoint_Intersects(t *testing.T) {
	if !p.Intersects(p) || p.Intersects(p2) {
		t.Error("Point Intersects failed")
	}
}

func BenchmarkPoint_Intersects(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p.Intersects(p2)
	}
}

// TODO: test touches
// TODO: test crosses
// TODO: test within
// TODO: test contains
// TODO: test overlaps

func TestPoint_Distance(t *testing.T) {
	p3 := NewPoint(p.Lat(), p.Lon()+offset, 0)
	p4 := NewPoint(p.Lat()+offset, p.Lon()+offset, 0)

	if dToP := p.Distance(p3); dToP != offset {
		t.Errorf("Point Distance (to Point) failed - expected: %f, got: %f", offset, dToP)
	}

	lineString := NewLineString(p3, p4)
	if dToL := p.Distance(lineString); dToL != offset {
		t.Errorf("Point Distance (to LineString) failed - expected: %f, got: %f", offset, dToL)
	}

	linearRing := NewLinearRing(p3, p4)
	polygon := NewPolygon(linearRing)
	if dToP := p.Distance(polygon); dToP != offset {
		t.Errorf("Point Distance (to Polygon) failed - expected: %f, got: %f", offset, dToP)
	}

	multiPoint := NewMultiPoint(p3, p4)
	if dToMP := p.Distance(multiPoint); dToMP != offset {
		t.Errorf("Point Distance (to MultiPoint) failed - expected: %f, got: %f", offset, dToMP)
	}

	multiLineString := NewMultiLineString(lineString)
	if dToMLS := p.Distance(multiLineString); dToMLS != offset {
		t.Errorf("Point Distance (to MultiLineString) failed - expected: %f, got: %f", offset, dToMLS)
	}

	multiPolygon := NewMultiPolygon(polygon)
	if dToMPG := p.Distance(multiPolygon); dToMPG != offset {
		t.Errorf("Point Distance (to MultiPolygon) failed - expected: %f, got: %f", offset, dToMPG)
	}
}

func BenchmarkPoint_DistancePoint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p.Distance(p2)
	}
}

func BenchmarkPoint_DistanceCurve(b *testing.B) {
	p3 := NewPoint(p.Lat(), p.Lon()+offset, 0)
	p4 := NewPoint(p.Lat()+offset, p.Lon()+offset, 0)
	linearRing := NewLinearRing(p3, p4)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		p.Distance(linearRing)
	}
}

func BenchmarkPoint_DistancePolygon(b *testing.B) {
	p3 := NewPoint(p.Lat(), p.Lon()+offset, 0)
	p4 := NewPoint(p.Lat()+offset, p.Lon()+offset, 0)
	linearRing := NewLinearRing(p3, p4)
	polygon := NewPolygon(linearRing)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		p.Distance(polygon)
	}
}

func BenchmarkPoint_DistanceGeometryCollection(b *testing.B) {
	p3 := NewPoint(p.Lat(), p.Lon()+offset, 0)
	p4 := NewPoint(p.Lat()+offset, p.Lon()+offset, 0)
	linearRing := NewLinearRing(p3, p4)
	polygon := NewPolygon(linearRing)
	multiPolygon := NewMultiPolygon(polygon)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		p.Distance(multiPolygon)
	}
}

func TestPoint_Buffer(t *testing.T) {
	distance := 10.0
	buffer := p.Buffer(distance)

	for _, lr := range buffer.(Polygon) {
		for _, bufferPoint := range lr.LineString {
			if math.Round(p.Distance(bufferPoint)) != distance {
				t.Error("Point Buffer failed")
			}
		}
	}
}

// TODO: test convexhull
// TODO: test intersection
// TODO: test union
// TODO: test difference
