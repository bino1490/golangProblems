// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// Check if left half is sorted
		if nums[left] <= nums[mid] {
			fmt.Println("goes here", mid)
			// Target lies in left half
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			fmt.Println("goes there", mid)
			// Right half is sorted
			// Target lies in right half
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func main() {
	a := []int{5, 6, 7, 8, 1, 2, 3, 4}
	k := 3 // Ans:=6
	x := search(a, k)
	fmt.Println(x)
}
