package main

import (
	"fmt"
	"strings"

	"github.com/neilgarb/advent2022/util"
)

func main() {
	util.MustDo(parts)
}

func parts() error {
	lines := util.MustReadFile("input.txt")
	var cur *dir
	allsizes := make(map[*dir]int)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			break
		}
		if i == 0 {
			cur = &dir{
				name: "",
			}
			continue
		}
		if line == "$ ls" {
			for i++; i < len(lines); i++ {
				line = lines[i]
				if strings.Contains(line, "$") {
					i--
					break
				}
				size, name, _ := strings.Cut(line, " ")
				if size == "dir" {
					cur.dirs = append(cur.dirs, &dir{
						parent: cur,
						name:   name,
					})
					continue
				}
				cur.files = append(cur.files, &file{
					name: name,
					size: util.MustParseInt(size),
				})
				c := cur
				for c != nil {
					c.size += cur.files[len(cur.files)-1].size
					allsizes[c] = c.size
					c = c.parent
				}
			}
			continue
		}
		if line == "$ cd .." {
			cur = cur.parent
			continue
		}
		if strings.HasPrefix(line, "$ cd ") {
			name := strings.TrimPrefix(line, "$ cd ")
			var found bool
			for _, d := range cur.dirs {
				if d.name == name {
					cur = d
					found = true
					break
				}
			}
			if !found {
				panic("dir not found")
			}
		}
	}

	var tot int
	for _, v := range allsizes {
		if v <= 100000 {
			tot += v
		}
	}

	fmt.Println(tot)

	for cur.parent != nil {
		cur = cur.parent
	}
	toDelete := 30000000 - (70000000 - cur.size)
	smallest := cur.size
	for _, v := range allsizes {
		if v >= toDelete && v < smallest {
			smallest = v
		}
	}

	fmt.Println(smallest)

	return nil
}

type dir struct {
	parent *dir
	name   string
	dirs   []*dir
	files  []*file
	size   int
}

type file struct {
	name string
	size int
}
