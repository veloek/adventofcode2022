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

	sum1, sum2 := solve(&input)

	fmt.Printf("Sum1: %d\n", sum1)
	fmt.Printf("Sum2: %d\n", sum2)
}

func solve(input *string) (int, int) {
	lines := strings.Split(*input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	nCols := len(lines[0])
	nRows := len(lines)

	sum := nCols*2 + (nRows-2)*2
	maxScenicScore := 0
	for i := 1; i < nRows-1; i++ {
		for j := 1; j < nCols-1; j++ {
			if isVisible(&lines, i, j) {
				sum++
			}

			s := scenicScore(&lines, i, j)
			if s > maxScenicScore {
				maxScenicScore = s
			}
		}
	}

	return sum, maxScenicScore
}

func scenicScore(rows *[]string, i, j int) int {
	left := 0
	for k := j - 1; k >= 0; k-- {
		left++

		if (*rows)[i][k] >= (*rows)[i][j] {
			break
		}
	}

	right := 0
	for k := j + 1; k < len((*rows)[0]); k++ {
		right++

		if (*rows)[i][k] >= (*rows)[i][j] {
			break
		}
	}

	top := 0
	for k := i - 1; k >= 0; k-- {
		top++

		if (*rows)[k][j] >= (*rows)[i][j] {
			break
		}
	}

	bottom := 0
	for k := i + 1; k < len(*rows); k++ {
		bottom++

		if (*rows)[k][j] >= (*rows)[i][j] {
			break
		}
	}

	return left * right * top * bottom
}

func isVisible(rows *[]string, i, j int) bool {
	visible := true
	for k := 0; k < j; k++ {
		if (*rows)[i][k] >= (*rows)[i][j] {
			visible = false
		}
	}
	if visible {
		return true
	}

	visible = true
	for k := j + 1; k < len((*rows)[0]); k++ {
		if (*rows)[i][k] >= (*rows)[i][j] {
			visible = false
		}
	}
	if visible {
		return true
	}

	visible = true
	for k := 0; k < i; k++ {
		if (*rows)[k][j] >= (*rows)[i][j] {
			visible = false
		}
	}
	if visible {
		return true
	}

	visible = true
	for k := i + 1; k < len(*rows); k++ {
		if (*rows)[k][j] >= (*rows)[i][j] {
			visible = false
		}
	}
	if visible {
		return true
	}

	return false
}
