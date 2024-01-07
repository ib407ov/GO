package main

import "fmt"

func ffunction(list []string) map[string]int {

	m := make(map[string]int)

	for i := 0; i < len(list); i++ {
		if _, ok := m[list[i]]; !ok {
			m[list[i]] = 1
			for j := i + 1; j < len(list); j++ {
				if list[i] == list[j] {
					m[list[i]]++
				}
			}
		}
	}

	return m
}

func main() {
	listEmployee := []string{
		"BTC",
		"Hello",
		"Hello",
		"Hello",
		"Hello",
		"ETH",
		"BTC",
		"BTC",
		"ETH",
	}

	fmt.Println(ffunction(listEmployee))

}
