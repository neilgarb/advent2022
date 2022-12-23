package main

import (
	"fmt"

	. "github.com/neilgarb/advent2022/util"
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

func part(part2 bool) {
	lines := MustReadFile("input.txt")
	elves := make(map[V2]bool)
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				elves[V2{x, y}] = true
			}
		}
	}

	if !part2 {
		var dirIdx int
		for i := 0; i < 10; i++ {
			round(elves, dirIdx)
			dirIdx = (dirIdx + 1) % 4
		}
		min, max := minmax(elves)
		fmt.Println((max.X-min.X+1)*(max.Y-min.Y+1) - len(elves))
	} else {
		var dirIdx int
		i := 1
		for {
			if !round(elves, dirIdx) {
				fmt.Println(i)
				break
			}
			dirIdx = (dirIdx + 1) % 4
			i++
		}
	}
}

func round(elves map[V2]bool, dirIdx int) bool {
	proposals := make(map[V2][]V2)

	// First half.
ElfLoop:
	for e := range elves {
		var haveNeighbours bool
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if x == 0 && y == 0 {
					continue
				} else if elves[e.Add(V2{x, y})] {
					haveNeighbours = true
				}
			}
		}
		if haveNeighbours {
			for i := 0; i < 4; i++ {
				switch dirs[(dirIdx+i)%4] {
				case 0:
					if !elves[e.Add(V2{1, 0})] && !elves[e.Add(V2{1, 1})] && !elves[e.Add(V2{1, -1})] {
						proposals[e.Add(V2{1, 0})] = append(proposals[e.Add(V2{1, 0})], e)
						continue ElfLoop
					}
				case 1:
					if !elves[e.Add(V2{0, 1})] && !elves[e.Add(V2{1, 1})] && !elves[e.Add(V2{-1, 1})] {
						proposals[e.Add(V2{0, 1})] = append(proposals[e.Add(V2{0, 1})], e)
						continue ElfLoop
					}
				case 2:
					if !elves[e.Add(V2{-1, 0})] && !elves[e.Add(V2{-1, 1})] && !elves[e.Add(V2{-1, -1})] {
						proposals[e.Add(V2{-1, 0})] = append(proposals[e.Add(V2{-1, 0})], e)
						continue ElfLoop
					}
				case 3:
					if !elves[e.Add(V2{0, -1})] && !elves[e.Add(V2{1, -1})] && !elves[e.Add(V2{-1, -1})] {
						proposals[e.Add(V2{0, -1})] = append(proposals[e.Add(V2{0, -1})], e)
						continue ElfLoop
					}
				}
			}
		}
	}

	// Second half.
	for k, v := range proposals {
		if len(v) > 1 {
			continue
		}
		delete(elves, v[0])
		elves[k] = true
	}

	return len(proposals) > 0
}

var dirs = []int{3, 1, 2, 0} // North, south, west, east

func minmax(elves map[V2]bool) (V2, V2) {
	var min, max V2
	for e := range elves {
		min.X = Min(min.X, e.X)
		min.Y = Min(min.Y, e.Y)
		max.X = Max(max.X, e.X)
		max.Y = Max(max.Y, e.Y)
	}
	return min, max
}
