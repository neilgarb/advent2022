package main

import (
	"fmt"
	"strings"

	"github.com/neilgarb/advent2022/util"
)

func main() {
	part1()
	part2()
}

func part1() {
	part(2)
}

func part2() {
	part(10)
}

func part(knotCount int) {
	lines := util.MustReadFile("input.txt")
	seen := make(map[util.V2]bool)
	rope := make([]util.V2, knotCount)
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
}

func moveTo(from, to util.V2) util.V2 {
	diffX := to.X - from.X
	diffY := to.Y - from.Y
	if util.Abs(diffX) > 1 && util.Abs(diffY) > 1 {
		from.X += util.Sign(diffX)
		from.Y += util.Sign(diffY)
	} else if util.Abs(diffX) > 1 {
		from.X += util.Sign(diffX)
		from.Y = to.Y
	} else if util.Abs(diffY) > 1 {
		from.Y += util.Sign(diffY)
		from.X = to.X
	}
	return from
}

func move(h util.V2, dir string) util.V2 {
	switch dir {
	case "U":
		h.Y++
	case "D":
		h.Y--
	case "L":
		h.X--
	case "R":
		h.X++
	}
	return h
}
