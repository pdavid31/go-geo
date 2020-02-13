package geo

import (
	"testing"
)

var mpoly = NewMultiPolygon(poly)

func TestMultiPolygon_GeometryType(t *testing.T) {
	if mpoly.GeometryType() != "MultiPolygon" {
		t.Error("MultiPolygon GeometryType failed")
	}
}

func TestMultiPolygon_SRID(t *testing.T) {
	if mpoly.SRID() != 4326 {
		t.Error("MultiPolygon SRID failed")
	}
}

func TestMultiPolygon_IsEmpty(t *testing.T) {
	if mpoly.IsEmpty() {
		t.Error("MultiPolygon IsEmpty failed")
	}
}

func TestMultiPolygon_Is3D(t *testing.T) {
	if !mpoly.Is3D() {
		t.Error("MultiPolygon Is3D failed")
	}
}
