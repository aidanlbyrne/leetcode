package main

import "fmt"

func maxDepth(s string) int{
	count := 0
	max := 0
	for _,c := range(s) {
		if c == 40 {
			count++
		}
		if c == 41 && count > 0{
			count--
		}
		if count > max {
			max = count
		}
	}
	return max
}
func main() {
	s := "(1)+(2*3)+(((3)))"
	res := maxDepth(s)
	fmt.Println(res)

}