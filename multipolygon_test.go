package geo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mpoly = NewMultiPolygon(poly)

func TestMultiPolygon_GeometryType(t *testing.T) {
	assert.Equal(t, "MultiPolygon", mpoly.GeometryType(), "MultiPolygon GeometryType failed")
}

func TestMultiPolygon_SRID(t *testing.T) {
	assert.Equal(t, 4326, mpoly.SRID(), "MultiPolygon SRID failed")
}

func TestMultiPolygon_AsText(t *testing.T) {
	asText := fmt.Sprintf("MULTIPOLYGON (((%f %f %f, %f %f %f, %f %f %f)))", p.Lat(), p.Lon(), p.Alt(), p2.Lat(), p2.Lon(), p2.Alt(), p.Lat(), p.Lon(), p.Alt())
	assert.Equal(t, asText, mpoly.AsText(), "MultiPoint AsText failed")
}

func TestMultiPolygon_IsEmpty(t *testing.T) {
	assert.False(t, mpoly.IsEmpty(), "MultiPolygon IsEmpty failed")
}

func TestMultiPolygon_Is3D(t *testing.T) {
	assert.True(t, mpoly.Is3D(), "MultiPolygon Is3D failed")
}
