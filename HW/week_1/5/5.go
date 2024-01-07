package main

import (
	"fmt"
	"sort"
)

// MergeAndSort объединяет и сортирует два среза
func MergeAndSort(a, b []int) []int {
	var arr []int

	arr = append(a, b...)
	sort.Ints(arr)

	return arr
}

func main() {
	a := []int{1, 2, 2, 2, 6, 6, 6}
	b := []int{0, 1, 3, 3, 3, 4, 4, 8, 8}
	fmt.Println(MergeAndSort(a, b))

}
