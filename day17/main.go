package main

import (
	"fmt"

	"github.com/neilgarb/advent2022/util"
)

func main() {
	util.MustDo(part1)
	util.MustDo(part2)
}

func part1() error {
	return part(false, 2022)
}

func part2() error {
	return part(true, 1_000_000_000_000)
}

func part(log bool, n int) error {
	lines := util.MustReadFile("input.txt")
	moves := lines[0]
	var move int
	grid := make(map[c]struct{})
	var maxHeight int
	for i := 0; i < n; i++ {
		if log {
			fmt.Println(i, maxHeight)
		}

		r := rocks[i%len(rocks)]
		width := widths[i%len(rocks)]
		rc := c{2, maxHeight + 3}

		for {
			if moves[move%len(moves)] == '>' {
				if canmove(r, width, rc, grid, c{1, 0}) {
					rc.x++
				}
			} else if moves[move%len(moves)] == '<' {
				if canmove(r, width, rc, grid, c{-1, 0}) {
					rc.x--
				}
			}
			move++
			if canmove(r, width, rc, grid, c{0, -1}) {
				rc.y--
			} else {
				for _, p := range r {
					grid[c{p.x + rc.x, p.y + rc.y}] = struct{}{}
					maxHeight = util.Max(maxHeight, p.y+rc.y+1)
				}
				break
			}
		}
	}

	fmt.Println(maxHeight)

	return nil
}

type c struct {
	x, y int
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

type rock []c

func canmove(r rock, width int, rc c, grid map[c]struct{}, off c) bool {
	newrc := c{rc.x + off.x, rc.y + off.y}
	if newrc.x < 0 || newrc.y < 0 {
		return false
	}
	if newrc.x+width > 7 {
		return false
	}
	for _, p := range r {
		if _, ok := grid[c{p.x + newrc.x, p.y + newrc.y}]; ok {
			return false
		}
	}
	return true
}
