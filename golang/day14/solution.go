package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(path string) string {
	dat, err := os.ReadFile(path)
	check(err)
	return string(dat)
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	check(err)

	return n
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

type point struct {
	x, y int
}

type line struct {
	start point
	end   point
}

func getStartAndEnd(a, b point) (start, end point) {
	if a.y == b.y {
		if a.x < b.x {
			return a, b
		} else {
			return b, a
		}
	} else {
		if a.y < b.y {
			return a, b
		} else {
			return b, a
		}
	}
}

func getSymbol(v int) string {
	switch v {
	case 1:
		return "#"
	case 2:
		return "o"
	default:
		return "."
	}
}

func drawCave(minX, maxX, minY, maxY int) {
	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			fmt.Print(getSymbol(cave[i][j]))
		}
		fmt.Println()
	}
}

// Simulate the bottom for Part 2
func isHittingTheBottom(p point) bool {
	return p.y+1 < maxY+2
}

func isDownable(p point) bool {
	return isHittingTheBottom(p) && cave[p.y+1][p.x] == 0
}

func isLeftDownable(p point) bool {
	return isHittingTheBottom(p) && cave[p.y+1][p.x-1] == 0
}

func isRightDownable(p point) bool {
	return isHittingTheBottom(p) && cave[p.y+1][p.x+1] == 0
}

func fall(isNotFallIntoAbyss func(curr point) bool) bool {
	curr := point{x: 500, y: 0}

	// The start point is blocked
	if cave[curr.y][curr.x] == 2 {
		return false
	}

	for isNotFallIntoAbyss(curr) {
		if isDownable(curr) {
			curr.y = curr.y + 1
		} else if isLeftDownable(curr) {
			curr.x = curr.x - 1
			curr.y = curr.y + 1
		} else if isRightDownable(curr) {
			curr.x = curr.x + 1
			curr.y = curr.y + 1
		} else {
			cave[curr.y][curr.x] = 2

			minX = min(curr.x, minX)
			maxX = max(curr.x, maxX)

			return true
		}
	}
	return false
}

var cave [1000][1000]int
var minX = 1000
var maxX = 0
var maxY = 0

func getLines(inputs []string) []line {
	var lines []line

	for _, input := range inputs {
		coordinates := strings.Split(input, "->")

		var curr point
		for i, c := range coordinates {
			nums := strings.Split(c, ",")
			x := toInt(strings.TrimSpace(nums[0]))
			y := toInt(strings.TrimSpace(nums[1]))

			p := point{x: x, y: y}
			if i != 0 {
				start, end := getStartAndEnd(curr, p)
				lines = append(lines, line{start: start, end: end})
			}

			minX = min(x, minX)
			maxX = max(x, maxX)
			maxY = max(y, maxY)

			curr = p
		}
	}

	return lines
}

func buildCave(lines []line) [1000][1000]int {
	var cave [1000][1000]int

	for _, l := range lines {
		start := l.start
		end := l.end

		if start.x == end.x {
			for j := start.y; j <= end.y; j++ {
				cave[j][start.x] = 1
			}
		} else {
			for j := start.x; j <= end.x; j++ {
				cave[start.y][j] = 1
			}
		}
	}

	return cave
}

func solvePart1(lines []line) int {
	cave = buildCave(lines)

	// For Part 1
	// The sand fall into abyss when it passes the last line
	isNotFallIntoAbyss := func(curr point) bool {
		return curr.y <= maxY
	}

	var ctr int
	for fall(isNotFallIntoAbyss) {
		ctr++
	}

	return ctr
}

func solvePart2(lines []line) int {
	cave = buildCave(lines)

	// For Part 2
	// The sand never fall into abyss because there is the bottom floor at maxY + 2
	isNotFallIntoAbyss := func(curr point) bool {
		return true
	}

	var ctr int
	for fall(isNotFallIntoAbyss) {
		ctr++
	}

	return ctr
}
func main() {
	txt := readInput("./day14/input.txt")

	inputs := strings.Split(txt, "\n")

	lines := getLines(inputs)

	fmt.Println("Part 1:", solvePart1(lines))
	fmt.Println("Part 2:", solvePart2(lines))
}
