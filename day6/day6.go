package main

import (
	"aoc/common"
	"fmt"
	"log"
	"sort"
)

func main() {
	input, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result1 := solve(&input, 4)
	result2 := solve(&input, 14)
	fmt.Printf("Result1: %v\n", result1)
	fmt.Printf("Result2: %v\n", result2)
}

func solve(input *string, n int) int {
	buf := newBuffer(n)
	for idx, c := range *input {
		buf.add(int(c))
		if idx >= n-1 && !buf.containsDuplicates() {
			return idx + 1
		}
	}
	return -1 // Not found
}

type buffer []int

func newBuffer(size int) *buffer {
	buf := make(buffer, size)
	return &buf
}

func (s *buffer) add(i int) {
	*s = append((*s)[1:], i)
}

func (s *buffer) containsDuplicates() bool {
	// Make a copy as sorting mutates the slice
	dup := make(buffer, len(*s))
	copy(dup, *s)
	sort.Sort(sort.IntSlice(dup))

	for idx, el := range dup {
		if idx > 0 && el == dup[idx-1] {
			return true
		}
	}
	return false
}
