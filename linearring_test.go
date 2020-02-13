package geo

import (
	"testing"
)

var lr = NewLinearring(p, p2)

func TestLinearring_GeometryType(t *testing.T) {
	if lr.GeometryType() != "Linearring" {
		t.Error("Linearring GeometryType failed")
	}
}

func TestLinearring_IsClosed(t *testing.T) {
	if !lr.IsClosed() {
		t.Error("Linearring IsClosed failed")
	}
}
