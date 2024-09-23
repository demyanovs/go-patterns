package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
The wait for task pattern is a foundational pattern used by larger patterns like pooling.

At the beginning, the function creates an unbuffered channel so there is a guarantee at the signaling level.
This is critically important for pooling so I can add mechanics later if needed to allow for timeouts and
cancellation. Once the channel is created, a child Goroutine is created immediately waiting for a signal
with data to perform work. The parent Goroutine begins to prepare that work and finally signals the work
to the child Goroutine. Since the guarantee is at the signaling level, the child Goroutine doesn’t
know how long it needs to wait.
*/
func main() {
	ch := make(chan string)

	go func() {
		msg := <-ch
		fmt.Println("child received signal: ", msg)
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- "Hello"
	fmt.Println("parent sent signal")

	time.Sleep(time.Second)
	fmt.Println("--------------------")
}
