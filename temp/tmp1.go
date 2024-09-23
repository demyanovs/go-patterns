package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	var wg sync.WaitGroup
	wg.Add(len(m))

	for k, _ := range m {
		k := k
		go func() {
			defer wg.Done()

			//time.Sleep(time.Millisecond * 100)
			m[k]++
		}()
	}

	wg.Wait()

	fmt.Println(m)
}
