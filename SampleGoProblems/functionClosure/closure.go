package main

import "fmt"

/*

A closure is essentially a function value that “captures” variables from its surrounding scope. In simpler terms, it’s a function that can access and manipulate variables that were defined outside of it, even after the outer function has finished executing.
Closures capture variables by reference, not by value.

Each call to the outer function can produce a separate closure with its own captured variables.

Useful for:
	Maintaining state between function calls
	Creating function factories
	Functional-style programming
*/

func sentenceBuilder() func(string) string {
	sentence := ""

	return func(word string) string {
		sentence += word + " "
		return sentence
	}
}

func main() {
	build := sentenceBuilder()

	build("Go")
	build("closures")
	build("are")
	build("really")
	build("cool")

	fmt.Println(build("!"))
}
