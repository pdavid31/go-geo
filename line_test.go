package geo

import (
	"testing"
)

var l = NewLine(p, p2)

func TestLine_Is3D(t *testing.T) {
	if !l.Is3D() {
		t.Error("Line Is3D failed")
	}
}
