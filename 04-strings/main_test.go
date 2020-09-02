package main

import "testing"

func TestPalindrom(t *testing.T) {
	s := "otto"
	ans := Palindrome(s)
	if !ans == true {
		t.Errorf("Reverse() %t; want true", ans)
	}
}
