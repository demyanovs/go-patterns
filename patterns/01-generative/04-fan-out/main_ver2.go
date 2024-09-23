package main

import (
	"fmt"
	"sync"
)

/**
Fan-out evenly distributes messages from an input channel to multiple output channels.

Applicability
Fan-out revives messages from an input channel, distributing them evenly among output channels, and is a useful pattern for parallelizing CPU and I/O utilization.

For example, imaging that you have an input source, such as Reader on input stream, or a listener on a message broker,
that provides the inputs for some resource-intensive unit of work.
Rather that coupling the input and computation processes, which would confine the effort to a single serial process,
you might prefer to parallelize the workload by distributing it among some number of concurrent worker processes

Participants
Source
An input channel. Accepted by Split.
Destinations
An output channel of the sa,e type as SSource. Created and provided by Split.
Split
A function that accepts Source and immediately returns Destinations. Any input from Source will be output to Destination.
*/

func main() {
	source := make(chan int)
	dests := Split(source, 5)

	go func() {
		for i := 0; i < 10; i++ {
			source <- i
		}

		close(source)
	}()

	var wg sync.WaitGroup
	wg.Add(len(dests))

	for i, ch := range dests {
		go func(i int, d <-chan int) {
			defer wg.Done()

			for val := range d {
				fmt.Printf("Worker %d got value %d\n", i, val)
			}
		}(i, ch)
	}

	wg.Wait()
}

func Split(source <-chan int, n int) []<-chan int {
	dests := make([]<-chan int, 0)

	for i := 0; i < n; i++ {
		ch := make(chan int)
		dests = append(dests, ch)

		go func() {
			defer close(ch)

			for val := range source {
				ch <- val
			}
		}()
	}

	return dests
}
