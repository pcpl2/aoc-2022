package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var roundNumber = 1
	var totalPointsPart1 = 0
	var totalPointsPart2 = 0
	for scanner.Scan() {
		round := strings.Split(scanner.Text(), " ")
		im := round[1]
		opponent := round[0]
		pointsPart1 := calculateResultPart1(opponent, im)
		pointsPart2 := calculateResultPart2(opponent, im)
		log.Printf("{Part1}[Round %d] Im: %s | Opponent: %s | Points: %d", roundNumber, im, opponent, pointsPart1)
		log.Printf("{Part2}[Round %d] Result: %s | Opponent: %s | Points: %d", roundNumber, im, opponent, pointsPart2)
		roundNumber = roundNumber + 1
		totalPointsPart1 = totalPointsPart1 + pointsPart1
		totalPointsPart2 = totalPointsPart2 + pointsPart2
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("{part1}Total points: %d", totalPointsPart1)
	log.Printf("{part2}Total points: %d", totalPointsPart2)
}

func calculateResultPart2(opponent, roundEnd string) int {
	switch {
	case roundEnd == "X":
		switch {
		case opponent == "A":
			return 0 + getPointsByShapePart2("C")
		case opponent == "B":
			return 0 + getPointsByShapePart2("A")
		case opponent == "C":
			return 0 + getPointsByShapePart2("B")
		}
	case roundEnd == "Y":
		return 3 + getPointsByShapePart2(opponent)
	case roundEnd == "Z":
		switch {
		case opponent == "A":
			return 6 + getPointsByShapePart2("B")
		case opponent == "B":
			return 6 + getPointsByShapePart2("C")
		case opponent == "C":
			return 6 + getPointsByShapePart2("A")
		}
	}
	return 0
}

func calculateResultPart1(opponent, im string) int {
	shapePoints := getPointsByShapePart1(im)
	switch {
	case opponent == "A":
		switch {
		case im == "X":
			return 3 + shapePoints
		case im == "Y":
			return 6 + shapePoints
		case im == "Z":
			return 0 + shapePoints
		}
	case opponent == "B":
		switch {
		case im == "X":
			return 0 + shapePoints
		case im == "Y":
			return 3 + shapePoints
		case im == "Z":
			return 6 + shapePoints
		}
	case opponent == "C":
		switch {
		case im == "X":
			return 6 + shapePoints
		case im == "Y":
			return 0 + shapePoints
		case im == "Z":
			return 3 + shapePoints
		}

	}
	return shapePoints
}

func getPointsByShapePart1(shape string) int {
	switch {
	case shape == "X":
		return 1
	case shape == "Y":
		return 2
	case shape == "Z":
		return 3
	}
	return 0
}

func getPointsByShapePart2(shape string) int {
	switch {
	case shape == "A":
		return 1
	case shape == "B":
		return 2
	case shape == "C":
		return 3
	}
	return 0
}
