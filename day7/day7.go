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

	root := parseInput(&input)

	fmt.Printf("Result 1: %d\n", solve1(root))
	fmt.Printf("Result 2: %d\n", solve2(root))
}

func solve1(root *dir) int {
	dirs := find(root, func(s int) bool {
		return s < 100000
	})

	sum := 0
	for _, d := range dirs {
		sum += d.size()
	}

	return sum
}

func solve2(root *dir) int {
	used := 70000000 - root.size()
	needed := 30000000 - used

	dirs := find(root, func(s int) bool {
		return s >= needed
	})

	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].size() < dirs[j].size()
	})

	return dirs[0].size()
}

func find(d *dir, predicate func(int) bool) []*dir {
	r := make([]*dir, 0)

	if predicate(d.size()) {
		r = append(r, d)
	}

	for _, s := range d.subdirs {
		r = append(r, find(s, predicate)...)
	}

	return r
}

func parseInput(input *string) *dir {
	lines := strings.Split(*input, "\n")
	cdRe := regexp.MustCompile("^\\$ cd (\\S+)$")
	dirRe := regexp.MustCompile("^dir (\\S+)$")
	fileRe := regexp.MustCompile("^(\\d+) (\\S+)$")

	var curDir *dir
	for _, line := range lines {
		if line == "" {
			continue
		}

		if r := cdRe.FindStringSubmatch(line); r != nil {
			if r[1] == "/" {
				curDir = newDir()
				curDir.name = "/"
			} else if r[1] == ".." {
				curDir = curDir.parent
			} else {
				for _, s := range curDir.subdirs {
					if s.name == r[1] {
						curDir = s
						break
					}
				}
			}
		} else if r := dirRe.FindStringSubmatch(line); r != nil {
			s := newDir()
			s.name = r[1]
			s.parent = curDir
			curDir.subdirs = append(curDir.subdirs, s)
		} else if r := fileRe.FindStringSubmatch(line); r != nil {
			f := file{}
			f.name = r[2]
			size, _ := strconv.Atoi(r[1])
			f.size = size
			curDir.files = append(curDir.files, &f)
		}
	}

	for curDir.parent != nil {
		curDir = curDir.parent
	}

	return curDir
}

type dir struct {
	parent  *dir
	name    string
	files   []*file
	subdirs []*dir
}

func newDir() *dir {
	d := dir{}
	d.files = make([]*file, 0)
	d.subdirs = make([]*dir, 0)
	return &d
}

func (d *dir) size() int {
	size := 0

	for _, f := range d.files {
		size += f.size
	}

	for _, d := range d.subdirs {
		size += d.size()
	}

	return size
}

type file struct {
	name string
	size int
}
