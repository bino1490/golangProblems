package main

import "fmt"

func variadic(a ...int) []int {
	fmt.Println(a)
	return a
}

func main() {

	a := [][]int{{1}, {2, 3}}
	fmt.Println(a[1])
	fmt.Println(a[1][0])

	fmt.Println(variadic(1, 2, 3, 4))
}
