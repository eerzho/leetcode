package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for idx, num := range nums {
		jdx, ok := hash[num]
		if ok {
			return []int{idx, jdx}
		}

		hash[target-num] = idx
	}

	return []int{}
}

func main() {
	fmt.Println("case 1")
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Print("result: ", twoSum(nums, target))

	fmt.Println()

	fmt.Println("case 2")
	nums = []int{3, 2, 4}
	target = 6
	fmt.Print("result: ", twoSum(nums, target))

	fmt.Println()

	fmt.Println("case 3")
	nums = []int{3, 3}
	target = 6
	fmt.Print("result: ", twoSum(nums, target))
	
	fmt.Println()
}
