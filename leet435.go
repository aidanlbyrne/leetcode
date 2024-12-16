package main

import (
	"fmt"
)

func main() {
	strs := []string{"flower", "flight", "box"}
	var res string

	for i := 1, i < len(strs); i++ {
		if strs[0][0] != strs[i][0] {
			return ""
		}
		for j := len(res)
	}
	fmt.Println(strs[0][0])
}