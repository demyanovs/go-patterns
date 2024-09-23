package main

import (
	"context"
	"fmt"
	"time"
)

/*
Берет один входной канал и произвольное кол-во выходных каналов и дублирует каждый вход в каждый выход.
Отличие от fdn-out - здесь сообщение в каждый канал дублируется. В случае с fan-out не зависимо от кол-ва
выходных каналов только один из них получит копию сообщения.
*/

func tee(ctx context.Context, input <-chan string, outputs []chan<- string) {
	for elem := range input {
		elem := elem
		for _, out := range outputs {
			out := out
			go func() {
				select {
				case out <- elem:
					break
				case <-ctx.Done():
					break
				}
			}()
		}
	}
}

func main() {
	in := make(chan string)
	out1 := make(chan string)
	out2 := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	outputs := []chan<- string{
		out1, out2,
	}

	go func() {
		in <- "A"
		in <- "B"
		in <- "C"
		close(in)
	}()

	tee(ctx, in, outputs)
	go func() {
		for {
			select {
			case <-ctx.Done():
				break
			case v := <-out1:
				fmt.Printf("out1 got value: %s\n", v)
			case v := <-out2:
				fmt.Printf("out2 got value: %s\n", v)
			}
		}
	}()

	time.Sleep(time.Second)
	cancel()
}
