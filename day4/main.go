package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ElfPair struct {
	ElfFirst  []int
	ElfSecond []int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var totalPairs = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := parsePairData(scanner.Text())
		pd := checkOneReangeFyllyContainOther(data)
		if pd {
			totalPairs = totalPairs + 1
		}
		log.Printf("Data :%v | %v", data, pd)
	}

	log.Printf("Pairs have range fully contain the orher: %d", totalPairs)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parsePairData(input string) ElfPair {
	elfs := strings.Split(input, ",")
	firstElf := make([]int, 0)
	secondElf := make([]int, 0)

	for elfi, elf := range elfs {
		elfRange := strings.Split(elf, "-")
		start, _ := strconv.Atoi(elfRange[0])
		end, _ := strconv.Atoi(elfRange[1])

		for i := start; i <= end; i++ {
			if elfi == 0 {
				firstElf = append(firstElf, i)
			} else if elfi == 1 {
				secondElf = append(secondElf, i)
			}
		}
	}

	return ElfPair{
		ElfFirst:  firstElf,
		ElfSecond: secondElf,
	}
}

func checkOneReangeFyllyContainOther(pair ElfPair) bool {
	totalRange := append(pair.ElfFirst, pair.ElfSecond...)
	totalRange = removeDuplacated(totalRange)

	sort.SliceStable(totalRange, func(i, j int) bool {
		return i < j
	})

	if len(totalRange) == len(pair.ElfFirst) || len(totalRange) == len(pair.ElfSecond) {
		return true
	}
	return false
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
