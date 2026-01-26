package main

import "fmt"

/*
Definition: Returning from a function without explicitly specifying return values, using named return variables.

A naked return is when a Go function returns without saying what it’s returning.

When it works: The function must have named return variables.
*/

func addAndMultiply(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return // naked return, uses sum and product
}

func main() {
	s, p := addAndMultiply(3, 4)
	fmt.Println("Sum:", s, "Product:", p)
}
