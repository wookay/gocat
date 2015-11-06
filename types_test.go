package gocat

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGocat(t *testing.T) {
	gocat()
	assert.Equal(t, nil, nil)
}

func TestNil(t *testing.T) {
	assert.Equal(t, nil, reflect.TypeOf(nil))
}

func TestInt(t *testing.T) {
	assert.Equal(t, 3, 1+2)
	assert.Equal(t, "int", reflect.TypeOf(2).String())
}

func TestFloat(t *testing.T) {
	assert.Equal(t, 3.14, 3.14)
	assert.Equal(t, "float64", reflect.TypeOf(3.14).String())
	assert.Equal(t, true, 0.1 > 0)
}

func TestString(t *testing.T) {
	assert.Equal(t, "abc", "abc")
	assert.Equal(t, "string", reflect.TypeOf("abc").String())
}

func TestArray(t *testing.T) {
	var a = []interface{}{
		"zero",
		1,
		3.14,
	}
	assert.Equal(t, "zero", a[0])
	assert.Equal(t, 1, a[1])
	assert.Equal(t, 3.14, a[2])
	assert.Equal(t, "[]interface {}", reflect.TypeOf(a).String())
}

func TestMultiDimensionalArray(t *testing.T) {
	var a = [][]string{
		[]string{"a", "b", "c"},
	}
	assert.Equal(t, "a", a[0][0])
	assert.Equal(t, "b", a[0][1])
	assert.Equal(t, "c", a[0][2])
}

func TestStruct(t *testing.T) {
	type ABC struct {
		v interface{}
	}
	var a = ABC{}
	var b = ABC{"A"}
	var c = ABC{1}
	assert.Equal(t, nil, a.v)
	assert.Equal(t, "A", b.v)
	assert.Equal(t, 1, c.v)
	assert.Equal(t, "gocat.ABC", reflect.TypeOf(a).String())
	assert.Equal(t, "gocat.ABC", reflect.TypeOf(b).String())
	assert.Equal(t, "gocat.ABC", reflect.TypeOf(c).String())
}

func TestFuncOf(t *testing.T) {
	var k = reflect.TypeOf(0)
	var e = reflect.TypeOf("")
	assert.Equal(t, "func(int) string", reflect.FuncOf([]reflect.Type{k}, []reflect.Type{e}, false).String())
}
