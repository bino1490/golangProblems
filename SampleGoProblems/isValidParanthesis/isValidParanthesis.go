package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}

	// map closing → opening
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range s {
		// If opening bracket → push
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			// If stack empty → invalid
			if len(stack) == 0 {
				return false
			}

			// Check match
			top := stack[len(stack)-1]
			if mapping[ch] != top {
				return false
			}

			// Pop
			stack = stack[:len(stack)-1]
		}
	}

	// Valid only if stack empty
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("{{[]}}")) // true
}
