package geo

import (
	"testing"
)

var poly = NewPolygon(lr)

func TestPolygon_GeometryType(t *testing.T) {
	if poly.GeometryType() != "Polygon" {
		t.Error("Polygon GeometryType failed")
	}
}

func TestPolygon_SRID(t *testing.T) {
	if poly.SRID() != 4326 {
		t.Error("Polygon SRID failed")
	}
}

func TestPolygon_AsText(t *testing.T) {
	if poly.AsText() != "POLYGON ((0.000000 0.000000 0.000000, 1.000000 1.000000 1.000000, 0.000000 0.000000 0.000000))" {
		t.Error("Polygon AsText failed")
	}
}

func TestPolygon_IsEmpty(t *testing.T) {
	if poly.IsEmpty() {
		t.Error("Polygon IsEmpty failed")
	}
}

func TestPolygon_Is3D(t *testing.T) {
	if !poly.Is3D() {
		t.Error("Polygon Is3D failed")
	}
}
