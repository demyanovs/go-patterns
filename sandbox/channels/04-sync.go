package main

import (
	"fmt"
)

func main() {
	i := 0
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	go func() {
		ch <- 1
	}()

	i += <-ch
	i += <-ch

	fmt.Println(i)
}
