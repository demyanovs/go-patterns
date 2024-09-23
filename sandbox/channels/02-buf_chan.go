package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	close(messages)

	time.Sleep(time.Second * 1)

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
