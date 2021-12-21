package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum, mul := calc(os.Args[1], os.Args[2])
	fmt.Println("Sum: ", sum)
	fmt.Println("Mul: ", mul)
}

func calc(a, b string) (sum int, mul int) {
	int1, _ := strconv.Atoi(a)
	int2, _ := strconv.Atoi(b)
	sum = int1 + int2
	mul = int1 * int2
	return sum, mul
}
