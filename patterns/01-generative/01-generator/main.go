package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Channel used to signal the generator to stop, with a message
	stop := make(chan string)

	// WaitGroup to ensure all goroutines finish before exiting
	wg := sync.WaitGroup{}
	wg.Add(2) // One for generator, one for consumer

	// Start the generator
	ch := makeGenerator(stop, &wg)

	// Consumer goroutine
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("value:", v)
		}
	}()

	// Let the generator run for a bit
	time.Sleep(time.Second * 1)

	// Send a stop signal with a message, then close the stop channel
	stop <- "finish job"
	close(stop)

	// Wait for both goroutines to finish
	wg.Wait()
}

// makeGenerator starts a goroutine that generates an increasing sequence of integers.
// It writes them to a buffered channel, and stops when it receives a message on the stop channel.
func makeGenerator(stop <-chan string, wg *sync.WaitGroup) <-chan int {
	ch := make(chan int, 1) // Buffered channel with capacity 1
	i := 0

	go func() {
		defer wg.Done()
		for {
			select {
			case msg := <-stop:
				// Stop signal received, print message and close the channel
				fmt.Printf("done, got message: %s\n", msg)
				close(ch)
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

//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//func main() {
//	done := make(chan string)
//	wg := sync.WaitGroup{}
//	wg.Add(2)
//
//	ch := makeGenerator(done, &wg)
//
//	go func() {
//		defer wg.Done()
//		for v := range ch {
//			fmt.Println("value:", v)
//		}
//	}()
//
//	time.Sleep(time.Second * 1)
//
//	done <- "finish job"
//	close(done)
//
//	wg.Wait()
//}
//
//func makeGenerator(done <-chan string, wg *sync.WaitGroup) <-chan int {
//	ch := make(chan int, 1)
//	var i = 0
//
//	go func() {
//		defer wg.Done()
//		for {
//			select {
//			case x := <-done:
//				close(ch)
//				fmt.Printf("done, got message: %s\n", x)
//				return
//			default:
//				time.Sleep(time.Millisecond * 250)
//				ch <- i
//				i++
//			}
//		}
//	}()
//
//	return ch
//}

/**
(input) -> (generator) -> (channel) -> (recipient)
Полезен, когда надо читать из очереди сообщений и обрабатывать в отдельных го-рутинах, не блокируя чтение из очереди.
Генератор будут заниматься только чтением этой очереди в буферизированный канал. Таким образом, запись в канал
не будет блокироваться, пока в буфере есть место под новые сообщения (в нашем примере это 1).
*/
