package main

import "fmt"

func main() {
	ch := make(chan int) // unbuffered channel

	// Sender goroutine
	//go func() {
	fmt.Println("Sending 42...")
	ch <- 42 // blocks until another goroutine receives
	fmt.Println("Sent 42")
	//}()

	// Receiver (in main goroutine)
	value := <-ch // blocks until a value is sent
	fmt.Println("Received:", value)
}
