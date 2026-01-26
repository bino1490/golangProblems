package main

import "fmt"

// Currying means breaking a function with many inputs into smaller functions that take one input at a time.
//It lets you reuse functions by filling in arguments step by step.

// Curried function
func order(size string) func(string) func(string) string {
	return func(kind string) func(string) string {
		return func(sugar string) string {
			return size + " " + kind + " coffee with " + sugar + " sugar"
		}
	}
}

func main() {
	// Step 1: choose size
	orderLarge := order("large")

	// Step 2: choose type
	orderLargeLatte := orderLarge("latte")

	// Step 3: choose sugar
	finalOrder := orderLargeLatte("less")

	fmt.Println(finalOrder)
}
