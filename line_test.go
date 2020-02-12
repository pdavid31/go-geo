package geo

import (
	"testing"
)

var l = NewLine(p, p2)

func TestLine_GeometryType(t *testing.T) {
	if l.GeometryType() != "Line" {
		t.Error("Line GeometryType failed")
	}
}
