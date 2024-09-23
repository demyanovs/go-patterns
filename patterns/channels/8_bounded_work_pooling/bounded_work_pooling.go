package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
The bounded work pooling pattern uses a pool of Goroutines to perform a fixed amount of known work.
*/
func main() {
	work := []string{"paper", "paper", "paper", "paper", 2000: "paper"}
	g := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(g)
	ch := make(chan string, g)

	for c := 0; c < g; c++ {
		go func(child int) {
			defer wg.Done()
			for wrk := range ch {
				fmt.Printf("child %d : recv'd signal : %s\n", child, wrk)
			}
			fmt.Printf("child %d : recv'd shutdown signal\n", child)
		}(c)
	}

	for _, wrk := range work {
		ch <- wrk
	}

	close(ch)
	wg.Wait()

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
