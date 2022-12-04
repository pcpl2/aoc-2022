package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Elf struct {
	TotalCalories int
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
	elfs = append(elfs, Elf{TotalCalories: 0})

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
			elfs = append(elfs, Elf{TotalCalories: 0})
			currentElf = len(elfs) - 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	elfs = sortElfs(elfs)
	log.Printf("Elf carrying the most Calories: %d", elfs[0].TotalCalories)
}

func sortElfs(elfs []Elf) []Elf {
	for i := 0; i < len(elfs)-1; i++ {
		for j := 0; j < len(elfs)-j-1; j++ {
			if elfs[j].TotalCalories < elfs[j+1].TotalCalories {
				elfs[j], elfs[j+1] = elfs[j+1], elfs[j]
			}
		}
	}
	return elfs
}
