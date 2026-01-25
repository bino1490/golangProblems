// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func reverseString(n string) string {

	// Way 1
	/*
		var rs string
		for _, v := range n {
			rs = string(v) + rs
		}*/

	// Alternative way
	var rs string
	l := len(n)
	for i := l - 1; i >= 0; i-- {
		rs = rs + string(n[i])
	}

	return rs
}

func main() {
	fmt.Println("Reverse String", reverseString("sad"))
}
