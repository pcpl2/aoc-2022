package main

import "testing"

func TestResultPart1(t *testing.T) {
	result := checkOneReangeFyllyContainOther(ElfPair{
		ElfFirst:  []int{2, 3, 4},
		ElfSecond: []int{6, 7, 8},
	})

	if result {
		t.FailNow()
	}

	result = checkOneReangeFyllyContainOther(ElfPair{
		ElfFirst:  []int{2, 3},
		ElfSecond: []int{4, 5},
	})

	if result {
		t.FailNow()
	}

	result = checkOneReangeFyllyContainOther(ElfPair{
		ElfFirst:  []int{5, 6, 7},
		ElfSecond: []int{7, 8, 9},
	})

	if result {
		t.FailNow()
	}

	result = checkOneReangeFyllyContainOther(ElfPair{
		ElfFirst:  []int{2, 3, 4, 5, 6, 7, 8},
		ElfSecond: []int{3, 4, 5, 6, 7},
	})

	if !result {
		t.FailNow()
	}

	result = checkOneReangeFyllyContainOther(ElfPair{
		ElfFirst:  []int{6},
		ElfSecond: []int{4, 5, 6},
	})

	if !result {
		t.FailNow()
	}

	result = checkOneReangeFyllyContainOther(ElfPair{
		ElfFirst:  []int{2, 3, 4, 5, 6},
		ElfSecond: []int{4, 5, 6, 7, 8},
	})

	if result {
		t.FailNow()
	}
}
