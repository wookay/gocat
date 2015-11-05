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
