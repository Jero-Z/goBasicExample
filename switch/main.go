package main

import (
	"fmt"
	"time"
)

func main() {
	//scanInput()

	//multiCase()
	//comparableCase()
	whatAmI := func(i any) { // default i interface{} ;golang 1.18 later can use any support
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

func comparableCase() {
	t := time.Now()
	switch {
	case t.Hour() < 14:
		fmt.Println("It's before noon")
	case t.Hour() < 18:
		fmt.Println("but to late")
	default:
		fmt.Println("It's after noon")
	}
}

func multiCase() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
}

func scanInput() {
	var i int
	fmt.Println("please input some number")
	_, err := fmt.Scan(&i)
	if err != nil {
		return
	}
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("nothing")
	}
}
