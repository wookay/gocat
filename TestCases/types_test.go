package main

import (
	"../listdict"
	"fmt"
	"reflect"
	"testing"
)

func assert_equal(expected, got interface{}) {
	if expected != got {
		fmt.Printf("Expected: %v\n     Got: %v\n", expected, got)
		panic("not equal")
	}
}

func TestInt(t *testing.T) {
	assert_equal(3, 1+2)
	assert_equal("int", reflect.TypeOf(2).String())
}

func TestFloat(t *testing.T) {
	assert_equal(3.14, 3.14)
	assert_equal("float64", reflect.TypeOf(3.14).String())
}

func TestString(t *testing.T) {
	assert_equal("abc", "abc")
	assert_equal("string", reflect.TypeOf("abc").String())
}

func TestArray(t *testing.T) {
	var a = []interface{}{
		"zero",
		1,
		3.14,
	}
	assert_equal("zero", a[0])
	assert_equal(1, a[1])
	assert_equal(3.14, a[2])
	assert_equal("[]interface {}", reflect.TypeOf(a).String())
}

func TestDict(t *testing.T) {
	var d = listdict.Dict{"one": 1, "two": 2}
	assert_equal(1, d["one"])
	assert_equal(2, d["two"])
	assert_equal("listdict.Dict", reflect.TypeOf(d).String())
}

func TestStruct(t *testing.T) {
	type ABC struct {
		v interface{}
	}
	var a = ABC{}
	var b = ABC{"A"}
	var c = ABC{1}
	assert_equal(nil, a.v)
	assert_equal("A", b.v)
	assert_equal(1, c.v)
	assert_equal("main.ABC", reflect.TypeOf(a).String())
	assert_equal("main.ABC", reflect.TypeOf(b).String())
	assert_equal("main.ABC", reflect.TypeOf(c).String())
}
