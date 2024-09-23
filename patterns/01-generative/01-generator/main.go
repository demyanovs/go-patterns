package main

import (
	"fmt"
	"sync"
	"time"
)

/**
(input) -> (generator) -> (channel) -> (recipient)
Полезен, когда надо читать из очереди сообщений и обрабатывать в отдельных го-рутинах, не блокируя чтение из очереди.
Генератор будут заниматься только чтением этой очереди в буферизированный канал. Таким образом, запись в канал
не будет блокироваться, пока в буфере есть место под новые сообщения (в нашем примере это 1).
*/

func main() {
	done := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(2)

	ch := makeGenerator(done, &wg)

	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("value:", v)
		}
	}()

	time.Sleep(time.Second * 1)

	done <- "finish job"
	//close(done)
	wg.Wait()
}

func makeGenerator(done <-chan string, wg *sync.WaitGroup) <-chan int {
	ch := make(chan int, 1)
	var i = 0

	go func() {
		defer wg.Done()
		for {
			select {
			case x := <-done:
				close(ch)
				fmt.Printf("done, got message: %s\n", x)
				return
			default:
				time.Sleep(time.Millisecond * 250)
				ch <- i
				i++
			}
		}
	}()

	return ch
}
