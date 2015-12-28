package gocat

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBackquote(t *testing.T) {
	assert.Equal(t, "\n1\n2\n3\n", `
1
2
3
`)
}
