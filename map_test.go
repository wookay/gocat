package gocat

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func keys(m map[string]bool) []string {
	a := []string{}
	for key, _ := range m {
		a = append(a, key)
	}
	return a
}

func TestMap(t *testing.T) {
	m := map[string]bool{}
	assert.Equal(t, map[string]bool{}, m)
	m["a"] = true
	assert.Equal(t, map[string]bool{"a": true}, m)
	m["a"] = true
	assert.Equal(t, map[string]bool{"a": true}, m)
	assert.Equal(t, []string{"a"}, keys(m))
	assert.Equal(t, false, m["b"])
}

type Any bool
type AnyMap map[string]*Any

func (m AnyMap) keys() []string {
	a := []string{}
	for key, _ := range m {
		a = append(a, key)
	}
	return a
}

func TestAnyMap(t *testing.T) {
	m := AnyMap{}
	assert.Equal(t, AnyMap{}, m)
	a := Any(true)
	m["a"] = &a
	assert.Equal(t, AnyMap{"a": &a}, m)
	m["a"] = &a
	assert.Equal(t, AnyMap{"a": &a}, m)
	assert.Equal(t, []string{"a"}, m.keys())
	assert.Equal(t, "*gocat.Any", reflect.TypeOf(m["b"]).String())
	bisnil := m["b"] == nil
	assert.Equal(t, true, bisnil)
}
