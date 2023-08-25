package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Printf("%T,%T", nextInt, newInt)
	fmt.Println()
	fmt.Printf("%v,%v", &nextInt, &newInt)
	fmt.Println(newInt())
}
