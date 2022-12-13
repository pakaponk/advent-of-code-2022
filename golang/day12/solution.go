package main

import (
	"fmt"
	"os"
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

func getHeightLevel(r rune) int {
	switch r {
	case 'S':
		return 0
	case 'E':
		return 26
	default:
		return int(r - 97)
	}
}

func toRune(s string) rune {
	return []rune(s)[0]
}

type position struct {
	x, y int
}

type node struct {
	pos    position
	height int
	next   []position
}

type answer struct {
	distance int
}

func getNextNodes(curr *node) []position {
	if curr.next == nil {
		x := curr.pos.x
		y := curr.pos.y

		right := position{x: x + 1, y: y}
		left := position{x: x - 1, y: y}
		top := position{x: x, y: y - 1}
		down := position{x: x, y: y + 1}

		var next []position
		if n, ok := nodeByPos[right]; ok && isReachable(curr, n) {
			next = append(next, right)
		}
		if n, ok := nodeByPos[left]; ok && isReachable(curr, n) {
			next = append(next, left)
		}
		if n, ok := nodeByPos[top]; ok && isReachable(curr, n) {
			next = append(next, top)
		}
		if n, ok := nodeByPos[down]; ok && isReachable(curr, n) {
			next = append(next, down)
		}

		curr.next = next
		return curr.next
	} else {
		return curr.next
	}
}

func isReachable(curr *node, next *node) bool {
	return curr.height-next.height <= 1
}

func bfs(curr *node, distance int, isTarget func(curr *node) bool) {

	if min, ok := minDisByPos[curr.pos]; !ok {
		minDisByPos[curr.pos] = distance
	} else {
		if min <= distance {
			// We have reached this point with fewer distance
			return
		} else {
			minDisByPos[curr.pos] = distance
		}
	}

	if isTarget(curr) {
		ans = append(ans, answer{distance: distance})
		return
	}

	next := getNextNodes(curr)
	if len(next) == 0 {
		return
	}

	// Coloring surrounding nodes before traversing
	for _, p := range next {
		if _, ok := visited[p]; !ok {
			if _, ok := minDisByPos[curr.pos]; !ok {
				minDisByPos[curr.pos] = distance
			}
		}
	}

	for _, p := range next {
		if _, ok := visited[p]; !ok {
			n := nodeByPos[p]

			visited[p] = true

			// copied := make([]pos, len(paths))
			// copy(copied, paths)
			// copied = append(copied, p)

			bfs(n, distance+1, isTarget)

			delete(visited, p)
		}
	}
}

var nodeByPos map[position]*node
var minDisByPos map[position]int
var visited map[position]bool
var ans []answer

func findShortestDistance(curr *node, isTarget func(curr *node) bool) int {
	visited = make(map[position]bool)
	minDisByPos = make(map[position]int)
	ans = nil

	bfs(curr, 0, isTarget)

	min := ans[0].distance
	for _, a := range ans {
		if min > a.distance {
			min = a.distance
		}
	}

	return min
}

func main() {
	txt := readInput("./day12/input.txt")

	lines := strings.Split(txt, "\n")

	nodeByPos = make(map[position]*node)

	var root position
	var part1Target node

	for i, l := range lines {
		heights := strings.Split(l, "")
		for j, h := range heights {
			pos := position{
				x: j,
				y: i,
			}

			n := node{
				pos:    pos,
				height: getHeightLevel(toRune(h)),
			}
			nodeByPos[pos] = &n

			switch h {
			case "S":
				part1Target = n
			case "E":
				root = pos
			default:
				// Do Nothing
			}
		}
	}

	isPart1Target := func(curr *node) bool {
		return curr.pos == part1Target.pos
	}
	isPart2Target := func(curr *node) bool {
		return curr.height == 0
	}

	curr := nodeByPos[root]

	fmt.Println("Shortest Distance from E -> S", findShortestDistance(curr, isPart1Target))
	fmt.Println("Shortest Distance from E -> any 'a' or 'S'", findShortestDistance(curr, isPart2Target))
}
