package main

func twoSum(nums []int, target int) []int {
	sum := make(map[int]int)
	for i, num := range nums {
		if mi, ok := sum[target - num]; ok {
			return []int{mi, i}
		}
		sum[num] = i
	}
	return nil
}

func main() {
	result := twoSum(...)
	
}