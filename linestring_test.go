package geo

import (
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
		t.Error("LineString GeometryN failed")
	}
}
