package main

import (
	"fmt"
)

func isValid(s string) bool {
	char := map[byte]byte {
		41 : 40,
		93 : 91,
		125 : 123,
	}

	var stack []byte

	for _, b := range s {
		if _, ok := char[b]; ok {
			stack = append(stack, b)
		} else if len(stack) == 0 || char[stack[len(stack)-1]] != b {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}