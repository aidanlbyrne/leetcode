package main

import (
	"fmt"
	"time"
	// "sort"

)

func timeTaken() func() {
    start := time.Now()

    return func() {
        elapsed := time.Since(start)
        fmt.Println(elapsed)
    }
}

func main() {
	defer timeTaken()()
	
	mat := [][]int{{7,3,1,9}, {3,4,6,9}, {6,9,6,6}, {9,5,8,5}}
	tot := 0
	m := len(mat)-1
	for i := range mat {
		tot += mat[i][i]
		tot += mat[i][m - i]
	}

	fmt.Println(tot)
}