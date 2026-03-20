package main

import "fmt"

func quickselct(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}

	pivotIndex := partition(nums, left, right)
	if k == pivotIndex {
		return nums[k]
	} else if k < pivotIndex {
		return quickselct(nums, left, pivotIndex-1, k)
	} else {
		return quickselct(nums, pivotIndex+1, right, k)
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left

	for j := left; j < right; j++ {
		if nums[j] <= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

func main() {

	nums := []int{3, 2, 5, 7, 6, 0, 1}
	k := 3 // Ans:2
	val := quickselct(nums, 0, len(nums)-1, k-1)
	fmt.Println(val)

}
