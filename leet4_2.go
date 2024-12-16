package main

import (
	"fmt"
)


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1i, n2i, l1, l2 := 0, 0, len(nums1), len(nums2)
	l3 := ((l1+l2)/2)+1
	
	nums3 := make([]int, l3)
	for n3i := 0; n3i < l3; n3i++ {
		if n1i < l1 && n2i < l2 {
			if nums1[n1i] <= nums2[n2i] {
				nums3[n3i] = nums1[n1i]
				n1i++
			} else {
				nums3[n3i] = nums2[n2i]
				n2i++
			}
		} else {
			if n1i < l1 {
				nums3[n3i] = nums1[n1i]
				n1i++
			} else {
				nums3[n3i] = nums2[n2i]
				n2i++
			}
		}
	}
	if even {
		return (float64(nums3[l3-2])+float64(nums3[l3-1]))/2
	}
	return float64(nums3[l3-1])
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3}

	res := findMedianSortedArrays(nums1, nums2)

	fmt.Println(res)
}
