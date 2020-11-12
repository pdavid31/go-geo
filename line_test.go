package geo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var l = NewLine(p, p2)

func TestLine_GeometryType(t *testing.T) {
	assert.Equal(t, "Line", l.GeometryType(), "Line GeometryType failed")
}
