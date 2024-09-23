package main

import "fmt"

func main() {
	messageCh := make(chan int, 5)
	disconnectCh := make(chan struct{}, 1)

	for i := 0; i < 10; i++ {
		messageCh <- i
	}
	disconnectCh <- struct{}{}

	for {
		select {
		case v := <-messageCh:
			fmt.Println(v)
		case <-disconnectCh:
			fmt.Println("disconnection, return")
			return
		}
	}

}
