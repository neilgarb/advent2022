package util

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFile(name string) ([]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()
	scanner := bufio.NewScanner(f)
	res := make([]string, 0)
	for scanner.Scan() {
		l := scanner.Text()
		res = append(res, l)
	}
	return res, scanner.Err()
}

func MustReadFile(name string) []string {
	ll, err := ReadFile(name)
	if err != nil {
		panic(err)
	}
	return ll
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
