package main

import (
	// "strconv"
	"fmt"
	"time"
	// "sort"
	// "math"
)

func timeTaken() func() {
    start := time.Now()

    return func() {
        elapsed := time.Since(start)
        fmt.Println(elapsed)
    }
}

func first(m [][]int, c chan int) {
	sum := 0
	for i := range m {
		sum += m[i][i]
	}
	c <- sum
}

func second(m [][]int, c chan int) {
	sum := 0
	for i := len(m)-1; i >= 0; i-- {
		fmt.Println(i)
		fmt.Println(m[i][i])
		
		sum += m[i][i]
	}
	c <- sum
}

func main() {
	defer timeTaken()()

	mat := [][]int{{7,3,1,9}, {3,4,6,9}, {6,9,6,6}, {9,5,8,5}}
	c := make(chan int)
	go first(mat, c)
	go second(mat, c)
	x, y := <-c, <-c
	z := x+y
	fmt.Println(x, y, z)
}
