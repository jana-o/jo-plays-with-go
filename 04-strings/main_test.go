package main

import "testing"

func TestIsPalindrom(t *testing.T) {
	s := "otto"
	ans := isPalindrome(s)
	if !ans == true {
		t.Errorf("IsPalindrom() %t; want true", ans)
	}
}
