package main

import (
	"fmt"
)

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	for key, value := range n2 {
		fmt.Println(key, value)
	}

	n3 := map[string]func(){
		"action": func() {
			fmt.Println("this is action func")
		},
		"step_one": func() {
			fmt.Println("this is step one func")
		},
	}
	for _, value := range n3 {
		value()
	}
}
