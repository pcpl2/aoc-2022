package main

import "testing"

var td01 = ElfPair{
	ElfFirst: ElfData{
		Start: 2,
		End:   4,
		Range: []int{2, 3, 4},
	},
	ElfSecond: ElfData{
		Start: 6,
		End:   8,
		Range: []int{6, 7, 8},
	},
}
var td02 = ElfPair{
	ElfFirst: ElfData{
		Start: 2,
		End:   3,
		Range: []int{2, 3},
	},
	ElfSecond: ElfData{
		Start: 4,
		End:   5,
		Range: []int{4, 5},
	},
}
var td03 = ElfPair{
	ElfFirst: ElfData{
		Start: 5,
		End:   7,
		Range: []int{5, 6, 7},
	},
	ElfSecond: ElfData{
		Start: 7,
		End:   9,
		Range: []int{7, 8, 9},
	},
}
var td04 = ElfPair{
	ElfFirst: ElfData{
		Start: 2,
		End:   8,
		Range: []int{2, 3, 4, 5, 6, 7, 8},
	},
	ElfSecond: ElfData{
		Start: 3,
		End:   7,
		Range: []int{3, 4, 5, 6, 7},
	},
}
var td05 = ElfPair{
	ElfFirst: ElfData{
		Start: 6,
		End:   6,
		Range: []int{6},
	},
	ElfSecond: ElfData{
		Start: 4,
		End:   6,
		Range: []int{4, 5, 6},
	},
}
var td06 = ElfPair{
	ElfFirst: ElfData{
		Start: 2,
		End:   6,
		Range: []int{2, 3, 4, 5, 6},
	},
	ElfSecond: ElfData{
		Start: 4,
		End:   8,
		Range: []int{4, 5, 6, 7, 8},
	},
}

func TestResultPart1(t *testing.T) {
	if checkOneReangeFyllyContainOther(td01) {
		t.FailNow()
	}

	if checkOneReangeFyllyContainOther(td02) {
		t.FailNow()
	}

	if checkOneReangeFyllyContainOther(td03) {
		t.FailNow()
	}

	if !checkOneReangeFyllyContainOther(td04) {
		t.FailNow()
	}

	if !checkOneReangeFyllyContainOther(td05) {
		t.FailNow()
	}

	if checkOneReangeFyllyContainOther(td06) {
		t.FailNow()
	}
}

func TestResultPart2(t *testing.T) {
	if checkPairRangeOverlap(td01) {
		t.FailNow()
	}

	if checkPairRangeOverlap(td02) {
		t.FailNow()
	}

	if !checkPairRangeOverlap(td03) {
		t.FailNow()
	}

	if !checkPairRangeOverlap(td04) {
		t.FailNow()
	}

	if !checkPairRangeOverlap(td05) {
		t.FailNow()
	}

	if !checkPairRangeOverlap(td06) {
		t.FailNow()
	}
}
