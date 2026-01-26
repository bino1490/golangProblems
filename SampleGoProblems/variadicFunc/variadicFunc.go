// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

// variadic function
func nums(n ...int) []int {
	return n
}

func nums1(n []int) []int {
	return n
}

func main() {

	fmt.Println(nums(1, 2, 3))
	fmt.Println(nums1([]int{1, 2, 3}))

	s := []int{2, 4, 6}
	// spread operator > s...
	fmt.Println(nums(s...))

}
