package palindrome

import (
	"github.com/jana-o/jo-plays-with-go/03-numbers/reverse"
)

//determine if given number is a palindromic number
func IsPalindrome(number int) bool {
	return number == reverse.Reverse(number)
}
