package main

import (
	"fmt"
)

func strStr(haystack, needle string) int {

	for i := 0; i < len(haystack) - len(needle); i++ {
		for j := 0; j <= len(needle); j++ {
			if j == len(needle) {
				return i
			}
			if haystack[i+j] != needle[j] {
				break
			}
		}
	}
	return -1
}

func main() {
	x := "mississippi"
	y := "pi"

	res := strStr(x, y)
	fmt.Println(res)
}
