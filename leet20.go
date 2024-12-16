package main


import (
	"fmt"
)

func isValid(s string) bool {
	
	apnd := map[byte]bool {
		40 : true,
		41 : false,
		91 : true,
		93 : false,
		123 : true,
		125 : false,
	}
	var stack []byte
	char := map[byte]byte {
		41 : 40,
		93 : 91,
		125 : 123,
	}
	var curChar byte

	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
		if apnd[s[i]] {
			stack = append(stack, s[i])
			curChar = s[i]
			fmt.Println(stack)
			continue
		}
		if char[s[i]] != curChar {
			return false
		}
		fmt.Println("didn't continue stack: ", stack)
		stack = stack[:len(stack)-1]
		if len(stack) > 0 {
			curChar = stack[len(stack)-1]
		} else {curChar = 0}
	}
	return true
}

func main() {
	// s := "({})"
	s := "(){}}{"
	result := isValid(s)
	fmt.Println(result)
}