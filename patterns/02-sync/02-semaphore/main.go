package main

/**
Еще один паттерн для организации rate-limiter.

Помогает ограничить кол-во параллельно исполняемых го-рутин.
Мы создаем семафор с определенным значением параллелизма, которое ограничивает сверху
кол-во одновременно исполняемых НЕ заблокируемых го-рутин.
*/

import (
	"fmt"
	"time"
)

type semaphore chan struct{}

func NewSemaphore(n int) semaphore {
	return make(semaphore, n)
}

func (s semaphore) Lock(n int) {
	e := struct{}{}
	for i := 0; i < n; i++ {
		s <- e
	}
}

func (s semaphore) Unlock(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

const N = 3
const TOTAL = 10

func main() {
	sem := NewSemaphore(N)
	done := make(chan bool)
	for i := 1; i <= TOTAL; i++ {
		sem.Lock(1)
		go func(v int) {
			defer sem.Unlock(1)
			process3(v)
			if v == TOTAL {
				done <- true
			}
		}(i)
	}
	<-done
}
func process3(id int) {
	fmt.Printf("[%s]: running task %d\n", time.Now().Format("15:04:05"), id)
	time.Sleep(time.Second)
}
