package main

import (
	"fmt"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // acts like infinity
	}

	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1) // --> core formula
			}
		}
	}
	fmt.Println(dp)
	if dp[amount] > amount {
		return -1
	}
	fmt.Println(dp)
	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11

	fmt.Println(coinChange(coins, amount)) // 3
}
