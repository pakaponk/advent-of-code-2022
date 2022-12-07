package main

import (
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func solve(input string, size int) int {
	for i := 0; i < len(input)-size-1; i++ {

		m := make(map[rune]int)
		isUniq := true
		for _, c := range input[i : i+size] {
			if m[c] == 0 {
				m[c] = 1
			} else {
				isUniq = false
			}
		}

		if isUniq {
			return i + size
		}
	}

	return -1
}

func main() {
	dat, err := os.ReadFile("./day6/input.txt")
	check(err)

	input := string(dat)

	part1Ans := solve(input, 4)
	part2Ans := solve(input, 14)

	fmt.Println("Part 1:", part1Ans)
	fmt.Println("Part 2:", part2Ans)
}
