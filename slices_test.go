package gocat

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestMake(t *testing.T) {
	v := make([]int, 5)
	assert.Equal(t, 5, len(v))
}

func TestSplit(t *testing.T) {
	v := strings.Split("1,2", ",")
	assert.Equal(t, []string{"1", "2"}, v)
}

func TestToTwin(t *testing.T) {
	s := []string{"1", "2"}
	s2 := [...]string{s[0], s[1]}
	assert.Equal(t, [2]string{"1", "2"}, s2)
}

func TestTwinTo(t *testing.T) {
	s2 := [2]string{"1", "2"}
	s := []string{s2[0], s2[1]}
	assert.Equal(t, []string{"1", "2"}, s)
}
