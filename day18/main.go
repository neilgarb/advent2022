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
	part(false)
}

func part2() {
	part(true)
}

func part(checkPockets bool) {
	lines := util.MustReadFile("input.txt")
	cubes := make(map[util.V3]bool)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		cubes[util.V3{
			util.MustParseInt(parts[0]),
			util.MustParseInt(parts[1]),
			util.MustParseInt(parts[2]),
		}] = true
	}
	var exposed int
	for cube := range cubes {
		for _, s := range sides {
			newCube := cube.Add(s)
			if !cubes[newCube] {
				exposed++
				if checkPockets {
					if isPocket(cubes, newCube, 1000, map[util.V3]bool{newCube: true}) {
						exposed--
					}
				}
			}
		}
	}
	fmt.Println(exposed)
}

var sides = []util.V3{
	{0, 0, 1},
	{0, 0, -1},
	{0, 1, 0},
	{0, -1, 0},
	{1, 0, 0},
	{-1, 0, 0},
}

func isPocket(cubes map[util.V3]bool, cube util.V3, max int, seen map[util.V3]bool) bool {
	if max == 0 {
		return false
	}
	res := true
	for _, s := range sides {
		t := cube.Add(s)
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
