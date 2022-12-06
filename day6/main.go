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
	return part(4)
}

func part2() error {
	return part(14)
}

func part(l int) error {
	lines := util.MustReadFile("input.txt")
	line := lines[0]
	known := make(map[byte]int)
	for i := 0; i < len(line); i++ {
		c := line[i]
		known[c]++
		if i >= l {
			known[line[i-l]]--
			if known[line[i-l]] == 0 {
				delete(known, line[i-l])
			}
		}
		if len(known) == l {
			fmt.Println(i + 1)
			break
		}
	}
	return nil
}
