package main

import "fmt"

func sunoftwo(nums []int, k int) []int {
	x := make(map[int]int)
	for i, val := range nums {

		if v, ok := x[val]; ok {
			return []int{v, i}
		}
		x[k-val] = i
	}
	return nil
}

func main() {
	fmt.Println(sunoftwo([]int{2, 7, 11, 15}, 19)) //Ans 1,2 -> 7+11=19

}
