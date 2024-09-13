package main

import "fmt"

func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	result := []int{}

	for left <= right {
		leftN := nums[left] * nums[left]
		rightN := nums[right] * nums[right]

		if leftN > rightN {
			result = append([]int{leftN}, result...)
			left++
		} else {
			result = append([]int{rightN}, result...)
			right--
		}
	}

	return result
}

func main() {
	fmt.Println(sortedSquares([]int{-4,-1,0,3,10}))
}