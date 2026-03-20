package main

import "fmt"

func movingAverage(nums []int, k int) []float64 {
	var result []float64
	windowSum := 0

	for i := 0; i < len(nums); i++ {
		windowSum += nums[i]

		// When window size reached
		if i >= k-1 {
			avg := float64(windowSum) / float64(k)
			result = append(result, avg)

			// Slide window: remove left element
			windowSum -= nums[i-k+1]
		}
	}

	return result
}

func main() {
	nums := []int{2, 5, 1, 6, 7}
	k := 3

	fmt.Println(movingAverage(nums, k))
}
