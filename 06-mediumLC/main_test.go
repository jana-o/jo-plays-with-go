package main

import (
	"reflect"
	"testing"
)

func TestMinOperations(t *testing.T) {
	var tests = []struct {
		input string
		want  []int
	}{
		{"110", []int{1, 1, 3}},
		{"001011", []int{11, 8, 5, 4, 3, 4}},
	}

	for _, tt := range tests {
		got := minOperations(tt.input)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}

}
