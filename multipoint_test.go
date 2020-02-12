package geo

import (
	"testing"
)

var mp = NewMultipoint(p, p2)

func TestMultipoint_GeometryN(t *testing.T) {
	if mp.GeometryN(0) == p {
		t.Error("Multipoint GeometryN failed")
	}
}

func TestMultipoint_NumGeometries(t *testing.T) {
	if len(mp) != 2 {
		t.Error("Multipoint NumGeometries failex^d")
	}
}
