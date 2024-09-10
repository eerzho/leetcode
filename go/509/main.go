package main

import "fmt"

func fib(n int) int {
	s := 1
	result := 0
	for i := 0; i < n; i++ {
		per := result
		result += s
		s = per
	}

	return result
}

func main() {
	fmt.Println("case 1")
	n := 2
	fmt.Println("result: ", fib(n))

	fmt.Println("case 2")
	n = 3
	fmt.Println("result: ", fib(n))

	fmt.Println("case 3")
	n = 4
	fmt.Println("result: ", fib(n))

	fmt.Println("case 4")
	n = 0
	fmt.Println("result: ", fib(n))
}