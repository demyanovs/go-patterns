package main

import (
	"fmt"
)

/*
Это структура данных, которая рассматривается как кольцевая, хотя ее реализация является линейной.
Кольцевой буфер является распространенной реализацией очереди. Используется он в основном в случаях, когда
нам нужно успевать обрабатывать входной поток данных, где некоторые элементы входного потока могут быть
без последствий отброшены. Таким образом прореживая слишком быстро поступающий поток. Конечно, если потребитель
успевает отрабатывать все приходящие данные, то прореживания данных не будет.
*/

func NewRingBuffer(inCh, outCh chan int) *ringBuffer {
	return &ringBuffer{
		inCh:  inCh,
		outCh: outCh,
	}
}

type ringBuffer struct {
	inCh  chan int
	outCh chan int
}

func (r *ringBuffer) Run() {
	defer close(r.outCh)
	for v := range r.inCh {
		select {
		case r.outCh <- v:
		default:
			<-r.outCh
			r.outCh <- v
		}
	}
}

func main() {
	inCh := make(chan int)
	outCh := make(chan int, 4)
	rb := NewRingBuffer(inCh, outCh)
	go rb.Run()

	for i := 1; i <= 10; i++ {
		inCh <- i
	}

	close(inCh)

	for res := range outCh {
		fmt.Println(res)
	}
}
