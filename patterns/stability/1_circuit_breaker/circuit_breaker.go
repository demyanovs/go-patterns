package main

/*
Circuit Breaker automatically degrades service functions in response to a likely fault,
preventing larger or cascading failures by eliminating recurring errors and providing reasonable error responses.
*/

import (
	"context"
	"errors"
	"sync"
	"time"
)

type Circuit func(context.Context) (string, error)

func main() {
	//ctx := context.Background()
	//circuit := func(ctx context.Context) (string, error) {
	// return "", nil
	//}
	//
	//res := Breaker(circuit, 3)
	//
}

func Breaker(circuit Circuit, failureThreshold uint) Circuit {
	var consecutiveFailures int = 0
	var lastAttempt = time.Now()
	var m sync.RWMutex

	return func(ctx context.Context) (string, error) {
		m.RLock() // Establish a "read lock"

		d := consecutiveFailures - int(failureThreshold)

		if d >= 0 {
			shouldRetryAt := lastAttempt.Add(time.Second * 2 << d)
			if !time.Now().After(shouldRetryAt) {
				m.RUnlock()
				return "", errors.New("service unreachable")
			}
		}

		m.RUnlock() // Unlock read lock

		response, err := circuit(ctx) // Issue request proper

		m.Lock() // Lock around shared resources
		defer m.Unlock()

		lastAttempt = time.Now() // Record time of attempt

		if err != nil { // Circuit returned an error,
			consecutiveFailures++ // so we count the failure
			return response, err  // and return
		}

		consecutiveFailures = 0 // Reset failures counter

		return response, nil
	}
}
