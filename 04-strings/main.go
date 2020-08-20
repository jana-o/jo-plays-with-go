package main

import "fmt"

func main() {
	s := "otto"
	fmt.Println(Palindrome(s))
}

func Palindrome(s string) bool {
	length := len(s)

	for i, j := 0, length-1; i < length/2; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
