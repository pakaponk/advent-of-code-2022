package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dat, err := os.ReadFile("./day1/input.txt")
	check(err)

	txt := string(dat)

	elves := strings.Split(txt, "\n\n")
	totalElves := len(elves)
	totalWeights := make([]int, totalElves)
	for i := 0; i < totalElves; i++ {
		weightTexts := strings.Split(elves[i], "\n")
		sum := 0
		for j := 0; j < len(weightTexts); j++ {
			weight, err := strconv.Atoi(weightTexts[j])
			check(err)

			sum += weight
		}

		totalWeights[i] = sum
	}

	sort.Ints(totalWeights[:])

	fmt.Printf("Part 1: %d\n", totalWeights[totalElves-1])
	fmt.Printf("Part 2: %d\n", totalWeights[totalElves-1]+totalWeights[totalElves-2]+totalWeights[totalElves-3])
}
