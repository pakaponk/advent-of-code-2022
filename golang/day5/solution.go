package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func parseDecimalInt(str string) int {
	num, err := strconv.Atoi(str)
	check(err)

	return num
}

func getProcedure(str string) (amount int, src int, des int) {
	r, _ := regexp.Compile("\\d+")

	nums := r.FindAllString(str, 3)

	amount = parseDecimalInt(nums[0])
	src = parseDecimalInt(nums[1]) - 1
	des = parseDecimalInt(nums[2]) - 1

	return
}

func readInput(path string) string {
	dat, err := os.ReadFile(path)
	check(err)

	return string(dat)
}

func pop(arr []string) (bool, string, []string) {
	if len(arr) == 0 {
		return false, "", arr
	}

	x, a := arr[len(arr)-1], arr[:len(arr)-1]

	return true, x, a
}

func moveVia9000(inputs []string, stacks [][]string) {
	for _, input := range inputs {
		amount, src, des := getProcedure(input)

		var popped string
		var isPopped bool
		for i := 0; i < amount; i++ {
			isPopped, popped, stacks[src] = pop(stacks[src])

			if isPopped {
				stacks[des] = append(stacks[des], popped)
			}

		}
	}
}

func moveVia9001(inputs []string, stacks [][]string) {
	for _, input := range inputs {
		amount, src, des := getProcedure(input)

		var popped string
		var isPopped bool
		temp := []string{}
		for i := 0; i < amount; i++ {
			isPopped, popped, stacks[src] = pop(stacks[src])
			if isPopped {
				temp = append(temp, popped)
			}
		}

		for i := len(temp) - 1; i >= 0; i-- {
			stacks[des] = append(stacks[des], temp[i])
		}
	}
}

func getTopOfAllStacks(stacks [][]string) string {
	ans := ""
	for _, stack := range stacks {
		isPopped, popped, _ := pop(stack)

		if isPopped {
			ans += popped
		}
	}

	return ans
}

func solvePart1(cratesText, procedureText string) string {
	crates := strings.Split(cratesText, "\n")
	stacks := make([][]string, len(crates))

	for i, crate := range crates {
		stack := strings.Split(crate, " ")
		stacks[i] = stack
	}

	procedureInputs := strings.Split(procedureText, "\n")
	moveVia9000(procedureInputs, stacks)

	return getTopOfAllStacks(stacks)
}

func solvePart2(cratesText, procedureText string) string {
	crates := strings.Split(cratesText, "\n")
	stacks := make([][]string, len(crates))

	for i, crate := range crates {
		stack := strings.Split(crate, " ")
		stacks[i] = stack
	}

	procedureInputs := strings.Split(procedureText, "\n")
	moveVia9001(procedureInputs, stacks)

	return getTopOfAllStacks(stacks)
}

func main() {
	procedureText := readInput("./day5/procedure.txt")
	cratesText := readInput("./day5/crates.txt")

	part1Ans := solvePart1(cratesText, procedureText)
	part2Ans := solvePart2(cratesText, procedureText)

	fmt.Println("Part 1:", part1Ans)
	fmt.Println("Part 2:", part2Ans)
}
