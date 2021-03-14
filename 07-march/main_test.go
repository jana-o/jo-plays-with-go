package main

import (
	"reflect"
	"testing"
)

func TestDistributedCandieså(t *testing.T) {
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
func TestDistributedCandies2(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1, 1, 2, 2, 3, 3}, 3},
		{[]int{1, 1, 2, 3}, 2},
		{[]int{6, 6, 6, 6}, 1},
	}

	for _, tt := range tests {
		got := distributedCandies2(tt.input)

		if got != tt.want {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}

}

func TestFindErrorNums(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 2, 4}, []int{2, 3}},
		{[]int{1, 1}, []int{1, 2}},
		{[]int{2, 2}, []int{2, 1}},
		{[]int{3,2,3,4,6,5}, []{3,1}},
	}
	for _, tt := range tests {
		got := findErrorNums(tt.input)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
