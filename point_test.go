package geo

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint_GeometryType(t *testing.T) {
	assert.Equal(t, "Point", p.GeometryType(), "Point GeometryType failed")
}

func TestPoint_SRID(t *testing.T) {
	assert.Equal(t, 4326, p.SRID(), "Point SRID failed")
}

// TODO: test envelope

func TestPoint_AsText(t *testing.T) {
	asText := fmt.Sprintf("POINT Alt (%f %f %f)", p.Lat(), p.Lon(), p.Alt())
	assert.Equal(t, asText, p.AsText())
}

func TestPoint_IsEmpty(t *testing.T) {
	assert.False(t, p.IsEmpty(), "Point IsEmpty failed")

	p3 := NewPoint()
	assert.True(t, p3.IsEmpty(), "Point IsEmpty failed")
}

func TestPoint_Is3D(t *testing.T) {
	assert.True(t, p.Is3D(), "Point Is3D failed")
}

func TestPoint_Equals(t *testing.T) {
	assert.True(t, p.Equals(p), "Point Equals failed")
	assert.False(t, p.Equals(p2), "Point Equals failed")
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
	assert.False(t, !p.Intersects(p) || p.Intersects(p2), "Point Intersects failed")
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

	dToP := p.Distance(p3)
	assert.Equal(t, offset, dToP, "Point Distance (to Point) failed - expected: %f, got: %f", offset, dToP)

	lineString := NewLineString(p3, p4)
	dToL := p.Distance(lineString)
	assert.Equal(t, offset, dToL, "Point Distance (to LineString) failed - expected: %f, got: %f", offset, dToL)

	linearRing := NewLinearRing(p3, p4)
	polygon := NewPolygon(linearRing)
	dToP = p.Distance(polygon)
	assert.Equal(t, offset, dToP, "Point Distance (to Polygon) failed - expected: %f, got: %f", offset, dToP)

	multiPoint := NewMultiPoint(p3, p4)
	dToMP := p.Distance(multiPoint)
	assert.Equal(t, offset, dToMP, "Point Distance (to MultiPoint) failed - expected: %f, got: %f", offset, dToMP)

	multiLineString := NewMultiLineString(lineString)
	dToMLS := p.Distance(multiLineString)
	assert.Equal(t, offset, dToMLS, "Point Distance (to MultiLineString) failed - expected: %f, got: %f", offset, dToMLS)

	multiPolygon := NewMultiPolygon(polygon)
	dToMPG := p.Distance(multiPolygon)
	assert.Equal(t, offset, dToMPG)
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
			assert.Equal(t, distance, math.Round(p.Distance(bufferPoint)))
		}
	}
}

// TODO: test convexhull
// TODO: test intersection
// TODO: test union
// TODO: test difference
