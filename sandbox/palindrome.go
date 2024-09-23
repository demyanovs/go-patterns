package main

import "fmt"

func main() {
	res := isPalindrome("abbccbba")
	fmt.Println("res:", res)
}

func isPalindrome(str string) bool {
	var reversed string

	for i := len(str) - 1; i > -1; i-- {
		reversed = reversed + string(str[i])
	}

	fmt.Println("original", str)
	fmt.Println("reversed", reversed)

	return str == reversed
}
