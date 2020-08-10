package reverse

import "testing"

func TestReverse(t *testing.T) {
	ans := Reverse(12)
	if ans != 21 {
		t.Errorf("Reverse(12) = %d; want 21", ans)
	}
}
