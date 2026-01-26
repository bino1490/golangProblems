package main

import "fmt"

func add(a, b int) int {
	return a + b
}

/*
A higher-order function is a function that:
Takes one or more functions as arguments OR
Returns a function
*/

// calculate is a higher-order function
func calculate(op func(int, int) int) int {
	return op(2, 3)
}

func main() {

	/*A first-class function is a function that can be:
	Assigned to a variable
	Passed as an argument
	Returned from another function*/

	var f func(int, int) int // -> This declares a variable that can hold a function value. This line enables first-class usage, but still no function value yet.
	f = add                  // -> This is first-class function usage,  add is treated as a value

	fmt.Println(f(2, 3)) // 5

	result := calculate(add) // passing function as argument
	fmt.Println(result)      // 5
}
