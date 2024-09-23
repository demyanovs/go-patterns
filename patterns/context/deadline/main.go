package deadline

func handler(ctx context.Context, ch chan Message) error {
	for {
		select {
		case msg := <-ch:
			// Do something with msg
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
