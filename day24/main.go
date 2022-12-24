package main

import (
	"fmt"
	"math"

	. "github.com/neilgarb/advent2022/util"
)

func main() {
	part1()
	part2()
}

func part1() {
	blizzards, width, height := parse("input.txt")
	freeRepeat := int(math.Sqrt(float64(height)-2) * (float64(width) - 2))
	start := V2{1, 0}
	end := V2{width - 2, height - 1}
	freeFn := makeFreeFn(blizzards, width, height)
	fmt.Println(simulate(freeFn, start, end, freeRepeat, 0))
}

func part2() {
	blizzards, width, height := parse("input.txt")
	freeRepeat := int(math.Sqrt(float64(height)-2) * (float64(width) - 2))
	start := V2{1, 0}
	end := V2{width - 2, height - 1}
	freeFn := makeFreeFn(blizzards, width, height)
	min := simulate(freeFn, start, end, freeRepeat, 0)
	min = simulate(freeFn, end, start, freeRepeat, min)
	min = simulate(freeFn, start, end, freeRepeat, min)
	fmt.Println(min)
}

func parse(file string) (map[V3]bool, int, int) {
	lines := MustReadFile(file)
	blizzards := make(map[V3]bool)
	width := len(lines[0])
	height := len(lines)
	for y, line := range lines {
		for x, c := range line {
			if c == '>' {
				blizzards[V3{x, y, 0}] = true
			} else if c == 'v' {
				blizzards[V3{x, y, 1}] = true
			} else if c == '<' {
				blizzards[V3{x, y, 2}] = true
			} else if c == '^' {
				blizzards[V3{x, y, 3}] = true
			}
		}
	}
	return blizzards, width, height
}

func simulate(freeFn func(int) map[V2]bool, start, end V2, freeRepeat, startMin int) int {
	set := []state{{startMin, start}}
	seen := map[state]bool{{startMin, start}: true}
	min := math.MaxInt
	for len(set) > 0 {
		// Pick the closest to the end.
		closestIdx := 0
		closestDist := man(set[0].pos, end)
		for i := 1; i < len(set); i++ {
			if d := man(set[i].pos, end); d < closestDist {
				closestIdx = i
				closestDist = d
			}
		}
		s := set[closestIdx]
		set = append(set[:closestIdx], set[closestIdx+1:]...)
		if s.min >= min {
			continue
		}

		if s.pos == end {
			if s.min < min {
				min = s.min
			}
			// Prune.
			for i := 0; i < len(set); i++ {
				if set[i].min >= s.min {
					set = append(set[:i], set[i+1:]...)
					i--
				}
			}
			continue
		}

		free := freeFn(s.min + 1)
		nextMin := s.min + 1

		// Wait.
		waitState := state{nextMin, s.pos}
		if free[s.pos] && !seen[state{nextMin % freeRepeat, s.pos}] {
			set = append(set, waitState)
			seen[state{nextMin % freeRepeat, waitState.pos}] = true
		}

		// Move.
		rightState := state{s.min + 1, s.pos.Add(V2{1, 0})}
		if free[rightState.pos] && !seen[state{nextMin % freeRepeat, rightState.pos}] {
			set = append(set, rightState)
			seen[state{nextMin % freeRepeat, rightState.pos}] = true
		}
		downState := state{s.min + 1, s.pos.Add(V2{0, 1})}
		if free[downState.pos] && !seen[state{nextMin % freeRepeat, downState.pos}] {
			set = append(set, downState)
			seen[state{nextMin % freeRepeat, downState.pos}] = true
		}
		leftState := state{s.min + 1, s.pos.Add(V2{-1, 0})}
		if free[leftState.pos] && !seen[state{nextMin % freeRepeat, leftState.pos}] {
			set = append(set, leftState)
			seen[state{nextMin % freeRepeat, leftState.pos}] = true
		}
		upState := state{s.min + 1, s.pos.Add(V2{0, -1})}
		if free[upState.pos] && !seen[state{nextMin % freeRepeat, upState.pos}] {
			set = append(set, upState)
			seen[state{nextMin % freeRepeat, upState.pos}] = true
		}
	}
	return min
}

type state struct {
	min int
	pos V2
}

func move(blizzards map[V3]bool, width, height int) map[V3]bool {
	res := make(map[V3]bool)
	for b := range blizzards {
		pos := V2{b.X, b.Y}
		switch b.Z {
		case 0:
			pos = pos.Add(V2{1, 0})
			if pos.X == width-1 {
				pos.X = 1
			}
		case 1:
			pos = pos.Add(V2{0, 1})
			if pos.Y == height-1 {
				pos.Y = 1
			}
		case 2:
			pos = pos.Add(V2{-1, 0})
			if pos.X == 0 {
				pos.X = width - 2
			}
		case 3:
			pos = pos.Add(V2{0, -1})
			if pos.Y == 0 {
				pos.Y = height - 2
			}
		}
		res[V3{pos.X, pos.Y, b.Z}] = true
	}
	return res
}

func makeFreeFn(blizzards map[V3]bool, width, height int) func(int) map[V2]bool {
	free := func(blizzards map[V3]bool) map[V2]bool {
		res := make(map[V2]bool)
		for y := 1; y < height-1; y++ {
			for x := 1; x < width-1; x++ {
				if !blizzards[V3{x, y, 0}] && !blizzards[V3{x, y, 1}] &&
					!blizzards[V3{x, y, 2}] && !blizzards[V3{x, y, 3}] {
					res[V2{x, y}] = true
				}
			}
		}
		res[V2{1, 0}] = true
		res[V2{width - 2, height - 1}] = true
		return res
	}
	knownBlizzards := []map[V3]bool{blizzards}
	knownFree := []map[V2]bool{free(blizzards)}
	return func(min int) map[V2]bool {
		if min < len(knownFree) {
			return knownFree[min]
		}
		b := knownBlizzards[len(knownBlizzards)-1]
		for i := len(knownFree); i <= min; i++ {
			b = move(knownBlizzards[i-1], width, height)
			knownBlizzards = append(knownBlizzards, b)
			knownFree = append(knownFree, free(b))
		}
		return knownFree[min]
	}
}

func man(a, b V2) int {
	return Diff(a.X, b.X) + Diff(a.Y, b.Y)
}
