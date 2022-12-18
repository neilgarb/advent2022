package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/neilgarb/advent2022/util"
)

func main() {
	part1()
	part2()
}

func part1() {
	lines := util.MustReadFile("input.txt")
	rock, _, max := parse(lines)
	sand := make(map[util.V2]bool)
SandLoop:
	for {
		cur := util.V2{500, 0}
	GrainLoop:
		for {
			for _, x := range []int{0, -1, 1} {
				t := cur.Add(util.V2{x, 1})
				if !rock[t] && !sand[t] {
					cur = t
					if cur.Y >= max.Y {
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
}

func part2() {
	lines := util.MustReadFile("input.txt")
	rock, _, max := parse(lines)
	sand := make(map[util.V2]bool)
SandLoop:
	for {
		cur := util.V2{500, 0}
	GrainLoop:
		for {
			for _, x := range []int{0, -1, 1} {
				t := cur.Add(util.V2{x, 1})
				if !rock[t] && !sand[t] {
					cur = t
					if cur.Y >= max.Y+1 {
						sand[cur] = true
						continue SandLoop
					}
					continue GrainLoop
				}
			}
			sand[cur] = true
			if cur.X == 500 && cur.Y == 0 {
				break SandLoop
			}
			continue SandLoop
		}
	}
	fmt.Println(len(sand))
}

func parse(lines []string) (map[util.V2]bool, util.V2, util.V2) {
	res := make(map[util.V2]bool)
	var min, max util.V2
	min.X, min.Y = math.MaxInt, math.MaxInt
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		for i := 1; i < len(parts); i++ {
			for _, p := range draw(parts[i-1], parts[i]) {
				res[p] = true
				if p.X < min.X {
					min.X = p.X
				}
				if p.Y < min.Y {
					min.Y = p.Y
				}
				if p.X > max.X {
					max.X = p.X
				}
				if p.Y > max.Y {
					max.Y = p.Y
				}
			}
		}
	}
	return res, min, max
}

func draw(fromStr, toStr string) []util.V2 {
	fromXStr, fromYStr, _ := strings.Cut(fromStr, ",")
	toXStr, toYStr, _ := strings.Cut(toStr, ",")
	fromX, fromY := util.MustParseInt(fromXStr), util.MustParseInt(fromYStr)
	tox, toy := util.MustParseInt(toXStr), util.MustParseInt(toYStr)
	if fromX > tox || fromY > toy {
		fromX, tox, fromY, toy = tox, fromX, toy, fromY
	}
	var res []util.V2
	for y := fromY; y <= toy; y++ {
		for x := fromX; x <= tox; x++ {
			res = append(res, util.V2{x, y})
		}
	}
	return res
}
