package main

import "fmt"

func main() {
	//var slice []string
	//
	//for i := 1; i < 1300; i++ {
	//	slice = append(slice, fmt.Sprintf("rec: %d", i))
	//	fmt.Println(fmt.Sprintf("len: %d, cap: %d", len(slice), cap(slice)))
	//}

	str := "G☺x"

	for i, s := range str {
		fmt.Println(i, s)
	}
	//
	//s := str[0]
	//
	//fmt.Println("len:", len(str))
	//fmt.Printf("%c of type %T", s, s)
}
