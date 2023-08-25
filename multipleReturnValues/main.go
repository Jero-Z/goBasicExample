package main

import "fmt"

func getMultiValue() (int, int) {
	return 3, 7
}

func main() {

	a, b := getMultiValue()
	fmt.Println(a, b)

	_, c := getMultiValue()
	fmt.Println(c)
}
