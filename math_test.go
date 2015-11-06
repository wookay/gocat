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

func TestSin(t *testing.T) {
	assert.Equal(t, 1.0, math.Sin(deg_to_rad(90)))
}

func TestPow(t *testing.T) {
	assert.Equal(t, 1, 2^3)
	assert.Equal(t, 8.0, math.Pow(2, 3))
}

func TestAbs(t *testing.T) {
	assert.Equal(t, 1.0, math.Abs(-1))
	assert.Equal(t, 1, int(math.Abs(-1)))
}
