package main

import (
	"fmt"
)

func isDupValExists(nums []int) bool {
	m := make(map[int]bool)

	for _, n := range nums {
		if m[n] {
			return true
		} else {
			m[n] = true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 4, 3, 1}

	fmt.Println("Dup Exists", isDupValExists(nums))

}
