package main

import (
	"fmt"
)


func mySqrt(x int) int {
    start := 0
    end := x
    var mid, ans int
    for start <= end {
        mid = (start + end)/2 // if mid gives value, return it
        if mid*mid == x {
            return mid
        }
        if (mid * mid) < x {
            start = mid+1		// if 'mid' gives lower result, we simply discard all the values lower than mid
            ans = mid			// an extra pointer 'ans' is maintained to keep track of only lowest 'mid' value
        } else {
            end = mid - 1		// if 'mid' value gives greater result, we simply discard all the values greater than mid
        }
    }
    return ans
}

func main() {
	x := 8
	result := mySqrt(x)

	fmt.Println(result)
}