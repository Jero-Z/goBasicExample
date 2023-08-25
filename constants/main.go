package main

import (
	"fmt"
	"math"
)

const name string = "jack"

func main() {
	fmt.Println(name)

	//name ="ss" :Cannot assign to name
	const n = 500000000
	const d = 3e20 / n // 6e+10^11

	fmt.Println(d)
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
