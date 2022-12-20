package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/neilgarb/advent2022/util"
)

func main() {
	part1()
	part2()
}

func part1() {
	blueprints := parseBlueprints("input.txt")
	var tot int
	for i, b := range blueprints {
		tot += (i + 1) * eval(b, map[string]int{"ore": 1}, map[string]int{}, map[string]map[string]int{}, 0)
	}
	fmt.Println(tot)
}

func part2() {
	blueprints := parseBlueprints("input.txt")
	tot := 1
	for i, b := range blueprints {
		tot *= eval(b, map[string]int{"ore": 1}, map[string]int{}, map[string]map[string]int{}, 0)
		if i == 2 {
			break
		}
	}
	fmt.Println(tot)
}

func eval(b blueprint, robots, resources map[string]int, best map[string]map[string]int, min int) int {
	if resources["clay"] > 80 {
		return 0
	}
	if resources["ore"] > 80 {
		return 0
	}
	if robots["ore"] > b["clay"]["ore"] && robots["ore"] > b["obsidian"]["ore"] {
		return 0
	}
	if robots["clay"] > b["obsidian"]["clay"] {
		return 0
	}
	if robots["obsidian"] > b["geode"]["obsidian"] {
		return 0
	}
	if min == 32 {
		return resources["geode"]
	}
	var max int
	for _, r := range resourceRanks {
		if canBuild(b[r], resources) {
			newResources := add(add(resources, neg(b[r])), robots)
			newRobots := add(robots, map[string]int{r: 1})
			key := makeKey(newRobots, min+1)
			if best[key] == nil || !allLess(newResources, best[key]) {
				best[key] = newResources
				max = util.Max(max, eval(
					b,
					newRobots,
					newResources,
					best,
					min+1,
				))
			}
			if r == "geode" {
				return max
			}
		}
	}
	newResources := add(resources, robots)
	key := makeKey(robots, min+1)
	if best[key] == nil || !allLess(newResources, best[key]) {
		best[key] = newResources
		max = util.Max(max, eval(
			b,
			robots,
			newResources,
			best,
			min+1,
		))
	}
	return max
}

func makeKey(robots map[string]int, min int) string {
	var vals []string
	for k, v := range robots {
		vals = append(vals, fmt.Sprintf("%s=%d", k, v))
	}
	sort.Strings(vals)
	return fmt.Sprintf("%d_%s", min, strings.Join(vals, "_"))
}

func allLess(m1, m2 map[string]int) bool {
	for k, v := range m1 {
		if v > m2[k] {
			return false
		}
	}

	return true
}

func canBuild(costs, resources map[string]int) bool {
	for k, v := range costs {
		if resources[k] < v {
			return false
		}
	}
	return true
}

func add(m1, m2 map[string]int) map[string]int {
	res := make(map[string]int)
	for k, v := range m1 {
		res[k] = v
	}
	for k, v := range m2 {
		res[k] += v
	}
	return res
}

func neg(m map[string]int) map[string]int {
	res := make(map[string]int)
	for k, v := range m {
		res[k] = -v
	}
	return res
}

var resourceRanks = []string{"geode", "obsidian", "clay", "ore"}

type blueprint map[string]map[string]int

var recipeRe = regexp.MustCompile(`Each (.*) robot costs (.*)`)
var costRe = regexp.MustCompile(`(\d+) (ore|clay|obsidian)`)

func parseBlueprints(file string) []blueprint {
	lines := util.MustReadFile(file)
	var res []blueprint
	for i, line := range lines {
		b := make(blueprint)
		line = strings.TrimPrefix(line, fmt.Sprintf("Blueprint %d: ", i+1))
		parts := strings.Split(line, ". ")
		for _, p := range parts {
			r := make(map[string]int)
			recipeMatches := recipeRe.FindStringSubmatch(p)
			costMatches := costRe.FindAllStringSubmatch(p, 2)
			for _, m := range costMatches {
				r[m[2]] = util.MustParseInt(m[1])
			}
			b[recipeMatches[1]] = r
		}
		res = append(res, b)
	}
	return res
}
