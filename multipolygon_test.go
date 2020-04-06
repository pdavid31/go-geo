package geo

import (
	"fmt"
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

func TestMultiPolygon_AsText(t *testing.T) {
	if mpoly.AsText() != fmt.Sprintf("MULTIPOLYGON (((%f %f %f, %f %f %f, %f %f %f)))", p.Lat(), p.Lon(), p.Z(), p2.Lat(), p2.Lon(), p2.Z(), p.Lat(), p.Lon(), p.Z()) {
		t.Error("MultiPoint AsText failed")
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
