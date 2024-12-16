package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	count := 0
	for i := len(s)-1; i >= 0; i-- {
		fmt.Println(s[i])
		if s[i] != 32 {
			count++
			continue
		}
		if count > 0 {
			break
		}
	}
	return count
}

func main() {
	// s := "   fly me   to  the moon  "
	s := "a"

	res := lengthOfLastWord(s)
	fmt.Println(res)
}