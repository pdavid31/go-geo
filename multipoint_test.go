package geo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mp = NewMultiPoint(p, p2)

func TestMultiPoint_GeometryType(t *testing.T) {
	assert.Equal(t, "MultiPoint", mp.GeometryType(), "MultiPoint GeometryType failed")
}

func TestMultiPoint_SRID(t *testing.T) {
	assert.Equal(t, 4326, mp.SRID(), "MultiPoint SRID failed")
}

func TestMultiPoint_AsText(t *testing.T) {
	asText := fmt.Sprintf("MULTIPOINT (%f %f %f, %f %f %f)", p.Lat(), p.Lon(), p.Z(), p2.Lat(), p2.Lon(), p2.Z())
	assert.Equal(t, asText, mp.AsText(), "MultiPoint AsText failed")
}

func TestMultiPoint_IsEmpty(t *testing.T) {
	assert.False(t, mp.IsEmpty(), "MultiPoint IsEmpty failed")
}

func TestMultiPoint_Is3D(t *testing.T) {
	assert.True(t, mp.Is3D(), "MultiPoint Is3D failed")
}

func TestMultiPoint_NumGeometries(t *testing.T) {
	assert.Equal(t, 2, len(mp), "MultiPoint NumGeometries failed")
}

func TestMultiPoint_GeometryN(t *testing.T) {
	assert.Equal(t, p, mp.GeometryN(0), "MultiPoint GeometryN failed")
}
