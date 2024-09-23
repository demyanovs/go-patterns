package main

import (
	"fmt"
	"time"
)

/*
The drop pattern is an important pattern for services that may experience heavy loads at times and
can drop requests when the service reaches a capacity of pending requests. As an example, a DNS service
would need to employ this pattern.
*/
func main() {
	const capacity = 100
	ch := make(chan string, capacity)

	go func() {
		for msg := range ch {
			fmt.Println("child received signal:", msg)
		}
	}()

	const work = 2000
	for w := 0; w < work; w++ {
		select {
		case ch <- "Hello":
			fmt.Println("parent sent signal:", w)
		default:
			fmt.Println("parent dropped data:", w)
		}
	}

	close(ch)
	fmt.Println("parent sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("--------------------------")
}
