package geo

import (
	"fmt"
	"math"
	"testing"
)

var mls = NewMultiLineString(ls)

func TestMultiLineString_GeometryType(t *testing.T) {
	if mls.GeometryType() != "MultiLineString" {
		t.Error("MultiLineString GeometryType failed")
	}
}

func TestMultiLineString_SRID(t *testing.T) {
	if mls.SRID() != 4326 {
		t.Error("MultiLineString SRID failed")
	}
}

func TestMultiLineString_AsText(t *testing.T) {
	if mls.AsText() != fmt.Sprintf("MULTILINESTRING ((%f %f %f, %f %f %f))", p.x, p.y, p.z, p2.x, p2.y, p2.z) {
		t.Error("MultiPoint AsText failed")
	}
}

func TestMultiLineString_IsEmpty(t *testing.T) {
	if mls.IsEmpty() {
		t.Error("MultiLineString IsEmpty failed")
	}
}

func TestMultiLineString_Is3D(t *testing.T) {
	if !mls.Is3D() {
		t.Error("MultiLineString Is3D failed")
	}
}

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
