package gocat

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	//"log"
	"math"
	"strconv"
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
	smap(f func(interface{}) string) [2]string
	fmap(f func(interface{}) float64) [2]float64
}

type Point [2]float64

func (point Point) smap(f func(interface{}) string) []string {
	return []string{f(point[0]), f(point[1])}
}

func TestPoint(t *testing.T) {
	floats := [2]float64{math.Pi, math.Pi}

	sprintFloat := func(value interface{}) string {
		return fmt.Sprintf("%.2f", value)
	}
	a := Point(floats).smap(sprintFloat)
	assert.Equal(t, []string{"3.14", "3.14"}, a)
	assert.Equal(t, "3.14,3.14", strings.Join(a, ","))
}

type PointString []string

func (point PointString) fmap(f func(string) float64) [2]float64 {
	return [2]float64{f(point[0]), f(point[1])}
}

func TestParse(t *testing.T) {
	s := "3.14,3.14"
	points := strings.Split(s, ",")
	parseFloat := func(value string) float64 {
		f, _ := strconv.ParseFloat(value, 64)
		return f
	}
	floats := PointString(points).fmap(parseFloat)
	assert.Equal(t, [2]float64{3.14, 3.14}, floats)
}
