package gocat

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"strings"
	"testing"
)

func TestFormatted(t *testing.T) {
	assert.Equal(t, "3.14", fmt.Sprintf("%.2f", math.Pi))
}

func TestFormattedSlices(t *testing.T) {
	a := [2]float64{math.Pi, math.Pi}
	assert.Equal(t, "3.14,3.14", fmt.Sprintf("%.2f,%.2f", a[0], a[1]))
}

type Map interface {
	fmap(f func(interface{}) string) []string
}

type Point [2]float64

func (point Point) fmap(f func(interface{}) string) []string {
	arr := make([]string, len(point))
	for idx, value := range point {
		arr[idx] = f(value)
	}
	return arr
}

func TestPoint(t *testing.T) {
	point := [2]float64{math.Pi, math.Pi}

	sprint_float := func(value interface{}) string {
		return fmt.Sprintf("%.2f", value)
	}
	a := Point(point).fmap(sprint_float)
	assert.Equal(t, []string{"3.14", "3.14"}, a)
	assert.Equal(t, "3.14,3.14", strings.Join(a, ","))
}
