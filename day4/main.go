package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ElfData struct {
	Start int
	End   int
	Range []int
}

type ElfPair struct {
	ElfFirst  ElfData
	ElfSecond ElfData
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var totalPairsPart1 = 0
	var totalPairsPart2 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := parsePairData(scanner.Text())
		pd := checkOneReangeFyllyContainOther(data)
		if pd {
			totalPairsPart1 = totalPairsPart1 + 1
		}
		pd2 := checkPairRangeOverlap(data)
		if pd2 {
			totalPairsPart2 = totalPairsPart2 + 1
		}
	}

	log.Printf("Pairs have range fully contain the other: %d", totalPairsPart1)
	log.Printf("Pairs have range overlap: %d", totalPairsPart2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parsePairData(input string) ElfPair {
	elfs := strings.Split(input, ",")
	firstElf := ElfData{
		Range: make([]int, 0),
	}
	secondElf := ElfData{
		Range: make([]int, 0),
	}

	for elfi, elf := range elfs {
		elfRange := strings.Split(elf, "-")
		start, _ := strconv.Atoi(elfRange[0])
		end, _ := strconv.Atoi(elfRange[1])
		if elfi == 0 {
			firstElf.Start = start
			firstElf.End = end
		} else if elfi == 1 {
			secondElf.Start = start
			secondElf.End = end
		}

		for i := start; i <= end; i++ {
			if elfi == 0 {
				firstElf.Range = append(firstElf.Range, i)
			} else if elfi == 1 {
				secondElf.Range = append(secondElf.Range, i)
			}
		}
	}

	return ElfPair{
		ElfFirst:  firstElf,
		ElfSecond: secondElf,
	}
}

func checkOneReangeFyllyContainOther(pair ElfPair) bool {
	totalRange := append(pair.ElfFirst.Range, pair.ElfSecond.Range...)
	totalRange = removeDuplacated(totalRange)

	sort.SliceStable(totalRange, func(i, j int) bool {
		return i < j
	})

	if len(totalRange) == len(pair.ElfFirst.Range) || len(totalRange) == len(pair.ElfSecond.Range) {
		return true
	}
	return false
}

func checkPairRangeOverlap(pair ElfPair) bool {
	log.Printf("%v | %d | %d", pair, pair.ElfSecond.Start, pair.ElfFirst.End)
	return !(pair.ElfFirst.End < pair.ElfSecond.Start || pair.ElfFirst.Start > pair.ElfSecond.End)
}

func removeDuplacated(elfTotalRange []int) []int {
	allData := make(map[int]bool)
	list := make([]int, 0)
	for _, item := range elfTotalRange {
		if _, value := allData[item]; !value {
			allData[item] = true
			list = append(list, item)
		}
	}
	return list
}
