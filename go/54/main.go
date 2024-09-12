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

	result := []int{}
	left, right, top, bottom := 0, m-1, 0, n-1

	for {
		// ➡️
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}

		// ⬇️
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if left > right {
			break
		}

		// ⬅️
		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])
		}
		bottom--
		if top > bottom {
			break
		}

		// ⬆️
		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return result
}

func main() {
	fmt.Println("case 1")
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("result: ", spiralOrder(matrix))

	fmt.Println("case 2")
	matrix = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println("result: ", spiralOrder(matrix))
}
