package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/neilgarb/advent2022/util"
)

func main() {
	part1()
	part2()
}

func part1() {
	lines := util.MustReadFile("input.txt")
	var start, end util.V2
	for y, line := range lines {
		if x := strings.Index(line, "S"); x > -1 {
			start = util.V2{x, y}
		}
		if x := strings.Index(line, "E"); x > -1 {
			end = util.V2{x, y}
		}
	}

	path := astar(start, end, lines)
	fmt.Println(len(path) - 1) // Number of steps.
}

func part2() error {
	lines := util.MustReadFile("input.txt")
	var end util.V2
	var starts []util.V2
	for y, line := range lines {
		for x, ch := range line {
			if ch == 'E' {
				end = util.V2{x, y}
			} else if ch == 'a' || ch == 'S' {
				starts = append(starts, util.V2{x, y})
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

func astar(start, end util.V2, lines []string) []util.V2 {
	open := []util.V2{start}
	openm := map[util.V2]bool{start: true}
	from := make(map[util.V2]util.V2)
	g := make(map[util.V2]int)
	g[start] = 0
	f := make(map[util.V2]int)
	f[start] = dis(start, end)
	w := len(lines[0])
	for len(open) > 0 {
		var cur util.V2
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
		for _, dir := range []util.V2{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			neigh := cur.Add(dir)
			if neigh.X < 0 || neigh.X >= w || neigh.Y < 0 || neigh.Y >= len(lines) {
				continue
			}
			elCur := lines[cur.Y][cur.X]
			if elCur == 'S' {
				elCur = 'a'
			} else if elCur == 'E' {
				elCur = 'z'
			}
			elNeigh := lines[neigh.Y][neigh.X]
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

func makePath(from map[util.V2]util.V2, cur util.V2) []util.V2 {
	path := []util.V2{cur}
	for {
		n, ok := from[cur]
		if !ok {
			break
		}
		path = append([]util.V2{n}, path...)
		cur = n
	}
	return path
}

func dis(from, to util.V2) int {
	return (to.X - from.X) + (to.Y - from.Y)
}
