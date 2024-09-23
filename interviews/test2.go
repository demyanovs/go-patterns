package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const timeout = 5
const retryCount = 5

func main() {
	ctx := context.Background()
	fmt.Printf("Discount: %v\n", getDiscount(ctx, retryCount))
}

func getDiscount(ctx context.Context, retryCount int) any {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://discount-service/discount", nil)
	if err != nil {
		return err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if retryCount > 0 {
			time.Sleep(1 * time.Second)
			return getDiscount(ctx, retryCount-1)
		}

		return err
	}

	discount := resp.Body

	//discount, _ := http.Get("http://discount-service/discount")

	return discount
}
