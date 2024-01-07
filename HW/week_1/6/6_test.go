package main

import (
	"reflect"
	"testing"
)

func TestFilterAndSortMap(t *testing.T) {
	testMap := map[string]string{
		"a": "apple",
		"b": "banana",
		"c": "cherry",
	}

	filterFunc := func(s string) bool {
		return s > "banana"
	}

	compareFunc := func(i, j string) bool {
		return i < j
	}

	expected := []string{"c"}
	result := FilterAndSortMap(testMap, filterFunc, compareFunc)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FilterAndSortMap(%v, filterFunc, compareFunc) = %v; want %v", testMap, result, expected)
	}
}
