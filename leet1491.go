package main

import (
	"fmt"
	"time"
	"sort"
)


func timeTaken() func() {
    start := time.Now()

    return func() {
        elapsed := time.Since(start)
        fmt.Println(elapsed)
    }
}

func main() {
	defer timeTaken()()
	arr := []int{1000, 4000, 6000, 11000, 13000, 17000, 28000, 31000, 33000, 37000, 39000, 45000, 48000, 54000, 59000, 67000, 77000, 78000, 93000, 99000}
	sort.Int(arr[:])
	last := len(arr)-2
	var total float64 = 0
	for i := 1; i <= last; i++ {
		total += float64(arr[i])
	}
	fmt.Println(float64(total)/float64(last))

	
	// for _, value := range arr {
	// 	fmt.Println(value)
	// }
}
