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
	fmt.Printf("Result2: %s\n", result2)
}

func solve(input *string) (int, string) {
	lines := strings.Split(*input, "\n")
	pc := 0
	x := 1
	sum := 0

	checkSignalStrength := func() {
		if (pc+20)%40 == 0 {
			sum += pc * x
		}
	}

	crt := newCRT(40, 6)

	for _, line := range lines {
		if line == "" {
			continue
		}

		if line == "noop" {
			// do nothing with x
			crt.draw(pc, x)

			pc++
			checkSignalStrength()
			continue
		}

		// addx
		split := strings.Split(line, " ")
		inc, _ := strconv.Atoi(split[1])
		// start first instruction
		crt.draw(pc, x)
		pc++
		checkSignalStrength()
		// start second instruction
		crt.draw(pc, x)
		pc++
		checkSignalStrength()
		// end second instruction
		x += inc
	}

	// render image
	var b strings.Builder
	for i := 0; i < len(crt); i++ {
		if i%40 == 0 {
			b.WriteByte('\n')
		}
		b.WriteByte(crt[i])
	}

	return sum, b.String()
}

type crt []byte

func newCRT(cols, rows int) crt {
	crt := make([]byte, cols*rows)
	return crt
}

func (crt *crt) draw(i, x int) {
	col := i % 40
	if col >= x-1 && col <= x+1 {
		(*crt)[i] = '#'
	} else {
		(*crt)[i] = '.'
	}
}
