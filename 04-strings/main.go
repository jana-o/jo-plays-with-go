package main

import (
	"fmt"
	"strings"
)

//reverse string using []rune
func reversed(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

//reverse a string
func reverse(s string) (reversed string) {
	for _, v := range s {
		reversed = string(v) + reversed
	}
	return
}

//determine if given string is a palindrom
func isPalindrome(input string) bool {
	return strings.ToUpper(input) == reversed(strings.ToUpper(input))
}

func main() {
	fmt.Println(isPalindrome("Madam"))
	fmt.Println(isPalindrome("not a palindrome"))
}
