package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/neilgarb/advent2022/util"
)

func main() {
	util.MustDo(part1)
	util.MustDo(part2)
}

func part1() error {
	return part(20, true)
}

func part2() error {
	return part(10_000, false)
}

func part(rounds int, divthree bool) error {
	monkeys := parse("input.txt")
	mod := 1
	for _, m := range monkeys {
		mod *= m.testdiv
	}
	for round := 0; round < rounds; round++ {
		for _, mon := range monkeys {
			for _, item := range mon.items {
				worry := mon.operate(item)
				if divthree {
					worry /= 3
				} else {
					worry %= mod
				}
				newmon := mon.test(worry)
				monkeys[newmon].items = append(monkeys[newmon].items, worry)
				mon.activity++
			}
			mon.items = nil
		}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].activity > monkeys[j].activity
	})
	fmt.Println(monkeys[0].activity * monkeys[1].activity)
	return nil
}

func parse(file string) []*monkey {
	lines := util.MustReadFile(file)
	cur := new(monkey)
	var monkeys []*monkey
	for _, line := range lines {
		if line == "" {
			monkeys = append(monkeys, cur)
			cur = new(monkey)
			continue
		}
		if strings.Contains(line, "Starting") {
			_, list, _ := strings.Cut(line, ":")
			for _, i := range strings.Split(list, ",") {
				cur.items = append(cur.items, util.MustParseInt(i))
			}
			continue
		}
		if strings.Contains(line, "Operation") {
			parts := strings.Fields(line)
			cur.op = parts[len(parts)-2]
			cur.operand = parts[len(parts)-1]
			continue
		}
		if strings.Contains(line, "Test") {
			parts := strings.Fields(line)
			cur.testdiv = util.MustParseInt(parts[len(parts)-1])
			continue
		}
		if strings.Contains(line, "true") {
			parts := strings.Fields(line)
			cur.truemon = util.MustParseInt(parts[len(parts)-1])
			continue
		}
		if strings.Contains(line, "false") {
			parts := strings.Fields(line)
			cur.falsemon = util.MustParseInt(parts[len(parts)-1])
			continue
		}
	}
	return append(monkeys, cur)
}

type monkey struct {
	items    []int
	op       string
	operand  string
	testdiv  int
	truemon  int
	falsemon int
	activity int
}

func (m *monkey) operate(i int) int {
	operand := i
	if m.operand != "old" {
		operand = util.MustParseInt(m.operand)
	}
	if m.op == "*" {
		return i * operand
	}
	return i + operand
}

func (m *monkey) test(i int) int {
	if i%m.testdiv != 0 {
		return m.falsemon
	}
	return m.truemon
}
