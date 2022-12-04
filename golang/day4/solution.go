package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Interval struct {
	start int
	end   int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getInterval(interval string) Interval {
	temp := strings.Split(interval, "-")

	start, err := strconv.Atoi(temp[0])
	check(err)

	end, err := strconv.Atoi(temp[1])
	check(err)

	return Interval{start, end}
}

func isFullyContained(itvA, itvB Interval) bool {
	if itvA.start < itvB.start {
		// |----| A
		//   |--| B
		return itvA.end >= itvB.end
	} else if itvA.start > itvB.start {
		//   |--| A
		// |----| B
		return itvB.end >= itvA.end
	} else {
		// If it starts at the same point, one will always fully contain the another one.
		// |----| A/B
		// |--|   B/A
		return true
	}
}

func isIntersected(itvA, itvB Interval) bool {
	if itvA.start <= itvB.start {
		return itvA.end >= itvB.start
	} else {
		return itvB.end >= itvA.start
	}
}

func main() {
	dat, err := os.ReadFile("./day4/input.txt")
	check(err)

	txt := string(dat)

	pairs := strings.Split(txt, "\n")
	part1Ans := 0
	part2Ans := 0
	for _, pair := range pairs {
		periods := strings.Split(pair, ",")

		a := getInterval(periods[0])
		b := getInterval(periods[1])

		if isFullyContained(a, b) {
			part1Ans += 1
		}

		if isIntersected(a, b) {
			part2Ans += 1
		}
	}

	fmt.Println("Part 1:", part1Ans)
	fmt.Println("Part 2:", part2Ans)
}
