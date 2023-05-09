package main

import (
	"fmt"
	// "time"
)

func main() {
	nums1 := []int{1, 5, 3, 1}
	nums2 := []int{3, 4, 5, 8}
	result := fmsa(nums1, nums2)
	fmt.Println(result)
}

func fmsa(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	m := len(nums2) 
	i, j := 0, 0
	r := make([]int, 0, n+m)
	for i < n && j < m {
		if nums1[i] < nums2[j] {
			r = append(r, nums1[i])
			i++
		} else {
			r = append(r, nums2[j])
			j++
		}
	}
	for i < n {
		r = append(r, nums1[i])
		i++
	}
	for j < m {
		r = append(r, nums2[j])
		j++
	}
	return r
}
