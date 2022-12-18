package main

import (
	"fmt"

	"github.com/neilgarb/advent2022/util"
)

func main() {
	part1()
	part2()
}

func part1() {
	trees := parseTrees()
	dim := len(trees)
	tot := make(map[util.V2]bool)
	for i := 0; i < dim; i++ {
		tot[util.V2{i, 0}] = true
		tot[util.V2{i, dim - 1}] = true
		tot[util.V2{0, i}] = true
		tot[util.V2{dim - 1, i}] = true
		bigX, bigY := -1, -1
		for j := 0; j < dim; j++ {
			if trees[j][i] > bigX {
				tot[util.V2{j, i}] = true
				bigX = trees[j][i]
			}
			if trees[i][j] > bigY {
				tot[util.V2{i, j}] = true
				bigY = trees[i][j]
			}
		}
		bigX, bigY = -1, -1
		for j := dim - 1; j >= 0; j-- {
			if trees[j][i] > bigX {
				tot[util.V2{j, i}] = true
				bigX = trees[j][i]
			}
			if trees[i][j] > bigY {
				tot[util.V2{i, j}] = true
				bigY = trees[i][j]
			}
		}
	}
	fmt.Println(len(tot))
}

func part2() {
	trees := parseTrees()
	dim := len(trees)
	high := -1
	for y := 0; y < dim; y++ {
		for x := 0; x < dim; x++ {
			if s := score(trees, x, y); s > high {
				high = s
			}
		}
	}
	fmt.Println(high)
}

func score(trees [][]int, x, y int) int {
	t := trees[y][x]
	s := 1
	for i := x + 1; i < len(trees); i++ {
		if trees[y][i] >= t || i == len(trees)-1 {
			s *= i - x
			break
		}
	}
	for i := x - 1; i >= 0; i-- {
		if trees[y][i] >= t || i == 0 {
			s *= x - i
			break
		}
	}
	for i := y + 1; i < len(trees); i++ {
		if trees[i][x] >= t || i == len(trees)-1 {
			s *= i - y
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		if trees[i][x] >= t || i == 0 {
			s *= y - i
			break
		}
	}
	return s
}

func parseTrees() [][]int {
	lines := util.MustReadFile("input.txt")
	trees := make([][]int, 0)
	for _, line := range lines {
		var row []int
		for _, c := range line {
			row = append(row, util.MustParseInt(string(c)))
		}
		trees = append(trees, row)
	}
	return trees
}
