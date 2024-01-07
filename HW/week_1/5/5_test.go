package main

import (
	"reflect"
	"testing"
)

func TestMergeAndSort(t *testing.T) {
	testCases := []struct {
		a, b, expected []int
	}{
		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{}, []int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{[]int{}, []int{}, []int{}},
	}

	for _, tc := range testCases {
		result := MergeAndSort(tc.a, tc.b)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("MergeAndSort(%v, %v) = %v; want %v", tc.a, tc.b, result, tc.expected)
		}
	}
}
