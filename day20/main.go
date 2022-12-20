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
	part(1, 1)
}

func part2() {
	part(811589153, 10)
}

func part(mul, mixes int) {
	lines := util.MustReadFile("input.txt")
	var cur, zero *node
	ordering := make(map[int]*node)
	for i := 0; i < len(lines); i++ {
		n := node{val: util.MustParseInt(lines[i]) * mul}
		if cur != nil {
			cur.next = &n
			n.prev = cur
		}
		cur = &n
		ordering[len(ordering)] = &n
		if n.val == 0 {
			zero = &n
		}
	}
	ordering[len(ordering)-1].next = ordering[0]
	ordering[0].prev = ordering[len(ordering)-1]
	for m := 0; m < mixes; m++ {
		for i := 0; i < len(lines); i++ {
			n := ordering[i]
			if n.val > 0 {
				nv := n.val % (len(lines) - 1)
				for j := 0; j < nv; j++ {
					t := n.next
					n.next = t.next
					t.prev = n.prev
					t.next.prev = n
					n.prev.next = t
					n.prev = t
					t.next = n
				}
			} else if n.val < 0 {
				nv := n.val % (len(lines) - 1)
				for j := 0; j > nv; j-- {
					t := n.prev
					n.prev = t.prev
					t.next = n.next
					t.prev.next = n
					n.next.prev = t
					n.next = t
					t.prev = n
				}
			}
		}
	}
	var tot int
	cur = zero
	for i := 0; i <= 3000; i++ {
		if i%1000 == 0 {
			tot += cur.val
		}
		cur = cur.next
	}
	fmt.Println(tot)
}

type node struct {
	val  int
	next *node
	prev *node
}
