package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	lcp := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j <= len(strs[i]) && j <= len(lcp); j++ {
			fmt.Println("current strs[i]: ", strs[i])
			if j == len(strs[i]) || j == len(lcp) {
				fmt.Println("j: ", j)
				fmt.Println(lcp)
				lcp = lcp[:j]
				fmt.Println(lcp)
				continue
			}
			// fmt.Println("should only happen twice")
			if lcp[j] == strs[i][j] {
				continue
			}
			lcp = lcp[:j]
			break
		}
		if len(lcp) == 0 {
			return ""
		}
	}
	return lcp
}


func main() {
	// strs1 := []string{"flower", "flow", "flight"}
	// strs2 := []string{"ab", "a"}
	strs3 := []string{"aaa", "aa", "aaa"}
	// result1 := longestCommonPrefix(strs1)
	// result2 := longestCommonPrefix(strs2)
	result3 := longestCommonPrefix(strs3)
	// fmt.Println(result1)
	// fmt.Println(result2)
	fmt.Println(result3)
}