package main

import (
	"aoc/common"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result1 := solve(&input, false)
	result2 := solve(&input, true)

	fmt.Printf("Result1: %d\n", result1)
	fmt.Printf("Result2: %d\n", result2)
}

const SIZE = 1000

func solve(input *string, part2 bool) int {
	lines := strings.Split(*input, "\n")

	// Parse map
	var grid [SIZE][SIZE]byte
	maxY := 0

	putBlock := func(x, y int) {
		if y > maxY {
			maxY = y
		}

		grid[y][x] = '#'
	}

	parseCoord := func(s string) (int, int) {
		split := strings.Split(s, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		return x, y
	}

	for _, line := range lines {
		if line == "" {
			continue
		}

		split := strings.Split(line, " -> ")

		for i := 1; i < len(split); i++ {
			x1, y1 := parseCoord(split[i-1])
			x2, y2 := parseCoord(split[i])

			if x1 == x2 {
				min, max := minmax(y1, y2)

				for j := min; j <= max; j++ {
					putBlock(x1, j)
				}
			} else {
				min, max := minmax(x1, x2)

				for j := min; j <= max; j++ {
					putBlock(j, y1)
				}
			}
		}
	}

	if part2 {
		maxY += 2

		for x := 0; x < SIZE; x++ {
			putBlock(x, maxY)
		}
	}

	// Simulate sand coming down from 500,0
	var nAtRest int
	curX := 500
	curY := 0
	for {
		// Flowing into the abyss below?
		if curY > maxY {
			break
		}
		// Can go down?
		if grid[curY+1][curX] == 0 {
			curY++
			continue
		}
		// Can go down and to the left?
		if grid[curY+1][curX-1] == 0 {
			curY++
			curX--
			continue
		}
		// Can go down and to the right?
		if grid[curY+1][curX+1] == 0 {
			curY++
			curX++
			continue
		}
		// Nowhere to go, sand comes to rest
		grid[curY][curX] = 'o'
		nAtRest++

		// Source is blocked?
		if curX == 500 && curY == 0 {
			break
		}

		// Reset and start over
		curX = 500
		curY = 0
	}

	return nAtRest
}

func minmax(a, b int) (int, int) {
	min := int(math.Min(float64(a), float64(b)))
	max := int(math.Max(float64(a), float64(b)))
	return min, max
}
