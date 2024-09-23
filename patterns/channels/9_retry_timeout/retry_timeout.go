package main

import (
	"context"
	"fmt"
	"time"
)

/*
*
The retry timeout pattern is great when I have to ping something (like a database)
which might fail, but I don’t want to fail immediately. I want to retry for a specified
amount of time before I fail.
*/
func retryTimeout(ctx context.Context, retryInterval time.Duration, check func(ctx context.Context) error) {
	for {
		fmt.Println("perform user check call")
		if err := check(ctx); err == nil {
			fmt.Println("work finished successfully")
			return
		}

		fmt.Println("check if timeout has expired")
		if ctx.Err() != nil {
			fmt.Println("time expired 1 :", ctx.Err())
			return
		}

		fmt.Printf("wait %s before trying again\n", retryInterval)
		t := time.NewTimer(retryInterval)

		select {
		case <-ctx.Done():
			fmt.Println("timed expired 2 :", ctx.Err())
			t.Stop()
			return
		case <-t.C:
			fmt.Println("retry again")
		}
	}
}
