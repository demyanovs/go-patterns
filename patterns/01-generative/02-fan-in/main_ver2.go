package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Fan-in multiplexes multiple input channels into a single output channel.

Applicability
Services that have some number of workers that all generate output may find it useful to combine
all of the workers' outputs to be processed as a single unified stream.

Participants
Sources
A set of one or more input channels with the same type. Accepted by Funnel.

Destination
An output channel of the sa,e  type as Sources. Created and provided by Funnel.

Funnel
Accepts Sources and immediately returns Destination. Any input from any Sources will be output by Destination.
*/

func main() {
	sources := make([]<-chan int, 0)

	for i := 0; i < 3; i++ {
		ch := make(chan int)
		sources = append(sources, ch)

		go func() {
			defer close(ch)
			time.Sleep(time.Second)
			ch <- i
		}()
	}

	dest := Funnel(sources...)
	for d := range dest {
		fmt.Println(d)
	}
}

func Funnel(sources ...<-chan int) <-chan int {
	dest := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(sources))

	for _, ch := range sources {
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				dest <- n
			}
		}(ch)
	}

	go func() {
		wg.Wait() // start a goroutine to close dest once all sources have closed
		close(dest)
	}()

	return dest
}
