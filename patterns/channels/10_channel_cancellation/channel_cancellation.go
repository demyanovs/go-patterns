package main

import (
	"context"
	"net/http"
)

/**
With channel cancellation, I can take an existing channel being used already
for cancellation purposes (legacy code) and convert its use with a context,
where a context is needed for a future function call.
*/

func channelCancellation(stop <-chan struct{}) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		select {
		case <-stop:
			cancel()
		case <-ctx.Done():
		}
	}()

	func(ctx context.Context) error {
		req, err := http.NewRequestWithContext(
			ctx,
			http.MethodGet,
			"https://www.ardanlabs.com/blog/index.xml",
			nil,
		)

		if err != nil {
			return err
		}
		_, err = http.DefaultClient.Do(req)

		if err != nil {
			return err
		}
		return nil
	}(ctx)
}
