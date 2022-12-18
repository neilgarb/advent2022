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

func part1() error {
	return part(false)
}

func part2() error {
	return part(true)
}

func part(checkPockets bool) error {
	lines := util.MustReadFile("input.txt")
	cubes := make(map[c]bool)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		cubes[c{
			util.MustParseInt(parts[0]),
			util.MustParseInt(parts[1]),
			util.MustParseInt(parts[2]),
		}] = true
	}

	var exposed int
	for cube := range cubes {
		for _, s := range sides {
			newc := c{cube.x + s.x, cube.y + s.y, cube.z + s.z}
			if !cubes[newc] {
				exposed++
				if checkPockets {
					if isPocket(cubes, newc, 1000, map[c]bool{newc: true}) {
						exposed--
					}
				}
			}
		}
	}

	fmt.Println(exposed)
	return nil
}

var sides = []c{
	{0, 0, 1},
	{0, 0, -1},
	{0, 1, 0},
	{0, -1, 0},
	{1, 0, 0},
	{-1, 0, 0},
}

type c struct {
	x, y, z int
}

func isPocket(cubes map[c]bool, cube c, max int, seen map[c]bool) bool {
	if max == 0 {
		return false
	}
	res := true
	for _, s := range sides {
		t := c{cube.x + s.x, cube.y + s.y, cube.z + s.z}
		if !cubes[t] && !seen[t] {
			seen[t] = true
			if !isPocket(cubes, t, max-1, seen) {
				res = false
				break
			}
		}
	}
	return res
}
