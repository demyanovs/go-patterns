package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Boat 🛥."
	fmt.Println("length of string?", len(s))
	fmt.Println("true length", utf8.RuneCountInString(s))

	for i, c := range s {
		fmt.Printf("Positiion %d of '%s'\n", i, string(c))
	}
}
