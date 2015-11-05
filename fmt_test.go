package gocat

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestFormatted(t *testing.T) {
	assert.Equal(t, "3.14", fmt.Sprintf("%.2f", math.Pi))
}

func TestFormattedSlices(t *testing.T) {
	a := [2]float64{math.Pi, math.Pi}
	assert.Equal(t, "3.14,3.14", fmt.Sprintf("%.2f,%.2f", a[0], a[1]))
}
