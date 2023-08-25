package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...\n")
	time.Sleep(time.Second * 4) // 4 second
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	<-done
}
