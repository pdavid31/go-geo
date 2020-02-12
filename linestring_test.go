package geo

import (
	"math"
	"reflect"
	"testing"
)

var ls = NewLinestring(p, p2)

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
