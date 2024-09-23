package main

import "fmt"

func main() {
	ch := make(chan string, 3)

	go func() {
		ch <- "hello1"
		ch <- "hello2"
		ch <- "hello3"
		ch <- "hello4"
		ch <- "hello5"
	}()

	msg := <-ch
	msg2 := <-ch
	msg2 = <-ch
	msg2 = <-ch
	msg2 = <-ch
	//msg2 = <-ch

	fmt.Println(msg)
	fmt.Println(msg2)
	fmt.Println("----")
}
