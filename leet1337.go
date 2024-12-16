package main

import (
	"fmt"
)

func kWeakestRows(mat [][]int, k int) []int {

	strMap := make(map[int]int)

	for rowI, row := range mat {
		count := 0
		for _, n := range row {
			if n == 0 {
				break
			}
			count++
		}
		strMap[count] = rowI
	}
}

func main() {
	mat := [][]int{{1,1,0,0,0},{1,1,1,1,0},{1,0,0,0,0},{1,1,0,0,0},{1,1,1,1,1}}
	// k := 3
	// res := kWeakestRows(mat, k)
	// fmt.Println(res)
}