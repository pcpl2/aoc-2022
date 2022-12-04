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
	var totalPoints = 0
	for scanner.Scan() {
		round := strings.Split(scanner.Text(), " ")
		im := round[1]
		opponent := round[0]
		points := calculateResult(opponent, im)
		log.Printf("[Round %d] Im: %s | Opponent: %s | Points: %d", roundNumber, im, opponent, points)
		roundNumber = roundNumber + 1
		totalPoints = totalPoints + points
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Total points: %d", totalPoints)
}

func calculateResult(opponent, im string) int {
	shapePoints := getPointsByShape(im)
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

func getPointsByShape(shape string) int {
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
