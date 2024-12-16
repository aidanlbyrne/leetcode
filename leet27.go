package main

import (
	"fmt"
)
// ****** in worked for the first part of the question, which isnt impressive. it failed making
// ****** the first k elements of nums all values != val in the way the question asked
// ****** this method focuses on elements that == val, when we really only care about 
// ****** elements != val
// func removeElement(nums []int, val int) int {
// 	replace := make([]int, 50)
// 	replaceCount := 0
// 	k := 0
// 	for i, n := range nums {
// 		if n == val {
// 			replace[k] = i
// 			k++
// 			continue
// 		}
// 		if k > replaceCount {
// 			nums[replace[replaceCount]] = n
// 			replaceCount++
// 		}
// 	}
// 	fmt.Println(nums)
// 	return k
// }

func removeElement(nums []int, val int) int {
	k := 0
	for _, n := range nums {
		if n != val {
			nums[k] = n
			k++
			continue
		}
		
	}
	fmt.Println(nums)
	return k
}

func main() {
	val := 2
	// nums := []int{3,2,2,3}
	nums := []int{0,1,2,2,3,0,4,2}

	result := removeElement(nums, val)
	fmt.Println(result)
}