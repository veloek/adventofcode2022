package main

import (
	"aoc/common"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := solve(&input)

	fmt.Printf("Result1: %d\n", result)
}

type sensor struct {
	pos    pos
	beacon pos
}

type pos struct {
	x, y int
}

func solve(input *string) int {
	lines := strings.Split(*input, "\n")

	sensors := make([]*sensor, len(lines))

	posRe := regexp.MustCompile("x=(-?\\d+), y=(-?\\d+)")
	parsePos := func(s string) pos {
		m := posRe.FindStringSubmatch(s)
		x, _ := strconv.Atoi(m[1])
		y, _ := strconv.Atoi(m[2])
		return pos{x, y}
	}

	for n, line := range lines {
		if line == "" {
			continue
		}
		split := strings.Split(line, ":")
		pos := parsePos(split[0])
		beacon := parsePos(split[1])
		sensors[n] = &sensor{pos, beacon}
	}

	occ := make(map[int]bool)

	targetY := 2000000

	for _, s := range sensors {
		if s == nil {
			continue
		}
		reach := distance(&s.pos, &s.beacon)
		target := pos{s.pos.x, targetY}
		dist := distance(&s.pos, &target)

		if dist > reach {
			continue
		}

		rrr := reach - dist

		for i := -rrr; i <= rrr; i++ {
			p := s.pos.x + i
			if _, found := occ[p]; !found {
				occ[p] = true
			}
		}
	}

	for _, x := range sensors {
		if x != nil && x.beacon.y == targetY {
			delete(occ, x.beacon.x)
		}
	}

	return len(occ)
}

func distance(a, b *pos) int {
	dx := math.Abs(float64(a.x) - float64(b.x))
	dy := math.Abs(float64(a.y) - float64(b.y))
	return int(dx + dy)
}
