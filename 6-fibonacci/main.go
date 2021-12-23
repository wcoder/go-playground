package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("What's the Fibonacci sequence you want? ")
	fmt.Scanln(&n)
	fmt.Println("The Fibonacci sequence is:", fibonacci(n))
}

func fibonacci(n int) []int {
	if n < 2 {
		return make([]int, 0)
	}

	nums := make([]int, n)
	nums[0], nums[1] = 1, 1

	for i := 2; i < n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}

	return nums
}
