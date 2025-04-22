package main

import (
	"fmt"
	"sync"
)

func split(source <-chan int, numWorkers int) []<-chan int {
	results := make([]<-chan int, 0)

	// Fan-out: start workers
	for i := 0; i < numWorkers; i++ {
		ch := make(chan int)
		results = append(results, ch)

		go func() {
			defer close(ch)

			for val := range source {
				ch <- val
			}
		}()
	}

	return results
}

func main() {
	const numWorkers = 5
	const numJobs = 10

	source := make(chan int)
	results := split(source, numWorkers)

	go func() {
		for i := 0; i < numJobs; i++ {
			source <- i
		}

		close(source)
	}()

	var wg sync.WaitGroup
	wg.Add(len(results))

	for i, ch := range results {
		go func(i int, d <-chan int) {
			defer wg.Done()

			for val := range d {
				fmt.Printf("Worker %d got value %d\n", i, val)
			}
		}(i, ch)
	}

	wg.Wait()
}
