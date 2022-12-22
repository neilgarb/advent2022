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
	part(false)
}

func part2() {
	part(true)
}

func part(part2 bool) {
	board, cur, instructions := parse("input.txt")
	var dir int
	wrapSpots := make(map[util.V3]util.V3)
	for x := 50; x < 100; x++ {
		wrapSpots[util.V3{x, 0, 3}] = util.V3{0, 100 + x, 0} // 1->6
	}
	for x := 100; x < 150; x++ {
		wrapSpots[util.V3{x, 0, 3}] = util.V3{x - 100, 199, 3} // 2->6
	}
	for y := 0; y < 50; y++ {
		wrapSpots[util.V3{149, y, 0}] = util.V3{99, 149 - y, 2} // 2->4
	}
	for x := 100; x < 150; x++ {
		wrapSpots[util.V3{x, 49, 1}] = util.V3{99, x - 50, 2} // 2->3
	}
	for y := 50; y < 100; y++ {
		wrapSpots[util.V3{99, y, 0}] = util.V3{y + 50, 49, 3} // 3->2
	}
	for y := 0; y < 50; y++ {
		wrapSpots[util.V3{50, y, 2}] = util.V3{0, 149 - y, 0} // 1->5
	}
	for y := 50; y < 100; y++ {
		wrapSpots[util.V3{50, y, 2}] = util.V3{y - 50, 100, 1} // 3->5
	}
	for y := 100; y < 150; y++ {
		wrapSpots[util.V3{99, y, 0}] = util.V3{149, 49 - (y - 100), 2} // 4->2
	}
	for x := 50; x < 100; x++ {
		wrapSpots[util.V3{x, 149, 1}] = util.V3{49, x + 100, 2} // 4->6
	}
	for x := 0; x < 50; x++ {
		wrapSpots[util.V3{x, 100, 3}] = util.V3{50, x + 50, 0} // 5->3
	}
	for y := 100; y < 150; y++ {
		wrapSpots[util.V3{0, y, 2}] = util.V3{50, 49 - (y - 100), 0} // 5->1
	}
	for y := 150; y < 200; y++ {
		wrapSpots[util.V3{0, y, 2}] = util.V3{y - 100, 0, 1} // 6->1
	}
	for x := 0; x < 50; x++ {
		wrapSpots[util.V3{x, 199, 1}] = util.V3{100 + x, 0, 1} // 6->2
	}
	for y := 150; y < 200; y++ {
		wrapSpots[util.V3{49, y, 0}] = util.V3{y - 100, 149, 3} // 6->4
	}
	for _, inst := range instructions {
		if inst.move > 0 {
			for i := 0; i < inst.move; i++ {
				next := cur.Add(dirs[dir])
				if board[next] == 0 {
					if part2 {
						wrap := util.V3{cur.X, cur.Y, dir}
						nextV3 := wrapSpots[wrap]
						next = util.V2{nextV3.X, nextV3.Y}
						if board[next] != 2 {
							dir = nextV3.Z
						}
					} else {
						next = cur
						for board[next] > 0 {
							next = next.Add(dirs[dir].Neg())
						}
						next = next.Add(dirs[dir])
					}
				}
				if board[next] == 2 { // Rock.
					break
				} else if board[next] == 1 {
					cur = next
				}
			}
		} else {
			dir = (dir + inst.turn) % 4
			if dir < 0 {
				dir += 4
			}
		}
	}
	fmt.Println(1000*(cur.Y+1) + 4*(cur.X+1) + dir)
}

var dirs = []util.V2{
	{1, 0},  // Right
	{0, 1},  // Down
	{-1, 0}, // Left
	{0, -1}, // Up
}

type instruction struct {
	turn int // 0=nothing, 1=R, -1=L
	move int
}

func parse(file string) (map[util.V2]int, util.V2, []instruction) {
	board := make(map[util.V2]int)
	var instructions []instruction
	lines := util.MustReadFile(file)
	var boardDone, haveStart bool
	var start util.V2
	for y, line := range lines {
		if line == "" {
			boardDone = true
		} else if boardDone {
			var num string
			for _, c := range line {
				if c == 'L' {
					if num != "" {
						instructions = append(instructions, instruction{move: util.MustParseInt(num)})
						num = ""
					}
					instructions = append(instructions, instruction{turn: -1})
				} else if c == 'R' {
					if num != "" {
						instructions = append(instructions, instruction{move: util.MustParseInt(num)})
						num = ""
					}
					instructions = append(instructions, instruction{turn: 1})
				} else {
					num += string(c)
				}
			}
			if num != "" {
				instructions = append(instructions, instruction{move: util.MustParseInt(num)})
			}
		} else {
			for x := 0; x < len(line); x++ {
				if line[x] == '.' {
					if !haveStart {
						start = util.V2{x, y}
						haveStart = true
					}
					board[util.V2{x, y}] = 1
				} else if line[x] == '#' {
					board[util.V2{x, y}] = 2
				}
			}
		}
	}
	return board, start, instructions
}
