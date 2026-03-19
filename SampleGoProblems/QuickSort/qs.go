package main

import (
	"fmt"
)

func partition(nums []int, left, right int) int {
	i := left
	pivot := nums[right]

	for j := left; j < right; j++ {
		if nums[j] <= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

func quicksort(nums []int, left, right int) {
	if left < right {
		pivotIndex := partition(nums, left, right)
		quicksort(nums, left, pivotIndex-1)
		quicksort(nums, pivotIndex+1, right)
	}
}

func main() {
	nums := []int{3, 2, 6, 4, 0, 7, 1}
	quicksort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
