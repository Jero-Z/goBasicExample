package main

import "fmt"

func main() {
	var a [4]int
	fmt.Println(a) // by default in array is zero value for int this is 0

	var b [4]string
	fmt.Println(b)
	var c [4]byte
	fmt.Println(c) // byte print auto trans int
	var d [4]complex64
	fmt.Println(d)
	a[3] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[3])
	fmt.Println("length:", len(a))

	e := [5]int{1, 2, 3, 4, 5} // init and set value
	fmt.Println("dcl:", e)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
