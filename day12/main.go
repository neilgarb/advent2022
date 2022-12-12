package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/neilgarb/advent2022/util"
)

func main() {
	util.MustDo(part1)
	util.MustDo(part2)
}

type c struct {
	x, y int
}

func part1() error {
	lines := util.MustReadFile("input.txt")
	var start, end c
	for y, line := range lines {
		if x := strings.Index(line, "S"); x > -1 {
			start = c{x, y}
		}
		if x := strings.Index(line, "E"); x > -1 {
			end = c{x, y}
		}
	}

	path := astar(start, end, lines)
	fmt.Println(len(path) - 1) // Number of steps.
	return nil
}

func part2() error {
	lines := util.MustReadFile("input.txt")
	var end c
	var starts []c
	for y, line := range lines {
		for x, ch := range line {
			if ch == 'E' {
				end = c{x, y}
			} else if ch == 'a' || ch == 'S' {
				starts = append(starts, c{x, y})
			}
		}
	}
	lowest := math.MaxInt
	for _, start := range starts {
		path := astar(start, end, lines)
		if len(path) == 0 {
			continue
		}
		if len(path)-1 < lowest {
			lowest = len(path) - 1
		}
	}
	fmt.Println(lowest)
	return nil
}

func astar(start, end c, lines []string) []c {
	open := []c{start}
	openm := map[c]bool{start: true}
	from := make(map[c]c)
	g := make(map[c]int)
	g[start] = 0
	f := make(map[c]int)
	f[start] = dis(start, end)
	w := len(lines[0])
	for len(open) > 0 {
		var cur c
		lowest := math.MaxInt
		var lowesti int
		for i, t := range open {
			if f[t] < lowest {
				cur = t
				lowest = f[t]
				lowesti = i
			}
		}
		if cur == end {
			return makePath(from, cur)
		}
		open = append(open[:lowesti], open[lowesti+1:]...)
		delete(openm, cur)
		for _, dir := range []c{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			neigh := c{cur.x + dir.x, cur.y + dir.y}
			if neigh.x < 0 || neigh.x >= w || neigh.y < 0 || neigh.y >= len(lines) {
				continue
			}
			elCur := lines[cur.y][cur.x]
			if elCur == 'S' {
				elCur = 'a'
			} else if elCur == 'E' {
				elCur = 'z'
			}
			elNeigh := lines[neigh.y][neigh.x]
			if elNeigh == 'E' {
				elNeigh = 'z'
			} else if elNeigh == 'S' {
				elNeigh = 'a'
			}
			diff := int(elNeigh) - int(elCur)
			if diff > 1 {
				continue
			}
			tg := g[cur] + 1
			gneigh, ok := g[neigh]
			if !ok {
				gneigh = math.MaxInt
			}
			if tg < gneigh {
				from[neigh] = cur
				g[neigh] = tg
				f[neigh] = tg + dis(neigh, end)
				if !openm[neigh] {
					open = append(open, neigh)
					openm[neigh] = true
				}
			}
		}
	}

	return nil
}

func makePath(from map[c]c, cur c) []c {
	path := []c{cur}
	for {
		n, ok := from[cur]
		if !ok {
			break
		}
		path = append([]c{n}, path...)
		cur = n
	}
	return path
}

func dis(from, to c) int {
	return (to.x - from.x) + (to.y - from.y)
}
