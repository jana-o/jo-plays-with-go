package main

import "fmt"

func main() {
	fmt.Println(minOperations("min"))
}

//1769: Minimum Number of Operations to Move All Balls to Each Box
// You have n boxes. You are given a binary string boxes of length n, where boxes[i] is '0' if the ith box is empty, and '1' if it contains one ball.
func minOperations(boxes string) []int {
	res := []int{1, 2}
	return res
}
