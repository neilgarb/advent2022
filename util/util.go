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

func MustDo(fn func() error) {
	if err := fn(); err != nil {
		panic(err)
	}
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
