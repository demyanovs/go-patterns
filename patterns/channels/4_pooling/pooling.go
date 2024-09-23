package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
The pooling pattern uses the wait for task pattern just described. The pooling pattern allows me to manage
resource usage across a well-defined number of Goroutines. As explained previously, in Go pooling is not
needed for efficiency in CPU processing like at the operating system. It’s more important for efficiency
in resource usage.
*/
func main() {
	ch := make(chan string)

	g := runtime.GOMAXPROCS(0)
	for c := 0; c < g; c++ {
		go func(child int) {
			for msg := range ch {
				fmt.Printf("child %d received signal %s\n", child, msg)
			}
			fmt.Printf("child %d received shutdown signal\n", child)
		}(c)
	}

	const work = 10
	for w := 0; w < work; w++ {
		ch <- "Hello"
		fmt.Println("parent sent signal: ", w)
	}

	close(ch)
	fmt.Println("parent : sent shutdown signal")
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
