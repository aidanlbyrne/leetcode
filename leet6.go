package main

import "fmt"

func convert(s string, numRows int) string {
	var res string
	slice := make([]string, 3)

	i := 0
	dir := -1
	for _,c := range s {
		fmt.Println(i)
		if i == numRows -1 || i == 0 {dir *= -1}
		slice[i] += string(c)
		i += dir
	}
	for _,c := range slice {
		res += string(c)
	}
	return res
}
func main() {
	s := "PAYPALISHIRING"
	num := 3

	var res string
	res = convert(s, num)
	fmt.Println(res)
}