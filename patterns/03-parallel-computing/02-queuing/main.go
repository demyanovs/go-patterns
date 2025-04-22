package main

import (
	"fmt"
	"sync"
	"time"
)

func process(payload int, queue chan struct{}, wg *sync.WaitGroup) {
	queue <- struct{}{}

	go func() {
		defer wg.Done()

		fmt.Printf("Start processing of %d\n", payload)
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("Completed processing of %d\n", payload)
		fmt.Printf("Queue length: %d\n\n", len(queue))

		<-queue
	}()
}

func main() {
	const numWorkers = 3
	const numMessages = 1000

	var wg sync.WaitGroup

	fmt.Println("Queue of length numWorkers:", numWorkers)

	// Buffered channel as semaphore
	queue := make(chan struct{}, numWorkers)

	wg.Add(numMessages)

	for w := 1; w <= numMessages; w++ {
		process(w, queue, &wg)
	}

	wg.Wait()

	close(queue)
	fmt.Println("Processing completed")
}

/*
Очередь
Позволяет принимать на обработку до N сообщений одновременно не дожидаясь их обработки.
*/
