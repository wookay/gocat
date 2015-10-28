package gocat

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {
	timestamp, err := strconv.ParseInt("1446375600", 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(timestamp, 0)
	assert.Equal(t, 2015, tm.Year())
	assert.Equal(t, "time.Month", reflect.TypeOf(tm.Month()).String())
	assert.Equal(t, time.November, tm.Month())
	assert.Equal(t, 11, int(tm.Month()))
	assert.Equal(t, 1, tm.Day())
	assert.Equal(t, "2015-11-01 20:00:00 +0900 KST", tm.String())
}
