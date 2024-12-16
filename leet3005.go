package main

import "fmt"


func maxFrequencyElements(nums []int) int {
	freqM := make(map[int]int)
	max := 0
	freq := 1
	for _, num := range nums {
		freqM[num]++
		if freqM[num] >= max {
			if freqM[num] > max {
				max = freqM[num]
				continue
			}
			if freqM[num] == max {
				freq++
			}
		}
	}
	fmt.Println(max)
	fmt.Println(freq)
	return max * freq
}

func main() {
	nums := []int{1,2,2,3,1,4}
	result := maxFrequencyElements(nums)
	fmt.Println(result)
}