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

func part(preserveOrder bool) error {
	lines := util.MustReadFile("input.txt")
	stacks := make([][]string, 100)
	var stackCount int
	for _, line := range lines {
		if strings.Contains(line, "[") {
			for j := 0; j < 9; j++ {
				c := line[j*4+1]
				if c >= 'A' && c <= 'Z' {
					stacks[j] = append([]string{string(c)}, stacks[j]...)
				}
				if j+1 > stackCount {
					stackCount = j + 1
				}
			}
			continue
		}

		if line == "" {
			stacks = stacks[:stackCount]
		}

		if strings.Contains(line, "move") {
			parts := strings.Fields(line)
			moveCount := util.MustParseInt(parts[1])
			from := util.MustParseInt(parts[3]) - 1
			to := util.MustParseInt(parts[5]) - 1
			if !preserveOrder {
				for i := 0; i < moveCount; i++ {
					stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
					stacks[from] = stacks[from][:len(stacks[from])-1]
				}
			} else {
				stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-moveCount:]...)
				stacks[from] = stacks[from][:len(stacks[from])-moveCount]
			}
		}
	}

	for i := 0; i < len(stacks); i++ {
		fmt.Print(stacks[i][len(stacks[i])-1])
	}
	fmt.Println()

	return nil
}
