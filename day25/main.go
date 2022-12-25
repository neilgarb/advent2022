package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	. "github.com/neilgarb/advent2022/util"
)

func main() {
	part1()
}

func part1() {
	lines := MustReadFile("input.txt")
	var tot int
	for _, line := range lines {
		tot += snafu2dec(line)
		if line != dec2snafu(snafu2dec(line)) {
			fmt.Println(line, snafu2dec(line), dec2snafu(snafu2dec(line)))
			os.Exit(1)
		}
	}
	fmt.Println(tot, dec2snafu(tot), snafu2dec(dec2snafu(tot)))
}

func dec2snafu(in int) string {
	// First convert to quinary.
	var quin string
	for in > 0 {
		rem := in % 5
		quin = strconv.Itoa(rem) + quin
		in /= 5
	}
	// Then convert to snafu.
	var res string
	var carry int
	for i := len(quin) - 1; i >= 0; i-- {
		v, _ := strconv.Atoi(string(quin[i]))
		v += carry
		c := v % 5
		carry = v / 5
		if c == 4 {
			res = "-" + res
			carry += 1
		} else if c == 3 {
			res = "=" + res
			carry += 1
		} else {
			res = strconv.Itoa(c) + res
		}
	}
	if carry > 0 {
		res = strconv.Itoa(carry) + res
	}
	return res
}

var chars = map[byte]float64{
	'=': -2,
	'-': -1,
	'0': 0,
	'1': 1,
	'2': 2,
}

func snafu2dec(s string) int {
	var tot int
	for i := len(s) - 1; i >= 0; i-- {
		tot += int(math.Pow(5, float64(len(s)-i-1)) * chars[s[i]])
	}
	return tot
}
