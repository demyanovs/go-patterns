package main

import "fmt"

func main() {
	fmt.Println(test1())
	test2()
}

func test1() (result int) {
	defer func() {
		result++
	}()

	return 0
}

func test2() {
	var i1 int = 10
	var k = 20
	var i2 *int = &k

	// Что выведет?
	defer printInt("i1", i1)
	defer printInt("i2 as value", *i2)
	defer printIntPointer("i2 as pointer", i2)

	i1 = 1010
	*i2 = 2020
}

func printInt(v string, i int) {
	fmt.Printf("%s=%d\n", v, i)
}

func printIntPointer(v string, i *int) {
	fmt.Printf("%s=%d\n", v, *i)
}
