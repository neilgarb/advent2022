package main

import (
	"fmt"
	"math"
	"regexp"

	"github.com/neilgarb/advent2022/util"
)

func main() {
	util.MustDo(part1)
	util.MustDo(part2)
}

var re = regexp.MustCompile(`\-?\d+`)

func part1() error {
	lines := util.MustReadFile("input.txt")
	sensors := make(map[c]c)
	beacons := make(map[c]bool)
	min, max := math.MaxInt, 0
	for _, line := range lines {
		matches := re.FindAllString(line, 4)
		s := c{util.MustParseInt(matches[0]), util.MustParseInt(matches[1])}
		b := c{util.MustParseInt(matches[2]), util.MustParseInt(matches[3])}
		sensors[s] = b
		beacons[b] = true
		min = util.Min(min, s.x-man(s, b))
		max = util.Max(max, s.x+man(s, b))
	}
	y := 2000000
	var tot int
	for x := min; x <= max; x++ {
		t := c{x, y}
		if _, ok := beacons[t]; ok {
			continue
		}
		if _, ok := sensors[t]; ok {
			continue
		}
		for s, b := range sensors {
			if man(t, s) <= man(b, s) {
				tot++
				break
			}
		}
	}
	fmt.Println(tot)
	return nil
}

func part2() error {
	lines := util.MustReadFile("input.txt")
	sensors := make(map[c]c)
	for _, line := range lines {
		matches := re.FindAllString(line, 4)
		s := c{util.MustParseInt(matches[0]), util.MustParseInt(matches[1])}
		b := c{util.MustParseInt(matches[2]), util.MustParseInt(matches[3])}
		sensors[s] = b
	}
	for s, b := range sensors {
		d := man(s, b)
		var tests []c
		for x := -d; x <= d; x++ {
			t := util.Abs(x)
			y := d - t
			tests = append(tests, c{s.x + x, s.y + y + 1})
			tests = append(tests, c{s.x + x, s.y - y - 1})
		}
		tests = append(tests, c{s.x - d - 1, s.y})
		tests = append(tests, c{s.x + d + 1, s.y})
		for _, t := range tests {
			if t.x < 0 || t.x > 4000000 || t.y < 0 || t.y > 4000000 {
				continue
			}
			var found bool
			for os, ob := range sensors {
				if man(os, t) <= man(os, ob) {
					found = true
					break
				}
			}
			if !found {
				fmt.Println(t.x*4000000 + t.y)
				return nil
			}
		}
	}
	return nil
}

type c struct {
	x, y int
}

func man(from, to c) int {
	return util.Diff(from.x, to.x) + util.Diff(from.y, to.y)
}
