package main

import (
	"fmt"
)

//reverse integer
func reverseInt(num int) int {
	var r int
	for num > 0 {
		//pop: %10 and x /=10
		r = r*10 + num%10
		num = num / 10
	}
	return r
}

//determine if palindromic number
func isPalindrome(num int) bool {
	if num < 0 {
		return false
	}
	return num == reverseInt(num)
}

func main() {
	fmt.Println(reverseInt(123))
	fmt.Println(isPalindrome(11))
	fmt.Println(isPalindrome(110))
}
