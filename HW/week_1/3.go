package main

import "fmt"

type Employee struct {
	Name       string
	Salary     int
	Department string
}

func function(list []Employee) map[string][]Employee {

	m := make(map[string][]Employee)

	for i := 0; i < len(list); i++ {
		if _, ok := m[list[i].Department]; !ok {
			m[list[i].Department] = []Employee{list[i]}
		}
		for j := i + 1; j < len(list); j++ {
			if list[i].Department == list[j].Department {
				m[list[i].Department] = append(m[list[i].Department], list[j])
			}
		}
	}

	return m
}

func main() {
	listEmployee := []Employee{
		{
			Name:       "Liza",
			Salary:     2000,
			Department: "manager",
		},
		{
			Name:       "Aekadij",
			Salary:     3000,
			Department: "owner",
		},
		{
			Name:       "Kasper",
			Salary:     4000,
			Department: "developer",
		},
		{
			Name:       "Emil",
			Salary:     5000,
			Department: "manager",
		},
		{
			Name:       "Anton",
			Salary:     6000,
			Department: "developer",
		},
		{
			Name:       "Rat",
			Salary:     7000,
			Department: "developer",
		},
	}

	fmt.Println(function(listEmployee))
}
