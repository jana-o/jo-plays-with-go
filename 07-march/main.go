package main

import "fmt"

func main() {
	fmt.Println(distributedCandies([]int{1, 1, 2, 3}))
}

// Alice has n candies, candyType[i]; only able to eat n/2 candies
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
