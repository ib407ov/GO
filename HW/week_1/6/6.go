package main

import (
	"fmt"
	"sort"
)

//Функцию FilterAndSortMap, которая принимает карту (map)
//с ключами и значениями типа string и две функции: фильтр
//(функция, которая принимает string и возвращает bool)
//и компаратор (функция, которая сравнивает две строки и
//возвращает bool). Функция должна возвращать отсортированный
//срез ключей, значения которых проходят через фильтр.

//Привет, по 6 заданию тебе нужно дописать
//функцию FilterAndSortMap и main перечитай
//еще раз задание, что у тебя подается на вход
//и что тебе надо получить на выходе. У тебя есть
//пример в виде теста.

func FilterAndSortMap(m map[string]string, filterFunc func(string) bool, compareFunc func(string, string) bool) []string {
	var keys []string

	for k, val := range m {
		if filterFunc(val) {
			keys = append(keys, k)
		}
	}

	sort.Slice(keys, func(i, j int) bool {
		return compareFunc(keys[i], keys[j])
	})

	return keys
}

func filterFunc(s string) bool {
	if len(s) > 6 {
		return true
	}

	return false
}

func compareFunc(a, b string) bool {
	if a < b {
		return true
	}
	return false
}

func main() {
	m := map[string]string{
		"a": "apple",
		"b": "banana",
		"c": "cherry",
	}

	fmt.Println()
	fmt.Println(FilterAndSortMap(m, filterFunc, compareFunc))
}
