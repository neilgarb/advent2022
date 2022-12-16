package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/neilgarb/advent2022/util"
)

func main() {
	util.MustDo(part1)
	util.MustDo(part2)
}

var rexp = regexp.MustCompile(`Valve (.*) has flow rate=(.*); tunnels? leads? to valves? (.*)`)

func part1() error {
	lines := util.MustReadFile("input.txt")
	rooms, cur := parseRooms(lines)

	// Pre-compute distances.
	dists := make(map[*room]map[*room]int)
	for _, i := range rooms {
		dists[i] = make(map[*room]int)
		for _, j := range rooms {
			dists[i][j] = dist(i, j)
		}
	}

	// Traverse.
	var max int
	traverse(rooms, dists, cur, nil, 0, 30, &max)
	fmt.Println(max)

	return nil
}

func part2() error {
	lines := util.MustReadFile("input.txt")
	rooms, cur := parseRooms(lines)

	// Pre-compute distances.
	dists := make(map[*room]map[*room]int)
	for _, i := range rooms {
		dists[i] = make(map[*room]int)
		for _, j := range rooms {
			dists[i][j] = dist(i, j)
		}
	}
	//fmt.Println(dists)

	// Traverse.
	var max int
	for _, c := range makecombos(rooms) {
		var max1, max2 int
		traverse(c, dists, cur, nil, 0, 26, &max1)
		cmap := make(map[*room]bool)
		for _, r := range c {
			cmap[r] = true
		}
		var otherrooms []*room
		for _, r := range rooms {
			if !cmap[r] && r.flow > 0 {
				otherrooms = append(otherrooms, r)
			}
		}
		traverse(otherrooms, dists, cur, nil, 0, 26, &max2)
		//fmt.Println(c, otherrooms, max1, max2, max1+max2)
		max = util.Max(max, max1+max2)
	}
	fmt.Println(max)

	return nil
}

type room struct {
	name string
	flow int
	next []*room
}

func (r *room) String() string {
	return r.name
}

func dist(from, to *room) int {
	dists := map[*room]int{from: 0}
	set := []*room{from}
	for len(set) > 0 {
		r := set[0]
		set = set[1:]
		if r == to {
			continue
		}
		for _, n := range r.next {
			_, ok := dists[n]
			if !ok || dists[r]+1 < dists[n] {
				dists[n] = dists[r] + 1
				set = append(set, n)
			}
		}
	}
	return dists[to]
}

func traverse(
	rooms []*room, dists map[*room]map[*room]int, cur *room,
	seen map[*room]int, min, max int, maxScore *int,
) {
	for _, k := range rooms {
		if k.flow > 0 && seen[k] == 0 && dists[cur][k] > 0 {
			seen2 := make(map[*room]int)
			for sk, v := range seen {
				seen2[sk] = v
			}
			d := min + 1 + dists[cur][k]
			if d >= max {
				continue
			}
			seen2[k] = seen2[cur] + (max-d)*k.flow
			if seen2[k] > *maxScore {
				*maxScore = seen2[k]
			}
			traverse(rooms, dists, k, seen2, d, max, maxScore)
		}
	}
}

func parseRooms(lines []string) ([]*room, *room) {
	rooms := make(map[string]*room)
	var cur *room
	for _, line := range lines {
		matches := rexp.FindStringSubmatch(line)
		name := matches[1]
		flow := util.MustParseInt(matches[2])
		next := strings.Split(matches[3], ", ")
		r, ok := rooms[name]
		if ok {
			r.flow = flow
		} else {
			r = &room{name: name, flow: flow}
			rooms[name] = r
		}
		if name == "AA" {
			cur = r
		}
		for _, n := range next {
			if _, ok := rooms[n]; !ok {
				rooms[n] = &room{name: n}
			}
			r.next = append(r.next, rooms[n])
		}
	}
	var roomsList []*room
	for _, v := range rooms {
		roomsList = append(roomsList, v)
	}
	return roomsList, cur
}

func makecombos(rooms []*room) [][]*room {
	var newrooms []*room
	for _, r := range rooms {
		if r.flow > 0 {
			newrooms = append(newrooms, r)
		}
	}
	var res [][]*room
	for i := 0; i < int(math.Pow(2, float64(len(newrooms)))); i++ {
		var combo []*room
		for j, ch := range fmt.Sprintf("%0"+strconv.Itoa(len(newrooms))+"b", i) {
			if ch == '1' {
				combo = append(combo, newrooms[j])
			}
		}
		if len(combo) < 5 {
			continue
		}
		res = append(res, combo)
	}
	return res
}
