package main

import "fmt"


func setZeroes(matrix [][]int) {
	row := len(matrix)
	col := len(matrix[0])
	colZ := make([]int, col)
	for i := 0; i < row; i++ {
		for j := 0, j < col; j++ {
			if matrix[i][j] == 0 {
				
				
			}
		}
	}
}

func main() {
	matrix := [][]int{{1,1,1}, {0,1,0}, {1,1,1}}
	fmt.Println(len(matrix))
	// setZeroes(matrix)
}