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

// Given the array candies and the integer extraCandies, where candies[i] represents the number of candies that the ith kid has.
// For each kid check if there is a way to distribute extraCandies among the kids such that he or she can have the greatest number of candies among them. Notice that multiple kids can have the greatest number of candies.
func kidsWithGreatesNrOfCandies(candies []int, extraCandies int) []bool {
	greatest := 0
	for _, v := range candies {
		if v > greatest {
			greatest = v
		}
	}

	result := []bool{}
	for _, v := range candies {
		switch {
		case v < greatest && v+extraCandies < greatest:
			result = append(result, false)
		default:
			result = append(result, true)
		}
	}
	return result
}

func main() {
	// fmt.Println(reverseInt(123))
	// fmt.Println(isPalindrome(11))
	// fmt.Println(isPalindrome(110))
	// fmt.Println(runningSumOf1dArr([]int{1, 2, 3, 4}))
	fmt.Println(kidsWithGreatesNrOfCandies([]int{2, 3, 5, 1, 3}, 3))
}
