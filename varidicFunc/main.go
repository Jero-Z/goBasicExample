package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	count := 0
	for _, num := range nums {
		total += num
		count++
	}
	fmt.Printf("total:%d,count:%d\n", total, count)
}

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
