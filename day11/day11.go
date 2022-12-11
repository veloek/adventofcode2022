package main

import (
	"aoc/common"
	"fmt"
	"log"
	"regexp"
	"sort"
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

func solve(input *string, part2 bool) int {
	var n int
	if part2 {
		n = 10000
	} else {
		n = 20
	}

	monkeys := parseMonkeys(input)

	// Find product of all divisors which we can use to
	// manage worry levels
	testProd := 1
	for _, m := range monkeys {
		testProd *= m.test
	}

	for i := 0; i < n; i++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				item = m.operation(item)
				if !part2 {
					item /= 3
				} else {
					item %= testProd
				}
				var receiver *monkey
				if item%m.test == 0 {
					receiver = monkeys[m.yes]
				} else {
					receiver = monkeys[m.no]
				}
				receiver.items = append(receiver.items, item)
				m.cnt++
			}
			m.items = []int{}
		}
	}

	sort.Slice(monkeys, func(a, b int) bool {
		return monkeys[a].cnt > monkeys[b].cnt
	})

	monkeyBusiness := monkeys[0].cnt * monkeys[1].cnt

	return monkeyBusiness
}

func parseMonkeys(input *string) []*monkey {
	specs := strings.Split(*input, "\n\n")

	monkeys := make([]*monkey, 0)

	for _, spec := range specs {
		m := parseMonkey(&spec)
		monkeys = append(monkeys, m)
	}
	return monkeys
}

func parseMonkey(input *string) *monkey {
	itemsRe := regexp.MustCompile("Starting items:(.*)$")
	operationRe := regexp.MustCompile("Operation: new = old (\\+|\\-|\\*|\\/) (\\d+|old)$")
	testRe := regexp.MustCompile("Test: divisible by (\\d+)$")
	ifRe := regexp.MustCompile("If (true|false): throw to monkey (\\d+)$")

	var items []int
	var operation func(int) int
	var test int
	var yes int
	var no int

	lines := strings.Split(*input, "\n")
	for _, line := range lines {
		if m := itemsRe.FindStringSubmatch(line); m != nil {
			split := strings.Split(m[1], ",")
			for _, i := range split {
				item, _ := strconv.Atoi(strings.Trim(i, " "))
				items = append(items, item)
			}
		}

		if m := operationRe.FindStringSubmatch(line); m != nil {
			operation = func(i int) int {
				var n int
				if nn, err := strconv.Atoi(m[2]); err != nil {
					n = i
				} else {
					n = nn
				}

				switch m[1] {
				case "+":
					return i + n
				case "-":
					return i - n
				case "*":
					return i * n
				case "/":
					return i / n
				}
				return 0
			}
		}

		if m := testRe.FindStringSubmatch(line); m != nil {
			test, _ = strconv.Atoi(m[1])
		}

		if m := ifRe.FindStringSubmatch(line); m != nil {
			i, _ := strconv.Atoi(m[2])
			if m[1] == "true" {
				yes = i
			} else {
				no = i
			}
		}
	}

	m := monkey{
		items:     items,
		operation: operation,
		test:      test,
		yes:       yes,
		no:        no,
	}
	return &m
}

type monkey struct {
	items     []int
	operation func(int) int
	test      int
	yes       int
	no        int
	cnt       int
}
