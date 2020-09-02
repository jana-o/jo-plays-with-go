package main

import "testing"

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
