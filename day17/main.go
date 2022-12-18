package main

import (
	"fmt"

	"github.com/neilgarb/advent2022/util"
)

func main() {
	part1()
	part2()
}

func part1() {
	part(false, 2022)
}

func part2() {
	part(true, 1_000_000_000_000)
}

func part(log bool, n int) {
	lines := util.MustReadFile("input.txt")
	moves := lines[0]
	var move int
	grid := make(map[util.V2]struct{})
	var maxHeight int
	for i := 0; i < n; i++ {
		if log {
			fmt.Println(i, maxHeight)
		}
		r := rocks[i%len(rocks)]
		width := widths[i%len(rocks)]
		rc := util.V2{2, maxHeight + 3}
		for {
			if moves[move%len(moves)] == '>' {
				if canMove(r, width, rc, grid, util.V2{1, 0}) {
					rc.X++
				}
			} else if moves[move%len(moves)] == '<' {
				if canMove(r, width, rc, grid, util.V2{-1, 0}) {
					rc.X--
				}
			}
			move++
			if canMove(r, width, rc, grid, util.V2{0, -1}) {
				rc.Y--
			} else {
				for _, p := range r {
					grid[p.Add(rc)] = struct{}{}
					maxHeight = util.Max(maxHeight, p.Y+rc.Y+1)
				}
				break
			}
		}
	}
	fmt.Println(maxHeight)
}

var rocks = []rock{
	{{0, 0}, {1, 0}, {2, 0}, {3, 0}},         // -
	{{0, 1}, {1, 0}, {1, 1}, {1, 2}, {2, 1}}, // +
	{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}, // _|
	{{0, 0}, {0, 1}, {0, 2}, {0, 3}},         // |
	{{0, 0}, {1, 0}, {0, 1}, {1, 1}},         // #
}

var widths = []int{
	4,
	3,
	3,
	1,
	2,
}

type rock []util.V2

func canMove(r rock, width int, rc util.V2, grid map[util.V2]struct{}, off util.V2) bool {
	newRock := rc.Add(off)
	if newRock.X < 0 || newRock.Y < 0 {
		return false
	}
	if newRock.X+width > 7 {
		return false
	}
	for _, p := range r {
		if _, ok := grid[p.Add(newRock)]; ok {
			return false
		}
	}
	return true
}
