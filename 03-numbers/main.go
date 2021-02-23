package main

import (
	"fmt"
	"strings"
)

//reverse string
func reverseStr(str string) string {
	words := strings.Fields(str)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

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

//Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
//Return the running sum of nums.
func runningSumOf1dArr(nums []int) []int {
	result := []int{}
	count := 0
	for _, v := range nums {
		count += v
		result = append(result, count)
	}
	return result
}

func main() {
	// fmt.Println(reverseInt(123))
	// fmt.Println(isPalindrome(11))
	// fmt.Println(isPalindrome(110))
	fmt.Println(runningSumOf1dArr([]int{1, 2, 3, 4}))
}
