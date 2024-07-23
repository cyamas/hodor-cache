package client

import "testing"

func TestValToIntArray(t *testing.T) {
	input := "(31, 30, 4, 0)"
	expected := []int64{31, 30, 4, 0}
	arr := valToIntArray(input)
	if len(arr) != 4 {
		t.Fatalf("Len should be 4. got %d", len(arr))
	}
	for i, item := range arr {
		if item != expected[i] {
			t.Fatalf("Item should be %d. got %v", expected[i], item)
		}
	}
}
