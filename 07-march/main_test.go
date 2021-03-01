package main

import "testing"

func TestDistributedCandies(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1, 1, 2, 2, 3, 3}, 3},
		{[]int{1, 1, 2, 3}, 2},
		{[]int{6, 6, 6, 6}, 1},
	}

	for _, tt := range tests {
		got := distributedCandies(tt.input)

		if got != tt.want {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}

}
