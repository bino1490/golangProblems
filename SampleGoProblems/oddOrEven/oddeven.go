package main

import (
	"fmt"
)

func oddOrEven(nums []int, ch chan string) {
	for _, n := range nums {
		if n%2 == 0 {
			ch <- fmt.Sprintf("%d is even", n)
		} else {
			ch <- fmt.Sprintf("%d is odd", n)
		}
	}
	close(ch)
}

func main() {
	nums := []int{2, 3, 4, 5, 6}
	ch := make(chan string)

	go oddOrEven(nums, ch)

	for msg := range ch { // -----------> this is important ch not <-ch
		fmt.Println(msg)
	}
}
