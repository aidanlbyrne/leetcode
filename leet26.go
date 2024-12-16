package main

import (
	"fmt"
)

// func removeDuplicates(nums []int) int {
//     uI := make(map[int]bool)
//     var que []int
//     for _, n := range nums {
//         if !uI[n] {
//             uI[n] = true
//             que = append(que, n)
//         }
//     }
// 	nums = append(que, nums...)
//     k := len(que)
//     return k

// }

func removeDuplicates(nums []int) int {
	k := 0
	for _, n := range nums {
		if nums[k] != n {
			k++
			nums[k] = n
		}
	}
	return k + 1
}

func main() {
	nums := []int{1,1,2}

	k := removeDuplicates(nums)

	fmt.Println("k: ",k)
	
}