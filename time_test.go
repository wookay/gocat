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
	timezone, offset := tm.Zone()
	if "KST" == timezone {
		assert.Equal(t, 9*60*60, offset)
		assert.Equal(t, "2015-11-01 20:00:00 +0900 KST", tm.String())
	} else {
		assert.Equal(t, "UTC", timezone)
		assert.Equal(t, 0, offset)
		assert.Equal(t, "2015-11-01 11:00:00 +0000 UTC", tm.String())
	}
}
