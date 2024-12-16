package main

import (
	"fmt"
)


func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	str := []byte(strs[0])

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) == 0 {
			return ""
		}
		for j := len(str)-1; j > 0; j-- {
			// fmt.Println(j)
			if j > len(strs[i])-1 {
				j = len(strs[i])-1
			}
			if str[j] != strs[i][j] {
				if j == 1 {
					return ""
				}
				str = str[:j]
			} else {break}
		}
	}
	return string(str)
}

func main() {
	strs := []string{"cir", "car"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)
	
}