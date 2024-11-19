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
	err := FetchUserDataWithError(ctx)
	//err := FetchUserDataWithoutError(ctx)
	if err != nil {
		fmt.Println("Error fetching user data:", err)
	}

	fmt.Println("Done")
}

func FetchUserDataWithError(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)
	//group := errgroup.Group{}
	//group.SetLimit(1)

	// Run first periodic task.
	group.Go(func() error {
		firstTask(ctx)
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

func FetchUserDataWithoutError(ctx context.Context) error {
	var group errgroup.Group

	// Run first periodic task.
	group.Go(func() error {
		thirdTask(ctx)
		return nil
	})

	// Run second task.
	group.Go(func() error {
		fourthTask()
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
			if counter > 10 {
				return
			}
			counter++
		}
	}
}

func secondTask() error {
	fmt.Println("second task start")
	time.Sleep(3 * time.Second)
	fmt.Println("log error", errFailure)
	return errFailure
}

func thirdTask(ctx context.Context) {
	var counter int
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println("some task")
			if counter > 10 {
				fmt.Println("third task finished")
				return
			}
			counter++
		}
	}
}

func fourthTask() {
	fmt.Println("fourth task start")
	time.Sleep(3 * time.Second)
	fmt.Println("fourth task log error", errFailure)
	return
}
