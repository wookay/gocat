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
