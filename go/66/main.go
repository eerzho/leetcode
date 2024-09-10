package main

import "fmt"

func plusOne(digits []int) []int {
	per := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] + per == 10 {
			digits[i] = 0
			per = 1
		} else {
			digits[i] += per
			per = 0
		}
	}

	if per != 0 {
		return append([]int{1}, digits...)
	}

	return digits
}

func main() {
	fmt.Println("case 1")
	digits := []int{1, 2, 3}
	fmt.Println("result: ", plusOne(digits))

	fmt.Println("case 2")
	digits = []int{4, 3, 2, 1}
	fmt.Println("result: ", plusOne(digits))

	fmt.Println("case 3")
	digits = []int{9}
	fmt.Println("result: ", plusOne(digits))

	fmt.Println("case 4")
	digits = []int{9, 9, 9}
	fmt.Println("result: ", plusOne(digits))
}
