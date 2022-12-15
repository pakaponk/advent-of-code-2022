package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func readInput(path string) string {
	dat, err := os.ReadFile(path)
	check(err)
	return string(dat)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func toInt(s string) int {
	num, err := strconv.Atoi(s)
	check(err)

	return num
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

type pair struct {
	s, b point
}

type coverage struct {
	start, end int
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func getManhattanDistance(p1, p2 point) int {
	diffX := abs(p1.x - p2.x)
	diffY := abs(p1.y - p2.y)

	return diffX + diffY
}

// Example
// const Target = 11
// const UpperBound = 20

const Target = 2000000
const UpperBound = 4000000

func solvePart1(inputs []string) int {
	pairs := getPairs(inputs)

	occupied := make(map[int]bool)
	coverage := make(map[int]bool)

	for _, p := range pairs {
		s := p.s
		b := p.b

		md := getManhattanDistance(s, b)

		if s.y == Target {
			occupied[s.x] = true
		}

		if b.y == Target {
			occupied[b.x] = true
		}

		if s.y-md <= Target && Target <= s.y+md {
			d := md - abs(s.y-Target)

			for i := s.x - d; i <= s.x+d; i++ {
				if _, ok := occupied[i]; !ok {
					coverage[i] = true
				}
				occupied[i] = true
			}
		}
	}
	return len(coverage)
}

func solvePart2(inputs []string) int {
	pairs := getPairs(inputs)

	occupied := make(map[int][]coverage)
	for _, p := range pairs {
		s := p.s
		b := p.b

		md := getManhattanDistance(s, b)

		for i := s.y - md; i <= s.y+md; i++ {
			if i < 0 || i > UpperBound {
				continue
			}

			d := md - abs(s.y-i)

			occupied[i] = append(occupied[i], coverage{start: max(s.x-d, 0), end: min(s.x+d, UpperBound)})
		}
	}

	for i := 0; i <= UpperBound; i++ {
		sort.Slice(occupied[i], func(a int, z int) bool {
			return occupied[i][a].start < occupied[i][z].start
		})

		m := mergeCoverages(occupied[i])

		if len(m) > 1 {
			x := m[0].end + 1
			y := i
			// fmt.Println("X", x, "Y", y, "Result", (x*4000000)+y)
			return (x * 4000000) + y
		}
	}

	return -1
}

func mergeCoverages(coverages []coverage) []coverage {
	curr := coverages
	merged := true

	for merged {
		merged = false
		var next []coverage

		a := curr[0]
		for i := 1; i < len(curr); i++ {
			b := curr[i]

			if v, ok := merge(a, b); ok {
				a = v
				merged = true
			} else {
				next = append(next, a)
				a = b
			}
		}
		next = append(next, a)

		if merged {
			curr = next
		}
	}

	return curr
}

// Assume that a.start <= b.start
func merge(a, b coverage) (coverage, bool) {
	if a.start <= b.start && b.start <= (a.end+1) {
		return coverage{start: min(a.start, b.start), end: max(a.end, b.end)}, true
	} else {
		return coverage{}, false
	}
}

func getPairs(inputs []string) []pair {
	r, err := regexp.Compile(`Sensor at x=(?P<sx>-?\d+), y=(?P<sy>-?\d+): closest beacon is at x=(?P<bx>-?\d+), y=(?P<by>-?\d+)`)
	check(err)

	pairs := make([]pair, len(inputs))
	for i, input := range inputs {
		matches := r.FindStringSubmatch(input)
		sx := toInt(matches[1])
		sy := toInt(matches[2])
		bx := toInt(matches[3])
		by := toInt(matches[4])

		s := point{x: sx, y: sy}
		b := point{x: bx, y: by}

		pairs[i] = pair{s: s, b: b}
	}

	return pairs
}

func main() {
	txt := readInput("./day15/input.txt")

	inputs := strings.Split(txt, "\n")

	fmt.Println("Part 1:", solvePart1(inputs))
	fmt.Println("Part 2:", solvePart2(inputs))
}
