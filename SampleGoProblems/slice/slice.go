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
	i := make([]int, 3, 8)
	fmt.Println(i)
	// Since i still has spare capacity:
	// append does NOT allocate a new array
	// j and i share the same backing array
	j := append(i, 4) // -> append is variadic function
	fmt.Println(i)
	fmt.Println(j)
	// Important: you append to i again, not j.
	// i still has capacity
	// append writes at index 3
	// This overwrites the same slot where 4 was stored
	k := append(i, 5)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println("------------------------")

	a := make([]int, 3, 3)
	fmt.Println(a)
	// Since capacity is full: Go allocates a new backing array
	// b gets its own array
	// a stays unchanged
	b := append(a, 4)
	fmt.Println(a)
	fmt.Println(b)
	// Since capacity is full: Go allocates a new backing array
	// c gets its own array
	// a stays unchanged
	c := append(a, 5)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	// i doesn't change because append does not change the slice header (len) of i.

	// Desclaimer: Do not append value to other slice
	// i = append(i, 4) -> Ok
	// a = append(i,4) -> try not to use this
}
