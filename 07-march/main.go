package main

import (
	"fmt"
)

func main() {
	// fmt.Println(distributedCandies([]int{1, 1, 2, 3}))
	fmt.Println(findErrorNums([]int{1, 1}))
}

// Day 01: Alice has n candies, candyType[i]; only able to eat n/2 candies
// find maximum number of different types of candies she can eat if she only eats n/2 of them
// n is always even
// runtime: 176 ms, beats 54.55%, memory usage 7.2MB, beats 29.55%
func distributedCandies(candyType []int) int {

	max := len(candyType) / 2
	types := make(map[int]int)

	for i := 0; i < len(candyType); i++ {
		types[candyType[i]] = candyType[i]
	}

	if max < len(types) {
		return max
	}
	return len(types)

}

//version 2 with range instead of for loop
func distributedCandies2(candyType []int) int {

	max := len(candyType) / 2
	types := make(map[int]int)

	for _, c := range candyType {
		types[c]++
	}

	if max < len(types) {
		return max
	}
	return len(types)

}

// Day 02: one of the numbers in set of integers got duplicated to another in the set, which results in repetition of one number and loss of another
// find the number that occurs twice and the number that is missing and return them in form of an array
func findErrorNums(nums []int) []int {

	result := []int{}
	seen := map[int]int{}

	for _, val := range nums {
		seen[val]++

		if t, ok := seen[val]; ok {
			if t == 2 {
				if _, ok := seen[val-1]; !ok {
					result = append(result, val, val-1)
				} else {
					result = append(result, val, val+1)
				}
			}

		}
	}
	fmt.Println(seen)
	return result
}
