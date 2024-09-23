package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

type Cache struct {
	rate int64
}

func main() {
	cache := NewCache(0)

	ctx := context.Background()
	done := make(chan struct{})
	defer close(done)

	tick := time.NewTicker(time.Second * 1)
	defer tick.Stop()

	go func() {
		for {
			select {
			case <-tick.C:
				cache.updateRate(ctx)
			case <-done:
				tick.Stop()
				fmt.Println("DONE SIGNAL")
				return
			case <-ctx.Done():

				return
			}
		}
	}()

	time.Sleep(time.Second * 4)
	done <- struct{}{}

	time.Sleep(time.Second * 4)
	fmt.Println("EXIT")
}

func NewCache(rate int64) *Cache {
	return &Cache{rate: rate}
}

func (c *Cache) GetRate() int64 {
	return c.rate
}

func (c *Cache) updateRate(ctx context.Context) {
	go func() {
		// TODO to http request
		fmt.Println("Cache", c.GetRate())
		atomic.StoreInt64(&c.rate, 56)
		fmt.Printf("updating cache at: %s. Val: %d\n", time.Now().Format("2006-01-02 15:04:05"), c.rate)
	}()
}
