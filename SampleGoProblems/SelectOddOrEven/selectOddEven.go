package main

import (
	"fmt"
)

var odd = make(chan string)
var even = make(chan string)

func oddoreven(nums []int) {
	for _, n := range nums {
		if n%2 == 0 {
			even <- fmt.Sprintf("%d is even", n)
		} else {
			odd <- fmt.Sprintf("%d is odd", n)
		}
	}
	close(even) // ---------------> close at sender end
	close(odd)
}

func main() {
	nums := []int{2, 3, 4, 5, 6}

	go oddoreven(nums)

	for even != nil || odd != nil {
		select {
		case msg, ok := <-even: // ------------> this is important in select this is <- channel and return value and ok
			if !ok {
				even = nil
				continue
			}
			fmt.Println(msg)

		case msg, ok := <-odd:
			if !ok {
				odd = nil
				continue
			}
			fmt.Println(msg)
		}
	}
}
