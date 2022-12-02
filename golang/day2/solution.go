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

func getChoiceScore(choice rune) int {
	ascii := int(choice)

	// X, Y, Z
	if ascii >= 88 {
		return ascii - 87
	} else {
		return ascii - 64
	}
}

func getMatchScore(a, b int) int {
	if a == b {
		return 3
	} else if b-a == 1 || b-a == -2 {
		return 6
	} else {
		return 0
	}
}

func getYourChoice(a, b int) int {
	if b == 2 {
		return a
	} else if b == 3 {
		if a == 3 {
			return 1
		} else {
			return a + 1
		}
	} else {
		if a == 1 {
			return 3
		} else {
			return a - 1
		}
	}
}

func main() {
	dat, err := os.ReadFile("./day2/input.txt")
	check(err)

	txt := string(dat)

	rows := strings.Split(txt, "\n")
	sum1 := 0
	sum2 := 0
	for _, row := range rows {
		choices := strings.Split(row, " ")

		recommendedChoice := choices[1]
		opponentChoice := choices[0]

		// Part 1
		a := getChoiceScore(rune(opponentChoice[0]))
		b := getChoiceScore(rune(recommendedChoice[0]))
		sum1 += getMatchScore(a, b) + b

		// Part 2
		// In Part 2, turn out the recommendedChoice is actually the recommended result
		// So we need to figure out our choice that would get us the recommended result
		c := getYourChoice(a, b)
		sum2 += getMatchScore(a, c) + c
	}

	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)
}
