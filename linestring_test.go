package geo

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

var ls = NewLinestring(p, p2)

func TestLinestring_GeometryType(t *testing.T) {
	if ls.GeometryType() != "Linestring" {
		t.Error("Linestring GeometryType failed")
	}
}

func TestLinestring_SRID(t *testing.T) {
	if ls.SRID() != 4326 {
		t.Error("Linestring SRID failed")
	}
}

func TestLinestring_AsText(t *testing.T) {
	if ls.AsText() != fmt.Sprintf("LINESTRING (%f %f %f, %f %f %f)", p.x, p.y, p.z, p2.x, p2.y, p2.z) {
		t.Error("Linestring AsText failed")
	}
}

func TestLinestring_IsEmpty(t *testing.T) {
	if ls.IsEmpty() {
		t.Error("Linestring IsEmpty failed")
	}
}

func TestLinestring_Is3D(t *testing.T) {
	if !ls.Is3D() {
		t.Error("Linestring Is3D failed")
	}
}

func TestLinestring_Length(t *testing.T) {
	correctLength := math.Sqrt(offset * 2)
	if ls.Length() != correctLength {
		t.Error("Linestring Length failed")
	}
}

func TestLinestring_StartPoint(t *testing.T) {
	if !reflect.DeepEqual(ls.StartPoint(), *p) {
		t.Error("Linestring StartPoint failed")
	}
}

func TestLinestring_EndPoint(t *testing.T) {
	if !reflect.DeepEqual(ls.EndPoint(), *p2) {
		t.Error("Linestring StartPoint failed")
	}
}

func TestLinestring_IsClosed(t *testing.T) {
	if ls.IsClosed() {
		t.Error("Linestring IsClosed failed")
	}

	ls2 := NewLinestring(p, p2, p)
	if !ls2.IsClosed() {
		t.Error("Linestring IsClosed failed")
	}
}
