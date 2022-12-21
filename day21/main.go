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
	monkeys := parse("input.txt")
	fmt.Println(eval(monkeys, "root"))
}

func part2() {
	monkeys := parse("input.txt")
	r := monkeys["root"]
	var humn int
	if haveHumn(monkeys, r.left) {
		right := eval(monkeys, r.right)
		eval2(monkeys, r.left, right, &humn)
	} else {
		left := eval(monkeys, r.left)
		eval2(monkeys, r.right, left, &humn)
	}
	fmt.Println(humn)
}

func eval2(monkeys map[string]monkey, root string, exp int, humn *int) {
	m := monkeys[root]
	if root == "humn" {
		*humn = exp
		return
	}
	if m.operation == "" {
		if exp != m.operand {
			fmt.Println(":/", exp, m.operand)
		}
		return
	}
	if haveHumn(monkeys, m.left) {
		right := eval(monkeys, m.right)
		switch m.operation {
		case "*":
			eval2(monkeys, m.left, exp/right, humn)
		case "/":
			eval2(monkeys, m.left, exp*right, humn)
		case "+":
			eval2(monkeys, m.left, exp-right, humn)
		case "-":
			eval2(monkeys, m.left, exp+right, humn)
		}
	} else {
		left := eval(monkeys, m.left)
		switch m.operation {
		case "*":
			eval2(monkeys, m.right, exp/left, humn)
		case "/":
			eval2(monkeys, m.right, left/exp, humn)
		case "+":
			eval2(monkeys, m.right, exp-left, humn)
		case "-":
			eval2(monkeys, m.right, left-exp, humn)
		}
	}
}

func eval(monkeys map[string]monkey, root string) int {
	m := monkeys[root]
	if m.operation == "" {
		return m.operand
	}
	left := eval(monkeys, m.left)
	right := eval(monkeys, m.right)
	var v int
	switch m.operation {
	case "*":
		v = left * right
	case "/":
		v = left / right
	case "+":
		v = left + right
	case "-":
		v = left - right
	}
	return v
}

func haveHumn(monkeys map[string]monkey, root string) bool {
	if root == "humn" {
		return true
	}
	m := monkeys[root]
	if m.operation == "" {
		return false
	}
	return haveHumn(monkeys, m.left) || haveHumn(monkeys, m.right)
}

type monkey struct {
	operand     int
	operation   string
	left, right string
}

func parse(file string) map[string]monkey {
	lines := util.MustReadFile(file)
	res := make(map[string]monkey)
	for _, line := range lines {
		name, rest, _ := strings.Cut(line, ": ")
		parts := strings.Split(rest, " ")
		var m monkey
		if len(parts) == 1 {
			m.operand = util.MustParseInt(parts[0])
		} else {
			m.left = parts[0]
			m.right = parts[2]
			m.operation = parts[1]
		}
		res[name] = m
	}
	return res
}
