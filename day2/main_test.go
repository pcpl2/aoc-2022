package main

import "testing"

func TestCalculateResultPart1(t *testing.T) {
	result := calculateResultPart1("A", "Y")

	if result != 8 {
		t.FailNow()
	}

	result = calculateResultPart1("B", "X")

	if result != 1 {
		t.FailNow()
	}

	result = calculateResultPart1("C", "Z")
	if result != 6 {
		t.FailNow()
	}

	result = calculateResultPart2("A", "Y")

	if result != 4 {
		t.FailNow()
	}

	result = calculateResultPart2("B", "X")

	if result != 1 {
		t.FailNow()
	}

	result = calculateResultPart2("C", "Z")
	if result != 7 {
		t.FailNow()
	}
}
