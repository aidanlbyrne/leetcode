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


func main() {
	defer timeTaken()()
	s, t := "anagram", "nagaram"
	// fmt.Println(s, t)
	count1 := make(map[byte]int)
	count2 := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		count1[s[i]] = 1 + count1[s[i]]
		count2[t[i]] = 1 + count2[t[i]]
	}
	if count1 == count2 {
		fmt.Println("hey")
	}
	fmt.Println(count1)
	fmt.Println(count2)

}
