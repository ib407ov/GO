package main

import "fmt"

func multiply(value *int) {
	*value *= *value
}

func main() {
	val := 4
	//fmt.Println(multiply(&val))
	multiply(&val)
	fmt.Println(val)
}
