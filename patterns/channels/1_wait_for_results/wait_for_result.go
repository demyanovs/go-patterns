package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
The wait for result pattern is a foundational pattern used by larger patterns like fan out/in.
In this pattern, a Goroutine is created to perform some known work and signals their result back to
the Goroutine that created them. This allows for the actual work to be placed on a Goroutine
that can be terminated or walked away from.
*/
func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "Hello"
		fmt.Println("child: sent signal")
	}()

	fmt.Println("no wait")

	msg := <-ch
	fmt.Println("parent: received signal:", msg)

	time.Sleep(time.Second)
	fmt.Println("------------------------")
}
