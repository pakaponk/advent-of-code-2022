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

type point struct {
	x, y int
}

func isVisibleFromLeft(trees [][]int, curr point) bool {
	for i := curr.x - 1; i >= 0; i-- {
		if trees[curr.y][i] >= trees[curr.y][curr.x] {
			return false
		}
	}
	return true
}

func isVisibleFromRight(trees [][]int, curr point) bool {
	for i := curr.x + 1; i < len(trees[0]); i++ {
		if trees[curr.y][i] >= trees[curr.y][curr.x] {
			return false
		}
	}
	return true
}

func isVisibleFromTop(trees [][]int, curr point) bool {
	for i := curr.y - 1; i >= 0; i-- {
		if trees[i][curr.x] >= trees[curr.y][curr.x] {
			return false
		}
	}
	return true
}

func isVisibleFromBottom(trees [][]int, curr point) bool {
	for i := curr.y + 1; i < len(trees); i++ {
		if trees[i][curr.x] >= trees[curr.y][curr.x] {
			return false
		}
	}
	return true
}

func getViewDistanceFromLeft(trees [][]int, curr point) int {
	var ctr int
	for i := curr.x - 1; i >= 0; i-- {
		ctr++
		if trees[curr.y][i] >= trees[curr.y][curr.x] {
			break
		}
	}
	return ctr
}

func getViewDistanceFromRight(trees [][]int, curr point) int {
	var ctr int
	for i := curr.x + 1; i < len(trees[0]); i++ {
		ctr++
		if trees[curr.y][i] >= trees[curr.y][curr.x] {
			break
		}
	}
	return ctr
}

func getViewDistanceFromTop(trees [][]int, curr point) int {
	var ctr int
	for i := curr.y - 1; i >= 0; i-- {
		ctr++
		if trees[i][curr.x] >= trees[curr.y][curr.x] {
			break
		}
	}
	return ctr
}

func getViewDistanceFromBottom(trees [][]int, curr point) int {
	var ctr int
	for i := curr.y + 1; i < len(trees); i++ {
		ctr++
		if trees[i][curr.x] >= trees[curr.y][curr.x] {
			break
		}
	}
	return ctr
}

func getScenicScore(trees [][]int, curr point) int {
	return getViewDistanceFromLeft(trees, curr) * getViewDistanceFromRight(trees, curr) * getViewDistanceFromTop(trees, curr) * getViewDistanceFromBottom(trees, curr)
}

func readInput(path string) string {
	dat, err := os.ReadFile(path)
	check(err)

	return string(dat)
}

func solvePart1(trees [][]int, rows, cols int) int {
	ctr := rows*2 + cols*2 - 4
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			curr := point{
				x: i,
				y: j,
			}
			if isVisibleFromLeft(trees, curr) || isVisibleFromRight(trees, curr) || isVisibleFromTop(trees, curr) || isVisibleFromBottom(trees, curr) {
				ctr++
			}
		}
	}

	return ctr
}

func solvePart2(trees [][]int, rows, cols int) int {
	var max int
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			curr := point{
				x: i,
				y: j,
			}
			score := getScenicScore(trees, curr)
			if max < score {
				max = score
			}
		}
	}

	return max
}

func main() {
	txt := readInput("./day8/input.txt")

	lines := strings.Split(txt, "\n")

	rows := len(lines)
	cols := len(lines[0])

	matrix := make([][]int, rows)

	for i, line := range lines {
		trees := strings.Split(line, "")
		matrix[i] = make([]int, cols)
		for j, t := range trees {
			h, err := strconv.Atoi(t)
			check(err)

			matrix[i][j] = h
		}
	}

	ctr := rows*2 + cols*2 - 4
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			curr := point{
				x: i,
				y: j,
			}
			if isVisibleFromLeft(matrix, curr) || isVisibleFromRight(matrix, curr) || isVisibleFromTop(matrix, curr) || isVisibleFromBottom(matrix, curr) {
				ctr++
			}
		}
	}

	fmt.Println("Part 1:", solvePart1(matrix, rows, cols))
	fmt.Println("Part 2:", solvePart2(matrix, rows, cols))
}
