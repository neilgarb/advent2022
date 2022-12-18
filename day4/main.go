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
	lines := util.MustReadFile("input.txt")
	var tot int
	for _, line := range lines {
		as, bs, _ := strings.Cut(line, ",")
		a, b := parse(as), parse(bs)
		if a.contains(b) || b.contains(a) {
			tot++
		}
	}
	fmt.Println(tot)
}

func part2() {
	lines := util.MustReadFile("input.txt")
	var tot int
	for _, line := range lines {
		as, bs, _ := strings.Cut(line, ",")
		a, b := parse(as), parse(bs)
		if a.overlaps(b) {
			tot++
		}
	}
	fmt.Println(tot)
}

type section struct {
	from, to int
}

func parse(s string) section {
	fs, ts, _ := strings.Cut(s, "-")
	return section{
		from: util.MustParseInt(fs),
		to:   util.MustParseInt(ts),
	}
}

func (s section) contains(t section) bool {
	return t.from >= s.from && t.to <= s.to
}

func (s section) overlaps(t section) bool {
	return (s.from >= t.from && s.from <= t.to) ||
		(s.to >= t.from && s.to <= t.to) ||
		(t.from >= s.from && t.from <= s.to) ||
		(t.to >= s.from && t.to <= s.to)
}
