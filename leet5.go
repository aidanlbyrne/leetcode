package main 

func longestPalindrome(s string) string {
	pal := string(s[0])
	var temp string;
	for i := 1; i < len(s); i++ {
		for j := 0; (i+j) < len(s) && (i-j) >= 0; j++ {
			if s[i+j] == s[i-j] {
				temp = append(temp, )
			}
		}
	}

	return pal
}

func main(){

	s := "abba"
	res := longestPalindrome(s)
	fmt.Println(res)
}