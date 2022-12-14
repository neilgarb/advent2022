package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/neilgarb/advent2022/util"
)

func main() {
	util.MustDo(part1)
	util.MustDo(part2)
}

func part1() error {
	lines := util.MustReadFile("input.txt")
	rock, _, max := parse(lines)
	sand := make(map[c]bool)

SandLoop:
	for {
		cur := c{500, 0}
	GrainLoop:
		for {
			for _, x := range []int{0, -1, 1} {
				t := c{cur.x + x, cur.y + 1}
				if !rock[t] && !sand[t] {
					cur = t
					if cur.y >= max.y {
						break SandLoop
					}
					continue GrainLoop
				}
			}
			sand[cur] = true
			continue SandLoop
		}
	}

	fmt.Println(len(sand))
	return nil
}

func part2() error {
	lines := util.MustReadFile("input.txt")
	rock, _, max := parse(lines)
	sand := make(map[c]bool)

SandLoop:
	for {
		cur := c{500, 0}
	GrainLoop:
		for {
			for _, x := range []int{0, -1, 1} {
				t := c{cur.x + x, cur.y + 1}
				if !rock[t] && !sand[t] {
					cur = t
					if cur.y >= max.y+1 {
						sand[cur] = true
						continue SandLoop
					}
					continue GrainLoop
				}
			}
			sand[cur] = true
			if cur.x == 500 && cur.y == 0 {
				break SandLoop
			}
			continue SandLoop
		}
	}

	fmt.Println(len(sand))
	return nil
}

func parse(lines []string) (map[c]bool, c, c) {
	res := make(map[c]bool)
	var min, max c
	min.x, min.y = math.MaxInt, math.MaxInt
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		for i := 1; i < len(parts); i++ {
			for _, p := range draw(parts[i-1], parts[i]) {
				res[p] = true
				if p.x < min.x {
					min.x = p.x
				}
				if p.y < min.y {
					min.y = p.y
				}
				if p.x > max.x {
					max.x = p.x
				}
				if p.y > max.y {
					max.y = p.y
				}
			}
		}
	}
	return res, min, max
}

type c struct {
	x, y int
}

func draw(froms, tos string) []c {
	fromxs, fromys, _ := strings.Cut(froms, ",")
	toxs, toys, _ := strings.Cut(tos, ",")
	fromx, fromy := util.MustParseInt(fromxs), util.MustParseInt(fromys)
	tox, toy := util.MustParseInt(toxs), util.MustParseInt(toys)
	if fromx > tox || fromy > toy {
		fromx, tox, fromy, toy = tox, fromx, toy, fromy
	}
	var res []c
	for y := fromy; y <= toy; y++ {
		for x := fromx; x <= tox; x++ {
			res = append(res, c{x, y})
		}
	}
	return res
}
