package main

import (
	"fmt"

	"github.com/neilgarb/advent2022/util"
)

func main() {
	util.MustDo(part1)
	util.MustDo(part2)
}

type coord struct {
	x, y int
}

func part1() error {
	trees := parseTrees()
	dim := len(trees)
	tot := make(map[coord]bool)
	for i := 0; i < dim; i++ {
		tot[coord{i, 0}] = true
		tot[coord{i, dim - 1}] = true
		tot[coord{0, i}] = true
		tot[coord{dim - 1, i}] = true
		bigx, bigy := -1, -1
		for j := 0; j < dim; j++ {
			if trees[j][i] > bigx {
				tot[coord{j, i}] = true
				bigx = trees[j][i]
			}
			if trees[i][j] > bigy {
				tot[coord{i, j}] = true
				bigy = trees[i][j]
			}
		}
		bigx, bigy = -1, -1
		for j := dim - 1; j >= 0; j-- {
			if trees[j][i] > bigx {
				tot[coord{j, i}] = true
				bigx = trees[j][i]
			}
			if trees[i][j] > bigy {
				tot[coord{i, j}] = true
				bigy = trees[i][j]
			}
		}
	}
	fmt.Println(len(tot))
	return nil
}

func part2() error {
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
	return nil
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
