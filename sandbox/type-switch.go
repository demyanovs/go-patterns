package main

import "fmt"

func main() {
	printType("hello")
	printType(32)
}

func printType(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%d is an integer\n", v)
	case string:
		fmt.Printf("%#v is a string\n", v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
