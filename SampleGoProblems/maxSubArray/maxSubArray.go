package main

import "fmt"

// Given an array of integers (can be +ve, -ve, 0):

// Find a continuous subarray that has the maximum sum
// Return that maximum sum

func maxSubArray(nums []int) int {
	maxSoFar := nums[0]
	curr := nums[0]

	for i := 1; i < len(nums); i++ {
		if curr < 0 {
			curr = nums[i]
		} else {
			curr += nums[i]
		}

		if curr > maxSoFar {
			maxSoFar = curr
		}
	}
	return maxSoFar
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("Max Subarray Sum:", maxSubArray(nums))
}
