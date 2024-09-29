package main

import "fmt"

func main() {
	var myArr [3]int

	var mySlice []int
	var mySlice2 []string
	mySlice3 := make([]string, 0)

	fmt.Printf("Arr equal %#v\n", myArr)

	fmt.Printf("Sl1 equal %#v\n", mySlice == nil)
	fmt.Printf("Sl2 equal %#v\n", mySlice2 == nil)
	fmt.Printf("Sl3 equal %#v\n", mySlice3 == nil)

}
