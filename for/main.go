package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i := 1; i <= 3; { // 同名变量在循环中变量的作用域仅在for中
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println(i)

	fmt.Println()

	for j := 8; j < 10; j++ {
		fmt.Println(j)
	}

	fmt.Println()
	for {
		fmt.Println("loop") // print once
		break
	}
	fmt.Println()
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n) // print 1 3 5
	}

	for {
		fmt.Println("this is long loop; but 4 second exit")

		time.Sleep(time.Second * 4)
		break
	}
}
