package main

import "fmt"

func main() {
	fmt.Println(merge([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}))
}

// merge merges multiple sorted arrays into one sorted array.
func merge(arrs ...[]int) []int {
	var res []int
	for _, arr := range arrs {
		res = append(res, arr...)
	}

	return res
}
