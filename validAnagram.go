package main

import (
	"fmt"
	"reflect"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s1 := make([]int, 26)
	s2 := make([]int, 26)

	for i, v := range s {
		s1[v%26]++
		s2[t[i]%26]++
	}
	if reflect.DeepEqual(s1, s2) {
		return true
	}
	return false
}

func main() {
	s, t := "anagram", "nagaram"
	result := isAnagram(s, t)

}
// func main() {
// 	s, t := "anagram", "nagaram"

// 	s1 := make([]int, 122)
// 	s2 := make([]int, 122)

// 	for _,v := range s {
// 		s1[v]++
// 	}
// 	for _,v := range t {
// 		s2[v]++
// 	}

// 	if s1 == s2 {
// 		fmt.Println("you")
// 	}
// }