package main

import (
	"fmt"
	"sync"
)

/**
Разветвитель.

										 ------->(out channel)
										|
(in channel) ------>(fan out)  ------>	------->(out channel)
										|
										 ------->(out channel)

Позволяет разделить один входной канал на несколько выходных. Это полезный паттерн для распределения обработки между
несколькими однородными го-рутинами.

Полезен для распределения нагрузки между каналами.
*/

func main() {
	wg := sync.WaitGroup{}
	work := []int{1, 2, 3, 4, 5, 6, 7, 8}

	wg.Add(1)
	in := generateWork(work, &wg)

	wg.Add(3)
	fanOut2(in, "Alice", &wg)
	fanOut2(in, "Jack", &wg)
	fanOut2(in, "Bob", &wg)

	wg.Wait()
}

func fanOut2(in <-chan int, name string, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()

		for data := range in {
			fmt.Println(name, "processed", data)
		}
	}()
}

func generateWork(work []int, wg *sync.WaitGroup) <-chan int {
	ch := make(chan int)

	go func() {
		defer wg.Done()
		defer close(ch)

		for _, w := range work {
			ch <- w
		}
		fmt.Println("All data written")
	}()

	return ch
}
