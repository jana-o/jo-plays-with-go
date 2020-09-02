package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	result := IsPalindrome(12321)
	if result != true {
		t.Errorf("got %t, want true", result)
	}
}
