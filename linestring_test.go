package geo

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

var ls = NewLineString(p, p2)

func TestLineString_GeometryType(t *testing.T) {
	if ls.GeometryType() != "LineString" {
		t.Error("LineString GeometryType failed")
	}
}

func TestLineString_SRID(t *testing.T) {
	if ls.SRID() != 4326 {
		t.Error("LineString SRID failed")
	}
}

func TestLineString_AsText(t *testing.T) {
	if ls.AsText() != fmt.Sprintf("LINESTRING (%f %f %f, %f %f %f)", p.x, p.y, p.z, p2.x, p2.y, p2.z) {
		t.Error("LineString AsText failed")
	}
}

func TestLineString_IsEmpty(t *testing.T) {
	if ls.IsEmpty() {
		t.Error("LineString IsEmpty failed")
	}
}

func TestLineString_Is3D(t *testing.T) {
	if !ls.Is3D() {
		t.Error("LineString Is3D failed")
	}
}

func TestLineString_Length(t *testing.T) {
	correctLength := math.Sqrt(offset * 2)
	if ls.Length() != correctLength {
		t.Error("LineString Length failed")
	}
}

func TestLineString_StartPoint(t *testing.T) {
	if !reflect.DeepEqual(ls.StartPoint(), p) {
		t.Error("LineString StartPoint failed")
	}
}

func TestLineString_EndPoint(t *testing.T) {
	if !reflect.DeepEqual(ls.EndPoint(), p2) {
		t.Error("LineString StartPoint failed")
	}
}

func TestLineString_IsClosed(t *testing.T) {
	if ls.IsClosed() {
		t.Error("LineString IsClosed failed")
	}

	ls2 := NewLineString(p, p2, p)
	if !ls2.IsClosed() {
		t.Error("LineString IsClosed failed")
	}
}

func TestLineString_NumPoints(t *testing.T) {
	if ls.NumPoints() != 2 {
		t.Error("LineString NumPoints failed")
	}
}

func TestLineString_PointN(t *testing.T) {
	if !reflect.DeepEqual(ls.PointN(0), p) {
		t.Error("LineString PointN failed")
	}
}
