package main

import (
	"fmt"
)

func FilterAndSortMap(m map[string]string, filterFunc func(string) bool, compareFunc func(string, string) bool) []string {
	var keys []string

	for k := range m {
		if filterFunc(k) {
			keys = append(keys, k)
		}
	}

	//sort.Slice(keys, func(i, j int) bool {
	//	return compareFunc(keys[i], keys[j])
	//})

	return keys
}

func filterFunc(s string) bool {
	return true
}

func compareFunc(a, b string) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]string{
		"a": "apple",
		"b": "banana",
		"z": "cherry",
	}

	fmt.Println(FilterAndSortMap(m, filterFunc, compareFunc))
}
