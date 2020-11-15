package geo

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mls = NewMultiLineString(ls)

func TestMultiLineString_GeometryType(t *testing.T) {
	assert.Equal(t, "MultiLineString", mls.GeometryType(), "MultiLineString GeometryType failed")
}

func TestMultiLineString_SRID(t *testing.T) {
	assert.Equal(t, 4326, mls.SRID(), "MultiLineString SRID failed")
}

func TestMultiLineString_AsText(t *testing.T) {
	asText := fmt.Sprintf("MULTILINESTRING ((%f %f %f, %f %f %f))", p.Lat(), p.Lon(), p.Z(), p2.Lat(), p2.Lon(), p2.Z())
	assert.Equal(t, asText, mls.AsText())
}

func TestMultiLineString_IsEmpty(t *testing.T) {
	assert.False(t, mls.IsEmpty(), "MultiLineString IsEmpty failed")
}

func TestMultiLineString_Is3D(t *testing.T) {
	assert.True(t, mls.Is3D(), "MultiLineString Is3D failed")
}

func TestMultiLineString_NumGeometries(t *testing.T) {
	assert.Equal(t, 1, mls.NumGeometries(), "MultiLineString NumGeometries failed")
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
	assert.Equal(t, correctLength, mls.Length(), "MultiLineString Length failed")
}

func TestMultiLineString_IsClosed(t *testing.T) {
	assert.False(t, mls.IsClosed(), "MultiLineString IsClosed failed")
}
