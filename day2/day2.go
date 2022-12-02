package main

import (
	"aoc/common"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	score := calculateScore(&input, func(opponent, shape string) int {
		//return shapeScore(shape) + outcomeScore(opponent, shape) // Part 1
		return outcomeScore2(opponent, shape) // Part 2
	})

	fmt.Printf("Score: %d\n", score)
}

func calculateScore(input *string, outcomeScore func(string, string) int) int {
	lines := strings.Split(*input, "\n")
	score := 0

	for _, round := range lines {
		if round == "" {
			continue
		}
		s := strings.Split(round, " ")
		score += outcomeScore(s[0], s[1])
	}

	return score
}

func shapeScore(shape string) int {
	switch shape {
	case "X": // Rock
		return 1
	case "Y": // Paper
		return 2
	case "Z": // Scissors
		return 3
	default:
		return 0
	}
}

func outcomeScore(opponent string, shape string) int {
	switch true {
	case opponent == "A" && shape == "X": // Rock vs. rock
		return 3
	case opponent == "A" && shape == "Y": // Rock vs. paper
		return 6
	case opponent == "A" && shape == "Z": // Rock vs. scissors
		return 0
	case opponent == "B" && shape == "X": // Paper vs. rock
		return 0
	case opponent == "B" && shape == "Y": // Paper vs. paper
		return 3
	case opponent == "B" && shape == "Z": // Paper vs. scissors
		return 6
	case opponent == "C" && shape == "X": // Scissors vs. rock
		return 6
	case opponent == "C" && shape == "Y": // Scissors vs. paper
		return 0
	case opponent == "C" && shape == "Z": // Scissors vs. scissors
		return 3
	default:
		return 0
	}
}

func outcomeScore2(opponent string, shape string) int {
	switch true {
	case opponent == "A" && shape == "X": // Loose: Rock vs. scissors
		return 3
	case opponent == "A" && shape == "Y": // Draw:  Rock vs. rock
		return 1 + 3
	case opponent == "A" && shape == "Z": // Win:   Rock vs. paper
		return 2 + 6
	case opponent == "B" && shape == "X": // Loose: Paper vs. rock
		return 1
	case opponent == "B" && shape == "Y": // Draw:  Paper vs. paper
		return 2 + 3
	case opponent == "B" && shape == "Z": // Win:   Paper vs. scissors
		return 3 + 6
	case opponent == "C" && shape == "X": // Loose: Scissors vs. paper
		return 2
	case opponent == "C" && shape == "Y": // Draw:  Scissors vs. scissors
		return 3 + 3
	case opponent == "C" && shape == "Z": // Win:   Scissors vs. rock
		return 1 + 6
	default:
		return 0
	}
}
