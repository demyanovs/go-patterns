package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int)

	out := merge(ch1, ch2)

	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()

	go func() {
		ch2 <- 3
		close(ch2)
	}()

	for val := range out {
		fmt.Println(val)
	}
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(cs))
		for _, c := range cs {
			go func() {
				defer wg.Done()
				for val := range c {
					out <- val
				}
			}()
		}

		wg.Wait()
		close(out)
	}()

	return out
}
