package main

import (
	"fmt"
	"sync"
)

type Mutex struct {
	s chan struct{}
}

func NewMutex() *Mutex {
	return &Mutex{
		s: make(chan struct{}, 1),
	}
}

func (m *Mutex) Lock() {
	m.s <- struct{}{}
}

func (m *Mutex) Unlock() {
	<-m.s
}

const numGoroutines = 1000

func main() {
	m := NewMutex()
	counter := 0
	var wg sync.WaitGroup

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			counter++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Mutex counter: %d\n", counter)
}
