package geo

import (
	"testing"
)

var lr = NewLinearRing(p, p2)

func TestLinearRing_GeometryType(t *testing.T) {
	if lr.GeometryType() != "LinearRing" {
		t.Error("LinearRing GeometryType failed")
	}
}

func TestLinearRing_IsClosed(t *testing.T) {
	if !lr.IsClosed() {
		t.Error("LinearRing IsClosed failed")
	}
}
