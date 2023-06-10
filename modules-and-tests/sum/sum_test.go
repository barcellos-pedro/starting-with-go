package sum

import "testing"

func TestSum(t *testing.T) {
	expected := 4
	result := Add(2, 2)
	if result != expected {
		t.Errorf("Sum = %d; expected %d", result, expected)
	}
}
