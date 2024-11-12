package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

var errFailure = errors.New("some error")

func main() {
	ctx := context.Background()
	err := FetchUserData(ctx)
	if err != nil {
		fmt.Println("Error fetching user data:", err)
	}
}

func FetchUserData(ctx context.Context) error {
	//group := errgroup.Group{}
	group, qctx := errgroup.WithContext(context.Background())

	// Run first periodic task.
	group.Go(func() error {
		firstTask(qctx)
		return nil
	})

	// Run second task.
	group.Go(func() error {
		if err := secondTask(); err != nil {
			return err
		}
		return nil
	})

	// Wait for all goroutines to finish and return the first error (if any)
	return group.Wait()
}

func firstTask(ctx context.Context) {
	var counter int
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println("some task")
			if counter > 30 {
				return
			}
			counter++
		}
	}
}

func secondTask() error {
	time.Sleep(3 * time.Second)
	return errFailure
}
