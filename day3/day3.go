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

	sum1, sum2 := findSum(&input)

	fmt.Printf("Sum1: %d\n", sum1)
	fmt.Printf("Sum2: %d\n", sum2)
}

func findSum(input *string) (sum1, sum2 int) {
	lines := strings.Split(*input, "\n")

	for i, rucksack := range lines {
		if rucksack == "" {
			continue
		}
		size := len(rucksack) / 2

		itemType := findCommonItemType(rucksack[:size], rucksack[size:])
		sum1 += getPriority(itemType)

		if (i+1)%3 == 0 {
			badgeType := findBadgeType(lines[i-2], lines[i-1], lines[i])
			sum2 += getPriority(badgeType)
		}
	}

	return
}

func findCommonItemType(compartment1, compartment2 string) (common rune) {
	for _, itemType1 := range compartment1 {
		for _, itemType2 := range compartment2 {
			if itemType1 == itemType2 {
				return itemType1
			}
		}
	}
	return
}

func findBadgeType(rucksack1, rucksack2, rucksack3 string) (common rune) {
	itemTypes1 := make(map[rune]bool)
	for _, r := range rucksack1 {
		itemTypes1[r] = true
	}

	itemTypes2 := make(map[rune]bool)
	for _, r := range rucksack2 {
		itemTypes2[r] = true
	}

	for _, r := range rucksack3 {
		f1 := itemTypes1[r]
		f2 := itemTypes2[r]

		if f1 && f2 {
			return r
		}
	}

	return
}

func getPriority(itemType rune) int {
	ascii := int(itemType)
	if ascii < 97 {
		return ascii - 38
	}
	return ascii - 96
}
