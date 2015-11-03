package gocat

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMake(t *testing.T) {
	var v = make([]int, 5)
	assert.Equal(t, 5, len(v))
}
