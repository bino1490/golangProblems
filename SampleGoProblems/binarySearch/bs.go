package main

import "fmt"

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 3, 5, 7, 9}
	target := 7

	fmt.Println("Index:", binarySearch(nums, target))
}
