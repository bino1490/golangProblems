package main

import "fmt"

func slidingwindow(nums []int, k int) int {
	maxval := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		if i <= k-1 {

			sum += nums[i]
			continue
		}
		sum = sum + nums[i] - nums[i-k]
		if sum > maxval {
			maxval = sum
		}

	}
	return maxval
}

func main() {
	fmt.Println(slidingwindow([]int{2, 1, 5, 1, 3, 2}, 3)) //9

}
