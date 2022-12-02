package main

import (
	"aoc/common"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	top1 := findTop1(&input)
	top3 := findTop3(&input)

	fmt.Printf("Top 1: %d\n", top1)
	fmt.Printf("Top 3: %d\n", top3)
}

func findTop1(input *string) int {
	lines := strings.Split(*input, "\n")
	max := 0
	sum := 0

	for _, cals := range lines {
		if cals == "" {
			if sum > max {
				max = sum
			}
			sum = 0
		} else {
			i, _ := strconv.Atoi(cals)
			sum += i
		}
	}

	return max
}

func findTop3(input *string) int {
	lines := strings.Split(*input, "\n")
	sums := make([]int, 0, len(lines))
	sum := 0

	for _, cals := range lines {
		if cals == "" {
			sums = append(sums, sum)
			sum = 0
		} else {
			i, _ := strconv.Atoi(cals)
			sum += i
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	return sums[0] + sums[1] + sums[2]
}
