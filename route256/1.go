/*
Напишите программу, которая выводит сумму двух целых чисел.
Входные данные
В первой строке входных данных содержится целое число t
(1<=t<=104) - количество наборов входных данных в тесте.

Далее следуют описания t наборов входных данных, один набор в строке.
В первой (и единственной) строке набора записаны два целых числа a и b
(−1000<=a,b<=1000).

Выходные данные
Для каждого набора входных данных выведите сумму двух заданных чисел, то есть a+b.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var n, m int
		fmt.Fscan(in, &n, &m)
		fmt.Fprintln(out, n+m)
	}
}
