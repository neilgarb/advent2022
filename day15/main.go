package main

import (
	"fmt"
	"math"
	"regexp"

	"github.com/neilgarb/advent2022/util"
)

func main() {
	part1()
	part2()
}

var re = regexp.MustCompile(`\-?\d+`)

func part1() error {
	lines := util.MustReadFile("input.txt")
	sensors := make(map[util.V2]util.V2)
	beacons := make(map[util.V2]bool)
	min, max := math.MaxInt, 0
	for _, line := range lines {
		matches := re.FindAllString(line, 4)
		s := util.V2{util.MustParseInt(matches[0]), util.MustParseInt(matches[1])}
		b := util.V2{util.MustParseInt(matches[2]), util.MustParseInt(matches[3])}
		sensors[s] = b
		beacons[b] = true
		min = util.Min(min, s.X-man(s, b))
		max = util.Max(max, s.X+man(s, b))
	}
	y := 2000000
	var tot int
	for x := min; x <= max; x++ {
		t := util.V2{x, y}
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

func part2() {
	lines := util.MustReadFile("input.txt")
	sensors := make(map[util.V2]util.V2)
	for _, line := range lines {
		matches := re.FindAllString(line, 4)
		s := util.V2{util.MustParseInt(matches[0]), util.MustParseInt(matches[1])}
		b := util.V2{util.MustParseInt(matches[2]), util.MustParseInt(matches[3])}
		sensors[s] = b
	}
	for s, b := range sensors {
		d := man(s, b)
		var tests []util.V2
		for x := -d; x <= d; x++ {
			t := util.Abs(x)
			y := d - t
			tests = append(tests, s.Add(util.V2{x, y + 1}))
			tests = append(tests, s.Add(util.V2{x, -y - 1}))
		}
		tests = append(tests, s.Add(util.V2{-d - 1, 0}))
		tests = append(tests, s.Add(util.V2{d + 1, 0}))
		for _, t := range tests {
			if t.X < 0 || t.X > 4000000 || t.Y < 0 || t.Y > 4000000 {
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
				fmt.Println(t.X*4000000 + t.Y)
				return
			}
		}
	}
}

func man(from, to util.V2) int {
	return util.Diff(from.X, to.X) + util.Diff(from.Y, to.Y)
}
