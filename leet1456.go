package main

import (
	"fmt"
)


func main() {
	s := "abciiidef"
	k := 3
	res := maxVowels(s, k)

	fmt.Println(res)
}

// first completed attempt: time complexity too large; timeout error


// func maxVowels(s string, k int) int {
//     maxLen := 0
// 	tempLen := 0
//  	vow := map[byte]int {
// 		'a' : 1,
// 		'e' : 2,
// 		'i' : 3,
// 		'o' : 4,
// 		'u' : 5,
// 	}

//     for i := range s {
//         if _, ok := vow[s[i]]; ok {
//             for j := i; j < i + k; j++ {
// 				if j == len(s) - 1  {
// 					break
// 				}
//                 if _, ok := vow[s[j]]; ok {
// 					tempLen++
// 					if tempLen == k{
// 						return k
// 					}
//                     if tempLen > maxLen {
// 						maxLen = tempLen
//                     }
// 					continue
// 				}
// 				tempLen=0
//             }
//         }
//     }
//  return maxLen
// }
