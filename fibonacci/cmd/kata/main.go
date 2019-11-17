package main

import (
	"fmt"

	"github.com/golang-katas/fibonacci/internal"
)

func main() {
	var fibonacci internal.Fibonacci

	for i := 0; i < 20; i++ {
		fmt.Printf("%d is the Fibonnaci result for %d steps\n", fibonacci.Calculate(i), i)
	}
}
