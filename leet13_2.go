package main

import (
	"fmt"
)

func romanToInt(s string) int {
    alph := map[string]int {
        "I" : 1,
		"IV" : 4,
        "V" : 5,
		"IX" : 9,
        "X" : 10,
		"XL" : 40,
        "L" : 50,
		"XC" : 90,
        "C" : 100,
		"CD" : 400,
        "D" : 500,
		"CM" : 900,
        "M" : 1000,
    }
	total := 0
	prevV := 0
	for i := len(s) - 1; i >= 0; i-- {
		if alph[string(s[i])] < prevV {
			total -= alph[string(s[i])]
			prevV = alph[string(s[i])]
			continue
		}
		total += alph[string(s[i])]
		prevV = alph[string(s[i])]
	}
	return total	
}


func main() {
// input := "III"
// input := "LVIII"
input := "MCMXCIV"
// input := "IV"
result := romanToInt(input)

fmt.Println(result)

}