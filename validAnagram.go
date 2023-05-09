package main

import "fmt"

func main() {
	s, t := "anagram", "nagaram"

	s1 := make([]int, 122)
	s2 := make([]int, 122)

	for _,v := range s {
		s1[v]++
	}
	for _,v := range t {
		s2[v]++
	}

	if s1 == s2 {
		fmt.Println("you")
	}

	
	
}