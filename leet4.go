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

func combine(ar1, ar2, []int) []int{
	i, j, l1, l2 := 0, 0, len(ar1), len(ar2)
	ar3 := make([]int, l1+l2)
	for k := 0; k <l1+l2; k++ {
		if i==l1 || i < l1 && j < l2 && ar1[i] > ar2[j] {
			ar3[k] = ar2[j]
			j++
			continue
		}
		if j == l2 || i < l1 && j < l2 && ar1[i] < = ar2[j] {
			ar3[k] = ar1[i]
			i++
		}
		return ar3
	}
}

func main() {
	defer timeTaken()()
	nums1 := []int{1, 2, 3, 3}
	nums2 := []int{3, 4, 5, 8}
	for _,num := range nums2 {
		nums1 = append(nums1[:], num)
	}
	mid := len(nums1)/2
	fmt.Println(nums1)
	fmt.Println((float64(nums1[mid]+nums1[mid - 1])/float64(2)))
}