package main

import (
	"fmt"
)

func main() {
	fmt.Println(minOperations("110"))
}

//1769: Minimum Number of Operations to Move All Balls to Each Box
// You have n boxes. You are given a binary string boxes of length n, where boxes[i] is '0' if the ith box is empty, and '1' if it contains one ball.
// 388ms 100% faster, 4.9MB 100% less memory usage
func minOperations(boxes string) []int {

	//use array here because length of result is equal to length or boxes
	result := make([]int, len(boxes))
	var counter int

	for i := range boxes {
		for j, v := range boxes {
			if i != j {
				//string conversion here because input is of type string
				if string(v) == "1" {

					if j < i {
						counter += i - j
					}

					if j > i {
						counter += j - i
					}

				}
			}
		}
		result[i] = counter
		counter = 0
	}

	return result
}
