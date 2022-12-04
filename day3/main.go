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

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func part1() error {
	lines := util.MustReadFile("input.txt")
	var tot int
	for _, line := range lines {
		seen := make(map[rune]bool)
		a, b := line[:len(line)/2], line[len(line)/2:]
		for _, c := range a {
			if strings.ContainsRune(b, c) && !seen[c] {
				tot += strings.IndexRune(alphabet, c) + 1
				seen[c] = true
			}
		}
	}
	fmt.Println(tot)
	return nil
}

func part2() error {
	lines := util.MustReadFile("input.txt")
	var tot int
	m := make([]map[rune]bool, 3)
	for i, line := range lines {
		m[i%3] = make(map[rune]bool)
		for _, c := range line {
			m[i%3][c] = true
		}
		if i%3 == 2 {
			for k := range m[0] {
				if m[1][k] && m[2][k] {
					tot += strings.IndexRune(alphabet, k) + 1
					break
				}
			}
		}
	}
	fmt.Println(tot)
	return nil
}
