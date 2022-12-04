package main

import "testing"

func TestCalculateResult(t *testing.T) {
	result := calculateResult("A", "Y")

	if result != 8 {
		t.FailNow()
	}

	result = calculateResult("B", "X")

	if result != 1 {
		t.FailNow()
	}

	result = calculateResult("C", "Z")
	if result != 6 {
		t.FailNow()
	}
}
