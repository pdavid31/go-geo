package geo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var poly = NewPolygon(lr)

func TestPolygon_GeometryType(t *testing.T) {
	assert.Equal(t, "Polygon", poly.GeometryType(), "Polygon GeometryType failed")
}

func TestPolygon_SRID(t *testing.T) {
	assert.Equal(t, 4326, poly.SRID(), "Polygon SRID failed")
}

func TestPolygon_AsText(t *testing.T) {
	asText := fmt.Sprintf("POLYGON ((%f %f %f, %f %f %f, %f %f %f))", p.Lat(), p.Lon(), p.Alt(), p2.Lat(), p2.Lon(), p2.Alt(), p.Lat(), p.Lon(), p.Alt())
	assert.Equal(t, asText, poly.AsText(), "Polygon AsText failed")
}

func TestPolygon_IsEmpty(t *testing.T) {
	assert.False(t, poly.IsEmpty(), "Polygon IsEmpty failed")
}

func TestPolygon_Is3D(t *testing.T) {
	assert.True(t, poly.Is3D(), "Polygon Is3D failed")
}

func TestPolygon_Equals(t *testing.T) {
	assert.True(t, poly.Equals(poly), "Polygon Equals failed")
}
