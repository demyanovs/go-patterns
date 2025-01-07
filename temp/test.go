package main

import (
	"fmt"
)

func main() {
	fmt.Println(3 % 5)
	//fizzBuzz(15)
}

func fizzBuzz(n int32) {
	for i := 1; i <= int(n); i++ {
		// Write your code here
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}

		if n%3 == 0 && n%5 != 0 {
			fmt.Println("Fizz")
			continue
		}

		if i%3 != 0 && i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}

		fmt.Println(i)
	}
}
