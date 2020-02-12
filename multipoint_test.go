package geo

import (
	"fmt"
	"reflect"
	"testing"
)

var mp = NewMultipoint(p, p2)

func TestMultipoint_GeometryType(t *testing.T) {
	if mp.GeometryType() != "Multipoint" {
		t.Error("Multipoint GeometryType failed")
	}
}

func TestMultipoint_SRID(t *testing.T) {
	if mp.SRID() != 4326 {
		t.Error("Multipoint SRID failed")
	}
}

func TestMultipoint_AsText(t *testing.T) {
	if mp.AsText() != fmt.Sprintf("MULTIPOINT (%f %f %f, %f %f %f)", p.x, p.y, p.z, p2.x, p2.y, p2.z) {
		t.Error("Multipoint AsText failed")
	}
}

func TestMultipoint_IsEmpty(t *testing.T) {
	if mp.IsEmpty() {
		t.Error("Multipoint IsEmpty failed")
	}
}

func TestMultipoint_Is3D(t *testing.T) {
	if !mp.Is3D() {
		t.Error("Multipoint Is3D failed")
	}
}

func TestMultipoint_NumGeometries(t *testing.T) {
	if len(mp) != 2 {
		t.Error("Multipoint NumGeometries failex^d")
	}
}

func TestMultipoint_GeometryN(t *testing.T) {
	if !reflect.DeepEqual(mp.GeometryN(0), p) {
		t.Error("Multipoint GeometryN failed")
	}
}
