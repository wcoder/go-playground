package main

import (
	"fmt"

	"github.com/wcoder/calculator"
	"rsc.io/quote"
)

func main() {
	total := calculator.Sum(1, 2)
	fmt.Println(total)
	fmt.Println("Version: ", calculator.Version)
	fmt.Println(quote.Hello())
}
