package geo

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ls = NewLineString(p, p2)

func TestLineString_Lat(t *testing.T) {
	assert.Equal(t, (p.Lat()+p2.Lat())/2, ls.Lat(), "LineString Lat failed")
}

func TestLineString_Lon(t *testing.T) {
	assert.Equal(t, (p.Lon()+p2.Lon())/2, ls.Lon(), "LineString Lon failed")
}

func TestLineString_Z(t *testing.T) {
	assert.Equal(t, (p.Alt()+p2.Alt())/2, ls.Alt(), "LineString Alt failed")
}

func TestLineString_GeometryType(t *testing.T) {
	assert.Equal(t, "LineString", ls.GeometryType(), "LineString GeometryType failed")
}

func TestLineString_SRID(t *testing.T) {
	assert.Equal(t, 4326, ls.SRID(), "LineString SRID failed")
}

func TestLineString_AsText(t *testing.T) {
	asText := fmt.Sprintf("LINESTRING (%f %f %f, %f %f %f)", p.Lat(), p.Lon(), p.Alt(), p2.Lat(), p2.Lon(), p2.Alt())
	assert.Equal(t, asText, ls.AsText(), "LineString AsText failed")
}

func TestLineString_IsEmpty(t *testing.T) {
	assert.False(t, ls.IsEmpty(), "LineString IsEmpty failed")
}

func TestLineString_Is3D(t *testing.T) {
	assert.True(t, ls.Is3D(), "LineString Is3D failed")
}

func TestLineString_Equals(t *testing.T) {
	assert.True(t, ls.Equals(ls), "LineString Equals failed")
	assert.True(t, ls.Equals(l.LineString), "LineString Equals failed")
	assert.False(t, ls.Equals(lr.LineString), "LineString Equals failed")
}

func TestLineString_Distance(t *testing.T) {
	p3 := NewPoint(p.Lat()-offset, p.Lon(), 0)
	p4 := NewPoint(p.Lat()-offset, p.Lon()+offset, 0)

	dToP := ls.Distance(p3)
	assert.Equal(t, offset, dToP, "LineString Distance (to Point) failed - expected: %f, got: %f", offset, dToP)

	lineString := NewLineString(p3, p4)
	dToL := l.Distance(lineString)
	assert.Equal(t, offset, dToL, "LineString Distance (to LineString) failed - expected: %f, got: %f", offset, dToL)

	linearRing := NewLinearRing(p3, p4)
	polygon := NewPolygon(linearRing)
	dToP = p.Distance(polygon)
	assert.Equal(t, offset, dToP, "LineString Distance (to Polygon) failed - expected: %f, got: %f", offset, dToP)

	multiPoint := NewMultiPoint(p3, p4)
	dToMP := p.Distance(multiPoint)
	assert.Equal(t, offset, dToMP, "LineString Distance (to MultiPoint) failed - expected: %f, got: %f", offset, dToMP)

	multiLineString := NewMultiLineString(lineString)
	dToMLS := p.Distance(multiLineString)
	assert.Equal(t, offset, dToMLS, "LineString Distance (to MultiLineString) failed - expected: %f, got: %f", offset, dToMLS)

	multiPolygon := NewMultiPolygon(polygon)
	dToMPG := p.Distance(multiPolygon)
	assert.Equal(t, offset, dToMPG, "LineString Distance (to MultiPolygon) failed - expected: %f, got: %f", offset, dToMPG)
}

func TestLineString_Length(t *testing.T) {
	correctLength := math.Sqrt(offset * 3)
	assert.Equal(t, correctLength, ls.Length(), "LineString Length failed")

	ls2 := NewLineString(p, p2, p)
	correctLength = correctLength * 2
	assert.Equal(t, correctLength, ls2.Length(), "LineString IsClosed failed")
}

func TestLineString_StartPoint(t *testing.T) {
	assert.Equal(t, p, ls.StartPoint(), "LineString StartPoint failed")
}

func TestLineString_EndPoint(t *testing.T) {
	assert.Equal(t, p2, ls.EndPoint(), "LineString StartPoint failed")
}

func TestLineString_IsClosed(t *testing.T) {
	assert.False(t, ls.IsClosed(), "LineString IsClosed failed")

	ls2 := NewLineString(p, p2, p)
	assert.True(t, ls2.IsClosed(), "LineString IsClosed failed")
}

func TestLineString_NumPoints(t *testing.T) {
	assert.Equal(t, 2, ls.NumPoints(), "LineString NumPoints failed")
}

func TestLineString_PointN(t *testing.T) {
	assert.Equal(t, p, ls.PointN(0), "LineString PointN failed")
}
