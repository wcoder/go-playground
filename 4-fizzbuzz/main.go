package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz(i))
	}
}

func fizzbuzz(num int) string {
	switch {
	case num%15 == 0:
		return "FizzBuzz"
	case num%3 == 0:
		return "Fizz"
	case num%5 == 0:
		return "Buzz"
	}
	return fmt.Sprintf("%d", num)
}
