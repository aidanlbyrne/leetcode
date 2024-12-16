package main

import "fmt"
import "sort"


func main() {
	nums := []int{4, 2, 3, 5}

	sort.Slice(nums, func(i,j int)bool{
		return nums[i] < nums[j]
	})

	for left,v := range nums {
		fmt.Println(left, v)
	}
}