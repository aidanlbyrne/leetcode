package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	l := len(needle)
	for i := 0; i < len(haystack)-l+1; i++ {
		fmt.Println(string(haystack[i]))
		if haystack[i] == needle[0] {
			for j := 1; j <= l; j++ {
				if j == l {
					return i
				}
				if haystack[i+j] != needle[j] {
					break
				}
			}
		}
	}
	return -1
}

func main() {
	haystack := "mississippi"
	needle := "pi"

	res := strStr(haystack, needle)

	fmt.Println(res)
}