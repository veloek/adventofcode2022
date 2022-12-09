package main

import (
	"aoc/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result1, result2 := solve(&input)

	fmt.Printf("Result1: %d\n", result1)
	fmt.Printf("Result2: %d\n", result2)
}

func solve(input *string) (int, int) {
	lines := strings.Split(*input, "\n")

	head := pos{x: 0, y: 0}
	tails := [9]pos{}
	for i := 0; i < len(tails); i++ {
		tails[i] = pos{x: 0, y: 0}
	}

	visitedFirst := make(map[pos]bool)
	visitedFirst[tails[0]] = true

	visitedLast := make(map[pos]bool)
	visitedLast[tails[len(tails)-1]] = true

	for _, line := range lines {
		if line == "" {
			continue
		}

		split := strings.Split(line, " ")
		steps, _ := strconv.Atoi(split[1])

		for i := 0; i < steps; i++ {
			switch split[0] {
			case "R":
				head.x--
				break
			case "L":
				head.x++
				break
			case "U":
				head.y--
				break
			case "D":
				head.y++
				break
			}

			follow(&tails[0], head)
			for j := 1; j < len(tails); j++ {
				follow(&tails[j], tails[j-1])
			}

			visitedFirst[tails[0]] = true
			visitedLast[tails[len(tails)-1]] = true

		}
	}

	return len(visitedFirst), len(visitedLast)
}

func follow(tail *pos, head pos) {
	if tail.x < head.x-1 || tail.x > head.x+1 || tail.y < head.y-1 || tail.y > head.y+1 {
		// Move tail
		if tail.y < head.y {
			tail.y++
		} else if tail.y > head.y {
			tail.y--
		}

		if tail.x < head.x {
			tail.x++
		} else if tail.x > head.x {
			tail.x--
		}
	}
}

type pos struct {
	x, y int
}
