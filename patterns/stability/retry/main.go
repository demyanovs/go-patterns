package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

type Effector func(ctx context.Context) (string, error)

var count int

func main() {
	r := Retry(EmulateTransientError, 5, 2*time.Second)
	res, err := r(context.Background())

	fmt.Println(res, err)
}

func EmulateTransientError(ctx context.Context) (string, error) {
	count++

	if count <= 3 {
		return "intentional fail", errors.New("error")
	} else {
		return "success", nil
	}
}

func Retry(effector Effector, retries int, delay time.Duration) Effector {
	return func(ctx context.Context) (string, error) {
		for r := 0; ; r++ {
			response, err := effector(ctx)
			if err == nil || r >= retries {
				return response, err
			}

			log.Printf("Attempt %d failed; retrying in %v", r+1, delay)

			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return "", ctx.Err()
			}
		}
	}
}
