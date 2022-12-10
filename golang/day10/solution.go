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

func isInRange(c, x int) bool {
	return x-1 <= c && c <= x+1
}

func draw(c, x int) string {
	if isInRange(c, x) {
		return "#"
	} else {
		return "."
	}
}

func main() {
	txt := readInput("./day10/input.txt")

	lines := strings.Split(txt, "\n")

	var ctr int

	xVals := make([]int, len(lines)*2)
	x := 1

	var messages []string

	for _, l := range lines {
		ops := strings.Split(l, " ")

		switch ops[0] {
		case "noop":
			messages = append(messages, draw(ctr%40, x))

			xVals[ctr] = x
			ctr++
			break
		case "addx":
			messages = append(messages, draw(ctr%40, x))

			xVals[ctr] = x
			ctr++

			messages = append(messages, draw(ctr%40, x))

			xVals[ctr] = x
			ctr++

			val, err := strconv.Atoi(ops[1])
			check(err)
			x += val
			break
		}
	}

	targets := [6]int{20, 60, 100, 140, 180, 220}
	var sum int
	for _, t := range targets {
		sum += t * xVals[t-1]
	}

	ans := "\n"
	for i := 0; i < 40*6; i++ {
		if i != 0 && i%40 == 0 {
			ans += "\n"
		}

		ans += " " + messages[i]
	}

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", ans)
}
