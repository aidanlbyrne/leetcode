package main

import "fmt"

func winnerOfGame(colors string) bool {
	AStarting := 0
	// BStarting := 0

	for i, v := range(colors) {
		fmt.Println(colors[i])
		if v == 65 {
			AStarting = i
			fmt.Println(AStarting)
		}
		if v == 66 {
			AStarting = -1
			fmt.Println(AStarting)
		}
	}
	return true
}

func main() {
	colors := "ABBBBBAAA"
	res := winnerOfGame(colors)
	fmt.Println(res)
}