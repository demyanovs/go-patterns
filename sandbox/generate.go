package main

import "fmt"

func main() {
	fmt.Println(generate(2, 5))
}

// Сгенерировать случайную матрицу размером n на m с неповторяющимися числами
func generate(n int, m int) [][]int {
	matrix := make([][]int, n)
	lastNum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			lastNum++
			matrix[i] = append(matrix[i], lastNum)
		}
	}

	return matrix
}
