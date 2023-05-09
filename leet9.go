package main

import (
	"fmt"
	"math"
)



func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	len := int(math.Log10(float64(x)))+1
	if len == 1 {
		return true
	}
	pal := 0
	pal1 := x
	for i := 0; i < len; i++ {
		pal = pal*10 + x%10
		x /= 10
	}
	
	return pal == pal1

}


func main() {
	result := isPalindrome(123321)

	fmt.Println(result)
}