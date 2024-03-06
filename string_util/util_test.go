package string_util

import "testing"

func TestString(t *testing.T) {
	result := SplitString("-30k tomatoes")
	expected := "-30000"

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}