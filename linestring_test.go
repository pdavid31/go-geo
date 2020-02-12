package geo

import (
	"math"
	"reflect"
	"testing"
)

var ls = NewLinestring(p, p2)

func TestLinestring_NumGeometries(t *testing.T) {
	if ls.NumGeometries() != 2 {
		t.Error("Linestring NumGeometries failed")
	}
}

func TestLinestring_GeometryN(t *testing.T) {
	if !reflect.DeepEqual(ls.GeometryN(0), p) {
		t.Error("Linestring GeometryN failed")
	}
}

func TestLinestring_Length(t *testing.T) {
	correctLength := math.Sqrt(offset * 2)
	if ls.Length() != correctLength {
		t.Error("Linestring Length failed")
	}
}

func TestLinestring_StartPoint(t *testing.T) {
	if !reflect.DeepEqual(ls.StartPoint(), p) {
		t.Error("Linestring StartPoint failed")
	}
}

func TestLinestring_EndPoint(t *testing.T) {
	if !reflect.DeepEqual(ls.EndPoint(), p2) {
		t.Error("Linestring StartPoint failed")
	}
}
