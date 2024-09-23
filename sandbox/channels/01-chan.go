package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		fmt.Println("inside goroutine 1")
		time.Sleep(time.Second * 3)
		fmt.Println("inside goroutine 2")

		messages <- "from go-routing"
	}()

	fmt.Println("main message")

	fmt.Println(<-messages)

	close(messages)
}
