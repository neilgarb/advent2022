package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/neilgarb/advent2022/util"
)

func main() {
	part1()
	part2()
}

func part1() {
	lines := util.MustReadFile("input.txt")
	var tot int
	for i := range lines {
		if i%3 != 1 {
			continue
		}
		_, left := parse(lines[i-1])
		_, right := parse(lines[i])
		if cmp(left, right) <= 0 {
			tot += i/3 + 1
		}
	}
	fmt.Println(tot)
}

func part2() {
	lines := util.MustReadFile("input.txt")
	_, p1 := parse("[[2]]")
	_, p2 := parse("[[6]]")
	parsed := []itemlist{p1, p2}
	for _, line := range lines {
		if line == "" {
			continue
		}
		_, p := parse(line)
		parsed = append(parsed, p)
	}
	sort.Slice(parsed, func(i, j int) bool {
		return cmp(parsed[i], parsed[j]) < 0
	})
	tot := 1
	for i, p := range parsed {
		if p.String() == "[[2]]" || p.String() == "[[6]]" {
			tot *= i + 1
		}
	}
	fmt.Println(tot)
}

type itemlist []*item

func (i itemlist) String() string {
	var res []string
	for _, ii := range i {
		res = append(res, ii.String())
	}
	return "[" + strings.Join(res, ",") + "]"
}

type item struct {
	isval bool
	val   int
	vals  itemlist
}

func (i *item) String() string {
	if i.isval {
		return strconv.Itoa(i.val)
	}
	return i.vals.String()
}

func parse(s string) (int, itemlist) {
	var curv string
	var res []*item
	for i := 1; i < len(s); i++ {
		if digit(s[i]) {
			curv += string(s[i])
		} else if s[i] == ',' {
			if curv != "" {
				res = append(res, &item{isval: true, val: util.MustParseInt(curv)})
				curv = ""
			}
		} else if s[i] == '[' {
			n, vals := parse(s[i:])
			res = append(res, &item{vals: vals})
			i += n + 1
		} else if s[i] == ']' {
			if curv != "" {
				res = append(res, &item{isval: true, val: util.MustParseInt(curv)})
			}
			return i - 1, res
		}
	}
	return 0, res
}

func digit(b byte) bool {
	return b >= '0' && b <= '9'
}

func cmp(left, right itemlist) int {
	if len(left) > len(right) {
		return -cmp(right, left)
	}
	for i := range left {
		if left[i].isval && right[i].isval {
			if left[i].val < right[i].val {
				return -1
			}
			if left[i].val > right[i].val {
				return 1
			}
			continue
		}
		var c int
		if left[i].isval && !right[i].isval {
			c = cmp(itemlist{left[i]}, right[i].vals)
		} else if !left[i].isval && right[i].isval {
			c = cmp(left[i].vals, itemlist{right[i]})
		} else {
			c = cmp(left[i].vals, right[i].vals)
		}
		if c != 0 {
			return c
		}
	}
	if len(left) == len(right) {
		return 0
	}
	return -1
}
