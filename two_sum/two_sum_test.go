package two_sum

import "testing"

func TestExecute(t *testing.T) {
	numbers := []int{1, 2, 3}
	target := 4
	expected := [2]int{0, 2}

	result := Execute(numbers, target)
	if result != expected {
		t.Fatalf("Failed! Expected: %v. Actual: %v", expected, result)
	}
}
