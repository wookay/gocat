package gocat

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func deg_to_rad(deg float64) float64 {
	return (deg * math.Pi / 180)
}

func rad_to_deg(rad float64) float64 {
	return (rad * 180 / math.Pi)
}

func TestRadian(t *testing.T) {
	assert.Equal(t, math.Pi, deg_to_rad(180))
	assert.Equal(t, 180.0, rad_to_deg(math.Pi))
}
