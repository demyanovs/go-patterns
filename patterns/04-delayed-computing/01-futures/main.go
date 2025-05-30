package main

import (
	"fmt"
	"time"
)

type data struct {
	Body  string
	Error error
}

func main() {
	future1 := future("https://example1.com")
	future2 := future("https://example2.com")

	fmt.Println("Requests started")

	body1 := <-future1
	body2 := <-future2

	fmt.Printf("Response 1: %v\n", body1)
	fmt.Printf("Response 2: %v\n", body2)
}

func future(url string) <-chan data {
	resultChan := make(chan data, 1)

	go func() {
		body, err := doGet(url)

		resultChan <- data{Body: body, Error: err}
	}()

	return resultChan
}

func doGet(url string) (string, error) {
	time.Sleep(time.Millisecond * 200)
	return fmt.Sprintf("Response of %s", url), nil
}

/*
Или промис.
Позволяет запустить вычисление некоторых данных в фоне не дожидаясь их обработки. В случае, если мы знаем, что эти
данные нам потребуются в дальнейшем, но не прямо сейчас.
*/
