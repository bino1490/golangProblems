package main

import (
	"fmt"
	"strings"
)

/*
Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.



Example 1:

Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.
Example 2:

Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.

*/

func strStr(haystack string, needle string) int {
	if ok := strings.Contains(haystack, needle); ok {
		for i := 0; i < len(haystack)-len(needle)+1; i++ {
			if haystack[i:len(needle)+i] == needle {
				return i
			}

		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("sadtest", "sad"))
	fmt.Println(strStr("testsad", "sad"))
	fmt.Println(strStr("leettotest", "leetti"))
}
