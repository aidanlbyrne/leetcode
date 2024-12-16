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
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			total += alph[string(s[i])]
			return total
		}
		comb := string(s[i]) + string(s[i+1])

		if _, ok := alph[comb]; ok {
			total += alph[comb]
			i++
			continue
		}
		total += alph[string(s[i])]
	}
	return total
}


func main() {
// input := "III"
// input := "LVIII"
input := "MCMXCIV"
result := romanToInt(input)

fmt.Println(result)

}