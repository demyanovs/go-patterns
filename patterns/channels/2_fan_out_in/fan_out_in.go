package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
The idea of this pattern is to create a Goroutine for each individual piece of work that is pending and
can be done concurrently. In this code sample, I am going to create 2000 child Goroutines
to perform 2000 individual pieces of work. I am going to use a buffered channel since there is only one receiver,
and it’s not important to have a guarantee at the signaling level. That will only create extra latency.
*/
func main() {
	children := 2000
	ch := make(chan string, children)

	for c := 0; c < children; c++ {
		go func(child int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "Hello"
			fmt.Println("child sent signal: ", child)
		}(c)

		for children < 0 {
			msg := <-ch
			children--
			fmt.Println(msg)
			fmt.Println("parent received signal: ", msg)
		}
	}

	time.Sleep(time.Second)
	fmt.Println("--------------------------")
}
