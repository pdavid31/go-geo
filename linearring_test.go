package geo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var lr = NewLinearRing(p, p2)

func TestLinearRing_GeometryType(t *testing.T) {
	assert.Equal(t, "LinearRing", lr.GeometryType(), "LinearRing GeometryType failed")
}

func TestLinearRing_IsClosed(t *testing.T) {
	assert.True(t, lr.IsClosed(), "LinearRing IsClosed failed")
}
