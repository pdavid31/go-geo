package geo

import (
	"math"
	"testing"
)

var mls = NewMultiLineString(ls)

func TestMultiLineString_NumGeometries(t *testing.T) {
	if mls.NumGeometries() != 1 {
		t.Error("MultiLineString NumGeometries failed")
	}
}

// TODO: compare Arrays of struct
//func TestMultiLineString_GeometryN(t *testing.T) {
//	fmt.Println(mls.GeometryN(0))
//	fmt.Println(ls)
//
//	if reflect.DeepEqual(mls.GeometryN(0), ls) {
//		t.Error("MultiLineString GeometryN failed")
//	}
//}

func TestMultiLineString_Length(t *testing.T) {
	correctLength := math.Sqrt(offset * 3)
	if mls.Length() != correctLength {
		t.Error("MultiLineString Length failed")
	}
}

func TestMultiLineString_IsClosed(t *testing.T) {
	if mls.IsClosed() {
		t.Error("MultiLineString IsClosed failed")
	}
}
