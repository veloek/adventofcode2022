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

	sum1 := findSum(&input, func(r1, r2 [2]int) bool {
		return contains(r1, r2) || contains(r2, r1)
	})

	sum2 := findSum(&input, func(r1, r2 [2]int) bool {
		return overlaps(r1, r2)
	})

	fmt.Printf("Sum1: %d\n", sum1)
	fmt.Printf("Sum2: %d\n", sum2)
}

func findSum(input *string, match func(r1, r2 [2]int) bool) (sum int) {
	lines := strings.Split(*input, "\n")

	for _, pair := range lines {
		if pair == "" {
			continue
		}

		split := strings.Split(pair, ",")

		range1 := parseRange(split[0])
		range2 := parseRange(split[1])

		if match(range1, range2) {
			sum++
		}
	}

	return
}

func parseRange(s string) [2]int {
	split := strings.Split(s, "-")

	start, _ := strconv.Atoi(split[0])
	end, _ := strconv.Atoi(split[1])

	return [2]int{start, end}
}

func contains(r1, r2 [2]int) bool {
	return r1[0] <= r2[0] && r1[1] >= r2[1]
}

func overlaps(r1, r2 [2]int) bool {
	return (r1[0] >= r2[0] && r1[0] <= r2[1]) || // ----2---1---2---1----
		(r1[1] >= r2[0] && r1[1] <= r2[1]) || //    ----1---2---1---2----
		(r1[0] < r2[0] && r1[1] > r2[1]) //         ----1---2---2---1----
}
