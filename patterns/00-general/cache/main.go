package main

import (
	"context"
	"fmt"
	"math/rand"
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

	fmt.Println("Exit")
}

func NewCache(rate int64) *Cache {
	return &Cache{rate: rate}
}

func (c *Cache) GetRate() int64 {
	return c.rate
}

func (c *Cache) updateRate(ctx context.Context) {
	go func() {
		fmt.Println("current value:", c.GetRate())
		atomic.StoreInt64(&c.rate, int64(rand.Intn(100)))
		fmt.Printf("updating cache at: %s. val: %d\n", time.Now().Format("2006-01-02 15:04:05"), c.rate)
	}()
}
