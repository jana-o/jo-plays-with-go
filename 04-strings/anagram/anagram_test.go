package anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	s1 := "tommarvoloriddle"
	s2 := "iamlordvoldemort"

	ans := isAnagram(s1, s2)
	if ans != true {
		t.Errorf("IsAnagram() want true, got %t", ans)
	}
}
