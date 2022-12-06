package util

import (
	"bufio"
	"os"
	"strconv"
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
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}