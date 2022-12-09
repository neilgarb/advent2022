package main

import (
	"fmt"
	"strings"

	"github.com/neilgarb/advent2022/util"
)

func main() {
	util.MustDo(part1)
	util.MustDo(part2)
}

type c struct {
	x, y int
}

func part1() error {
	return part(2)
}

func part2() error {
	return part(10)
}

func part(knotCount int) error {
	lines := util.MustReadFile("input.txt")
	seen := make(map[c]bool)
	rope := make([]c, knotCount)
	seen[rope[knotCount-1]] = true
	for _, line := range lines {
		dir, cnt, _ := strings.Cut(line, " ")
		for i := 0; i < util.MustParseInt(cnt); i++ {
			rope[0] = move(rope[0], dir)
			for j := 1; j < knotCount; j++ {
				rope[j] = moveTo(rope[j], rope[j-1])
			}
			seen[rope[knotCount-1]] = true
		}
	}

	fmt.Println(len(seen))
	return nil
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func sign(i int) int {
	if i < 0 {
		return -1
	}
	return 1
}

func moveTo(from, to c) c {
	diffX := to.x - from.x
	diffY := to.y - from.y
	if abs(diffX) > 1 && abs(diffY) > 1 {
		from.x += sign(diffX)
		from.y += sign(diffY)
	} else if abs(diffX) > 1 {
		from.x += sign(diffX)
		from.y = to.y
	} else if abs(diffY) > 1 {
		from.y += sign(diffY)
		from.x = to.x
	}
	return from
}

func move(h c, dir string) c {
	switch dir {
	case "U":
		h.y++
	case "D":
		h.y--
	case "L":
		h.x--
	case "R":
		h.x++
	}
	return h
}
