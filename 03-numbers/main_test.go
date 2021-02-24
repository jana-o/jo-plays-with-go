package main

import (
	"reflect"
	"testing"
)

func TestReverseInt(t *testing.T) {
	ans := reverseInt(12)
	if ans != 21 {
		t.Errorf("Reverse(12) = %d; want 21", ans)
	}
}

func TestIsPalindrome(t *testing.T) {
	result := isPalindrome(12321)
	if result != true {
		t.Errorf("got %t, want true", result)
	}
}

func TestRunningSumOf1dArr(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
	}
	for _, tt := range tests {
		//use reflect.DeepEqual because Go does not allow equality operator with slices bc ref type
		got := runningSumOf1dArr(tt.input)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}

func TestKidsWithGreatesNrOfCandies(t *testing.T) {
	type input struct {
		candies []int
		extra   int
	}
	var tests = []struct {
		input input
		want  []bool
	}{
		{input{[]int{2, 3, 5, 1, 3}, 3}, []bool{true, true, true, false, true}},
		{input{[]int{4, 2, 1, 1, 2}, 1}, []bool{true, false, false, false, false}},
		{input{[]int{12, 1, 12}, 10}, []bool{true, false, true}},
	}
	for _, tt := range tests {
		got := kidsWithGreatesNrOfCandies(tt.input.candies, tt.input.extra)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}

func TestKidsWithGreatesNrOfCandies2(t *testing.T) {
	type input struct {
		candies []int
		extra   int
	}
	var tests = []struct {
		input input
		want  []bool
	}{
		{input{[]int{2, 3, 5, 1, 3}, 3}, []bool{true, true, true, false, true}},
		{input{[]int{4, 2, 1, 1, 2}, 1}, []bool{true, false, false, false, false}},
		{input{[]int{12, 1, 12}, 10}, []bool{true, false, true}},
	}
	for _, tt := range tests {
		got := kidsWithGreatesNrOfCandies2(tt.input.candies, tt.input.extra)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}

func BenchmarkKidsWithGreatesNrOfCandies(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kidsWithGreatesNrOfCandies([]int{2, 3, 5, 1, 3}, 3)
	}
}
func BenchmarkKidsWithGreatesNrOfCandies2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kidsWithGreatesNrOfCandies2([]int{2, 3, 5, 1, 3}, 3)
	}
}

func TestShuffleArr(t *testing.T) {
	type input struct {
		nums []int
		n    int
	}
	var tests = []struct {
		input input
		want  []int
	}{
		{input{[]int{2, 5, 1, 3, 4, 7}, 3}, []int{2, 3, 5, 4, 1, 7}},
		{input{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4}, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{input{[]int{1, 1, 2, 2}, 2}, []int{1, 2, 1, 2}},
	}
	for _, tt := range tests {
		got := shuffleArr(tt.input.nums, tt.input.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v, want%v", got, tt.want)
		}
	}
}
