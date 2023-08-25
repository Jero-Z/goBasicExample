package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received section one")
		fmt.Println("this won't block ")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("section two input msg to  message", msg)
	default:
		fmt.Println("no message sent")

	}
	time.Sleep(2 * time.Second)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
