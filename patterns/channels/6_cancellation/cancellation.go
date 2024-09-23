package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/*
The cancellation pattern is used to tell a function performing some I/O how long I am willing to wait for
the operation to complete. Sometimes I can cancel the operation, and sometimes all I can do is just walk away.
*/
func main() {
	duration := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	}()

	select {
	case d := <-ch:
		fmt.Println("work complete:", d)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}

	time.Sleep(time.Second)
	fmt.Println("------------------------")
}
