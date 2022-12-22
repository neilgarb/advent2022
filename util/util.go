package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func MustReadFile(name string) []string {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()
	scanner := bufio.NewScanner(f)
	res := make([]string, 0)
	for scanner.Scan() {
		l := scanner.Text()
		res = append(res, l)
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return res
}

func MustParseInt(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return i
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Sign(i int) int {
	if i < 0 {
		return -1
	}
	if i > 0 {
		return 1
	}
	return 0
}

type V2 struct {
	X, Y int
}

func (v V2) Add(w V2) V2 {
	return V2{v.X + w.X, v.Y + w.Y}
}

func (v V2) Neg() V2 {
	return V2{-v.X, -v.Y}
}

type V3 struct {
	X, Y, Z int
}

func (v V3) Add(w V3) V3 {
	return V3{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}
