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

func getMap(str string) map[rune]bool {
	m := make(map[rune]bool)
	for _, c := range str {
		m[c] = true
	}

	return m
}

func getPriority(c rune) int {
	if c >= 97 {
		// a - z
		return int(c) - 96
	} else {
		// A - Z
		return int(c) - 64 + 26
	}
}

func solvePart1(bags []string) int {
	totalBags := len(bags)

	priorities := make([]int, totalBags)
	for i, bag := range bags {
		size := len(bag)

		first := bag[0 : size/2]
		second := bag[size/2:]

		m1 := getMap(first)
		m2 := getMap(second)

		var same rune
		for _, c := range first {
			if m1[c] && m2[c] {
				same = c
			}
		}

		priorities[i] = getPriority(same)
	}

	sum := 0
	for _, num := range priorities {
		sum += num
	}

	return sum
}

const TotalElvesInGroup = 3

func solvePart2(bags []string) int {
	totalBags := len(bags)

	priorities := make([]int, totalBags/3)
	group := make([]string, TotalElvesInGroup)
	for i, bag := range bags {
		group[i%TotalElvesInGroup] = bag

		if i%TotalElvesInGroup == TotalElvesInGroup-1 {
			first := group[0]
			second := group[1]
			third := group[2]

			m1 := getMap(first)
			m2 := getMap(second)
			m3 := getMap(third)

			var same rune
			for _, c := range first {
				if m1[c] && m2[c] && m3[c] {
					same = c
				}
			}

			priorities[i/TotalElvesInGroup] = getPriority(same)
		}

	}

	sum := 0
	for _, num := range priorities {
		sum += num
	}

	return sum
}

func main() {
	dat, err := os.ReadFile("./day3/input.txt")
	check(err)

	txt := string(dat)

	bags := strings.Split(txt, "\n")

	fmt.Println("Part 1:", solvePart1(bags))
	fmt.Println("Part 2:", solvePart2(bags))
}
