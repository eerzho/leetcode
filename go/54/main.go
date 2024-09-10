package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	if n == 0 {
		return []int{}
	}

	m := len(matrix[0])
	if m == 0 {
		return []int{}
	}

	f := matrix[0][0]
	result := []int{}

	return append([]int{f}, result...)
}

func main() {
	fmt.Println("case 1")
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("result: ", spiralOrder(matrix))

	fmt.Println("case 2")
	matrix = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println("result: ", spiralOrder(matrix))
}
