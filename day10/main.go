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

func part(render bool) error {
	lines := util.MustReadFile("input.txt")
	c := cpu{x: 1, render: render}
	for _, line := range lines {
		if line == "noop" {
			c.noop()
		} else {
			_, val, _ := strings.Cut(line, " ")
			vali := util.MustParseInt(val)
			c.addx(vali)
		}
	}
	if !render {
		fmt.Println(c.score)
	}
	return nil
}

type cpu struct {
	cyc    int
	x      int
	score  int
	render bool
}

func (c *cpu) noop() {
	c.cycle()
}

func (c *cpu) addx(i int) {
	c.cycle()
	c.cycle()
	c.x += i
}

func (c *cpu) cycle() {
	if c.render {
		if c.cyc%40 == 0 {
			fmt.Println()
		}
		if c.cyc%40 == c.x || c.cyc%40 == c.x-1 || c.cyc%40 == c.x+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}

	c.cyc++
	if (c.cyc-20)%40 == 0 {
		c.score += c.x * c.cyc
	}
}
