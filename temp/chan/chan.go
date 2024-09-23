package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		ch <- "hi"
	}()

	//msg := <-ch

	go func() {
		//time.Sleep(1 * time.Second)
		ch <- "hi2"
		//close(ch)
	}()

	//go func() {
	//	//time.Sleep(1 * time.Second)
	//	ch <- "hi3"
	//}()

	//fmt.Println(msg)

	for msg := range ch {
		fmt.Println(msg)
	}
}
