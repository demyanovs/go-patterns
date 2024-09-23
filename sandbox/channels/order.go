package main

import "fmt"

func main() {
	main3()
}

// A send on a channel happens before the corresponding receive from that channel completes.
// In the next example, a parent goroutine increments a variable before a send, while another goroutine reads it after a channel read
// The order is as follows:
// variable increment < channel send < channel receive < variable read
func main1() {
	i := 0
	ch := make(chan string)
	fmt.Println("START 1")

	go func() {
		fmt.Println("START 3")
		var x string
		x = <-ch
		fmt.Println("START 4", x, i)
		//fmt.Println(i)
	}()
	i++

	fmt.Println("START 2")
	ch <- "Hello"
	fmt.Println("START 5")
}

// Closing a channel happens before a receive of this closure.
// The next example is similar to the previous one, except that instead of sending a message, we close the channel:
func main2() {
	i := 0
	ch := make(chan struct{})
	fmt.Println("START 1")
	go func() {
		//fmt.Println("START 4")
		<-ch
		//fmt.Println("START 5", i)
	}()
	i++

	fmt.Println("START 2")
	close(ch)
	fmt.Println("START 3")
}

// a receive from an unbuffered channel happens before the send on that channel completes.
// The write is guaranteed to happen before the read.
func main3() {
	i := 0
	ch := make(chan struct{}, 1)
	fmt.Println("START 1")
	go func() {
		i = 1
		//fmt.Println("START 4")
		<-ch
		fmt.Println("START 5", i)
	}()
	fmt.Println("START 2")
	ch <- struct{}{}
	fmt.Println("START 3", i)
}
