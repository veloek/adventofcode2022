package main

import (
	"aoc/common"
	"fmt"
	"log"
	"regexp"
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
	fmt.Printf("Result1: %v\n", result1)
	fmt.Printf("Result2: %v\n", result2)
	// Result1: TGWSMRBPN
	// Result2: TZLTLWRNF
}

func solve(input *string, part2 bool) string {
	lines := strings.Split(*input, "\n")
	stacks := make([][]byte, (len(lines[0])+1)/4)
	operation := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")
	parseOperations := false

	for _, line := range lines {
		if line == "" {
			parseOperations = true
			continue
		}

		if parseOperations {
			res := operation.FindStringSubmatch(line)
			if len(res) == 4 {
				n, _ := strconv.Atoi(res[1])
				from, _ := strconv.Atoi(res[2])
				to, _ := strconv.Atoi(res[3])

				if part2 {
					newTo := append([]byte(nil), stacks[from-1][:n]...)
					stacks[to-1] = append(newTo, stacks[to-1]...)
				} else {
					for i := 0; i < n; i++ {
						stacks[to-1] = append([]byte{stacks[from-1][i]}, stacks[to-1]...)
					}
				}
				stacks[from-1] = stacks[from-1][n:]
			}
		} else {
			for i := 0; i < len(line); i += 4 {
				if line[i+1] >= 65 && line[i+1] <= 90 {
					stacks[i/4] = append(stacks[i/4], line[i+1])
				}
			}
		}
	}

	result := make([]byte, len(stacks))
	for i := 0; i < len(stacks); i++ {
		result[i] = stacks[i][0]
	}

	return string(result)
}
