package main

import "fmt"

func moveZeroes(nums []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			per := nums[k]
			nums[k] = nums[i]
			nums[i] = per
			k++
		}
	}
}

func main() {
	per := []int{0,1,0,3,12}
	moveZeroes(per)
	fmt.Println(per)
}
