package main

import "fmt"


func search(nums []int, target int) int {
	start := 0
	end := len(nums)
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2

	result := search(nums, target)
	fmt.Println(result)
}