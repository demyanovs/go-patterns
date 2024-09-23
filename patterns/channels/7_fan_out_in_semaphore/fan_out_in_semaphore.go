package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

/*
The fan out/in semaphore pattern provides a mechanic to control the number of Goroutines executing work
at any given time while still creating a unique Goroutine for each piece of work.
*/
func main() {
	children := 1000
	ch := make(chan string, children)

	g := runtime.GOMAXPROCS(0)
	sem := make(chan bool, g)

	for c := 0; c < children; c++ {
		go func(child int) {
			sem <- true
			{
				t := time.Duration(rand.Intn(10200)) * time.Millisecond
				time.Sleep(t)
				ch <- fmt.Sprintf("data: %d", child)
				fmt.Println("child sent signal:", child)
			}
			<-sem
		}(c)
	}

	for children > 0 {
		msg := <-ch
		children--
		fmt.Println("parent received signal:", msg)
	}

	time.Sleep(time.Second)
	fmt.Println("----------------------------")
}
