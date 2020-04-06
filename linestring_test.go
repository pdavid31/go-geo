package geo

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

var ls = NewLineString(p, p2)

func TestLineString_Lat(t *testing.T) {
	if ls.Lat() != (p.Lat()+p2.Lat())/2 {
		t.Error("LineString Lat failed")
	}
}

func TestLineString_Lon(t *testing.T) {
	if ls.Lon() != (p.Lon()+p2.Lon())/2 {
		t.Error("LineString Lon failed")
	}
}

func TestLineString_Z(t *testing.T) {
	if ls.Z() != (p.Z()+p2.Z())/2 {
		t.Error("LineString Z failed")
	}
}

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
	if ls.AsText() != fmt.Sprintf("LINESTRING (%f %f %f, %f %f %f)", p.Lat(), p.Lon(), p.Z(), p2.Lat(), p2.Lon(), p2.Z()) {
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

func TestLineString_Equals(t *testing.T) {
	if !ls.Equals(ls) {
		t.Error("LineString Equals failed")
	}

	if !ls.Equals(l.LineString) {
		t.Error("LineString Equals failed")
	}

	if ls.Equals(lr.LineString) {
		t.Error("LineString Equals failed")
	}
}

func TestLineString_Distance(t *testing.T) {
	p3 := NewPoint(2, 1, 1)
	p4 := NewPoint(3, 0, 0)

	ls2 := NewLineString(p3, p4)

	if ls.Distance(ls2) != 2.0 {
		t.Error("LineString Distance failed")
	}
}

func TestLineString_Length(t *testing.T) {
	correctLength := math.Sqrt(offset * 3)
	if ls.Length() != correctLength {
		t.Error("LineString Length failed")
	}

	ls2 := NewLineString(p, p2, p)
	correctLength = correctLength * 2
	if ls2.Length() != correctLength {
		t.Error("LineString IsClosed failed")
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
