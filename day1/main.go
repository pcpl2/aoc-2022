package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Elf struct {
	TotalCalories int
	ElfNumber     int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var elfs []Elf

	//Init list
	elfs = make([]Elf, 0)
	elfs = append(elfs, Elf{TotalCalories: 0, ElfNumber: 1})

	var currentElf = len(elfs) - 1
	for scanner.Scan() {
		if scanner.Text() != "" {
			log.Println(scanner.Text())
			value, err := strconv.Atoi(scanner.Text())
			if err == nil {
				elfs[currentElf].TotalCalories = elfs[currentElf].TotalCalories + value
			}
		} else {
			log.Println("New elf")
			elfs = append(elfs, Elf{TotalCalories: 0, ElfNumber: currentElf + 1})
			currentElf = len(elfs) - 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	elfs = sortElfs(elfs)
	log.Printf("Total number of calories carried by the elf having the most calories: %d", elfs[0].TotalCalories)
	totalOfTopThree := elfs[0].TotalCalories + elfs[1].TotalCalories + elfs[2].TotalCalories
	log.Printf("Total calories carried by the 3 elves with the most calories: %d", totalOfTopThree)
}

func sortElfs(elfs []Elf) []Elf {
	for i := 0; i < len(elfs)-1; i++ {
		for j := 0; j < len(elfs)-i-1; j++ {
			if elfs[j].TotalCalories < elfs[j+1].TotalCalories {
				elfs[j], elfs[j+1] = elfs[j+1], elfs[j]
			}
		}
	}
	return elfs
}
