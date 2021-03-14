package main

import (
	"fmt"
)

//reverse integer
//constraints from leetcode:
// If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
//0ms 100%faster and 2.2MB less than 100% of Go submissions
func reverseInt(num int) int {
	var r int

	//if num negative -> convert to positve first
	if num < 0 {
		num = num * -1
		for num > 0 {
			//pop: %10 and x /=10
			r = r*10 + num%10
			num = num / 10
		}
		//if r is out of range
		if r > 2147483647 || r < -2147483647 {
			return 0
		}
		return r * -1
	}
	for num > 0 {
		r = r*10 + num%10
		num = num / 10
	}

	//if r is out of range
	if r > 2147483647 || r < -2147483647 {
		return 0
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

//1480: Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
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

// 1431: Given the array candies and the integer extraCandies, where candies[i] represents the number of candies that the ith kid has.
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

// version2: zero value of bool is false
// 0ms, faster than 100%; 2.3MB
func kidsWithGreatesNrOfCandies2(candies []int, extraCandies int) []bool {
	greatest := 0
	result := make([]bool, len(candies))

	for _, v := range candies {
		if v > greatest {
			greatest = v
		}
	}
	for i, v := range candies {
		if v+extraCandies >= greatest {
			result[i] = true
		}
	}
	return result
}

// 1470: Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].
// Return the array in the form [x1,y1,x2,y2,...,xn,yn].
// 4ms, faster than 94.56%, 3.7MB, less than 98.58% of Go submissions
func shuffleArr(nums []int, n int) []int {
	shuffled := make([]int, len(nums))
	j := 0
	for i := 0; i < n; i++ {
		shuffled[j] = nums[i]
		j++
		shuffled[j] = nums[i+n]
		j++
	}
	return shuffled
}

// 1512: Given an array of integers nums. A pair (i,j) is called good if nums[i] == nums[j] and i < j.
// Return the number of good pairs.
// use hashmap here to find pairs: 0ms, 2MB
func numIdenticalPairs(nums []int) int {
	var pairs int
	hm := make(map[int]int)

	//find count of number occurences of number then store in map
	for _, n := range nums {
		pairs += hm[n]
		hm[n]++
		fmt.Println("pairs: ", pairs, "hm:", hm) //final hm in ex.: 1:3, 2:1, 3:2
	}
	return pairs
}

//remove duplicates in slice
func removeDuplicates(a []int) []int {
	result := []int{}
	seen := map[int]int{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}

// 1512: same as above using two for loops
func numIdenticalPairsLoop(nums []int) int {
	var pairs int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				pairs++
			}
		}
	}
	return pairs
}

func main() {
	// fmt.Println(reverseInt(1534236469))
	// fmt.Println(isPalindrome(11))
	// fmt.Println(isPalindrome(110))
	// fmt.Println(runningSumOf1dArr([]int{1, 2, 3, 4}))
	// fmt.Println(kidsWithGreatesNrOfCandies([]int{2, 3, 5, 1, 3}, 3))
	// fmt.Println(kidsWithGreatesNrOfCandies2([]int{2, 3, 5, 1, 3}, 3))
	// fmt.Println(shuffleArr([]int{2, 5, 1, 3, 4, 7}, 3))
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
	// fmt.Println(numIdenticalPairsLoop([]int{1, 2, 3, 1, 1, 3}))
}
