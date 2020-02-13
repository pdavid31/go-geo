package geo

import (
	"fmt"
	"reflect"
	"testing"
)

var mp = NewMultiPoint(p, p2)

func TestMultiPoint_GeometryType(t *testing.T) {
	if mp.GeometryType() != "MultiPoint" {
		t.Error("MultiPoint GeometryType failed")
	}
}

func TestMultiPoint_SRID(t *testing.T) {
	if mp.SRID() != 4326 {
		t.Error("MultiPoint SRID failed")
	}
}

func TestMultiPoint_AsText(t *testing.T) {
	if mp.AsText() != fmt.Sprintf("MULTIPOINT (%f %f %f, %f %f %f)", p.x, p.y, p.z, p2.x, p2.y, p2.z) {
		t.Error("MultiPoint AsText failed")
	}
}

func TestMultiPoint_IsEmpty(t *testing.T) {
	if mp.IsEmpty() {
		t.Error("MultiPoint IsEmpty failed")
	}
}

func TestMultiPoint_Is3D(t *testing.T) {
	if !mp.Is3D() {
		t.Error("MultiPoint Is3D failed")
	}
}

func TestMultiPoint_NumGeometries(t *testing.T) {
	if len(mp) != 2 {
		t.Error("MultiPoint NumGeometries failex^d")
	}
}

func TestMultiPoint_GeometryN(t *testing.T) {
	if !reflect.DeepEqual(mp.GeometryN(0), p) {
		t.Error("MultiPoint GeometryN failed")
	}
}
