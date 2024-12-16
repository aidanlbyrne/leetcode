package main

import "fmt"

func sortByParity(nums[] int) []int {
	arr := make([]int, len(nums))
	front, back := 0, len(nums)-1

	for _,v := range(nums) {
		if v%2==0 {
			arr[front] = v
			front++
		} else {
			arr[back] = v
			back--
		}
	}
	return arr

}

func main() {
	nums := []int{1,2,4,5,3,6}
	res := sortByParity(nums)
	fmt.Println(res)

}