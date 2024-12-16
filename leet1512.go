package main

import "fmt"

func numIdenticalPairs(nums []int) int {
    m := make(map[int]int)
    count := 0
    for _, v := range(nums) {
        if v, ok := m[v]; ok {
            count += m[v]
            m[v]++
            continue
        }
        m[v] = 1
    }
    return count
}

func main() {
	nums := []int{1,1,1,1}
	res := numIdenticalPairs(nums)

	fmt.Println(res)
}