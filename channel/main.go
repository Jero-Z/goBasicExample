package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string) // buffer zero

	go func() {
		//time.Sleep(3 * time.Second)
		tick := time.Tick(1 * time.Second)
		add := time.Now().Add(5 * time.Second)
		for i := range tick {
			fmt.Println(i)
			if i.After(add) {
				messages <- "ping"
				break
			}
		}
	}()

	msg := <-messages // this will block
	fmt.Println(msg)
}
