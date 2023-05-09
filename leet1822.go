package main

import (
	"fmt"
	"time"
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
	nums := []int64{-1, 0, 0, 0, -2, 1, 2, 2, 2, 2, 1, 2, 2, -1, 1, 1, 2, 2, 1, 0, 2, 1, 2, -1, -2, 2, 2, -2, -2, -2, 2, 0, 2, -2, -2, 1, -1, 0, -2, -1, 0, 2, -1, 1, 0, -1, 0, 0, 1, -1, 1, 2, -1, 1, -1, -1, -1, -1, -2, 2, -2, -2, 2, 2, 1, 2, 1, 2, -2, -2, -1, 0, 2, 0, 2, -1, -1, 0, 1, -2, 0, -1, 0, 2, 1, 0, 0, 0, 0, -1, -1, -1, 0, 0, 2, -1, 2, -2, -2, -2}
	// var prod int64 = 1
	// for _, num := range nums {
	// 	if num != 0 {
	// 		prod *= num
	// 		// fmt.Println(prod)
	// 	}
	// }
	// if prod >= 1 {
	// 	fmt.Println(1)
	// } else {
	// 	fmt.Println(-1)
	// }
	posCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			posCount++
		}
	}
	if posCount%2 == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(1)
	}
}