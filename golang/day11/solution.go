package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
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

type monkey struct {
	items    []int
	divisor  int
	opr, opd string
	t1, t2   int
	ctr      int
}

func getNum(s string) int {
	r, err := regexp.Compile("\\d+")
	check(err)

	str := r.FindString(s)
	num, err := strconv.Atoi(str)
	check(err)

	return num
}

func getItems(s string) []int {
	r, err := regexp.Compile("\\d+")
	check(err)

	strs := r.FindAllString(s, -1)
	nums := make([]int, len(strs))
	for i, s := range strs {
		num, err := strconv.Atoi(s)
		check(err)

		nums[i] = num
	}

	return nums
}

func getFormula(s string) (operator, operand string) {
	r, err := regexp.Compile(`new = old (?P<Operator>[\+\-\*\/]) (?P<Operand>\d+|old)`)
	check(err)

	strs := r.FindStringSubmatch(s)

	operator = strs[1]
	operand = strs[2]

	return operator, operand
}

func getOperand(old int, operand string) int {
	if operand == "old" {
		return old
	} else {
		num, err := strconv.Atoi(operand)
		check(err)

		return num
	}
}

func calWorryLvl1(old, lcm int, operator, operand string) int {
	opd1 := old
	opd2 := getOperand(old, operand)

	var out int
	switch operator {
	case "+":
		out = (opd1 + opd2)
	case "-":
		out = (opd1 - opd2)
	case "*":
		out = (opd1 * opd2)
	case "/":
		out = (opd1 / opd2)
	default:
		out = opd1
	}

	return out / 3
}

func calWorryLvl2(old, lcm int, operator, operand string) int {
	opd1 := old
	opd2 := getOperand(old, operand)

	var out int
	switch operator {
	case "+":
		out = (opd1 + opd2)
	case "-":
		out = (opd1 - opd2)
	case "*":
		out = (opd1 * opd2)
	case "/":
		out = (opd1 / opd2)
	default:
		out = opd1
	}

	// Mod by multiple of all monkeys' divisor here to prevent the output from overflow int
	// Since all divisors are prime number, the multiple of them are LCM (least common multiple)
	// LCM works because `worry level % lcm % factor of lcm = worry level % factor of lcm`

	// Assume that divisors are 2, 3 so lcm is 6
	// N % 6 = r where 0 <= r < 6
	//   N   : 0 1 2 3 4 5 | 6 7 8 9 10 11 | 12 13 ..
	// N % r : 0 1 2 3 4 5 | 0 1 2 3  4  5 |  0  1 ..
	// N % 3 : 0 1 2|0 1 2 | 0 1 2 0  1  2 |  0  1 ..
	// N % 2 : 0 1|0 1|0 1 | 0 1 0 1  0  1 |  0  1 ..

	// From the above example
	// We can see that Range of N % 3 and Range of N % 2 fitted perfectly in the Range of N % r
	// So N % 3 = N % r % 3 and N % 2 = N % r % 2
	return out % lcm
}

func simulate(monkeys []monkey, rounds int, cal func(old, lcm int, operator, operand string) int) int {
	lcm := 1
	for _, m := range monkeys {
		lcm *= m.divisor
	}

	for i := 0; i < rounds; i++ {
		for j, monkey := range monkeys {
			thrownItems := make([]int, len(monkey.items))
			for k, item := range monkey.items {
				thrownItems[k] = cal(item, lcm, monkey.opr, monkey.opd)
			}

			for _, item := range thrownItems {
				if item%monkey.divisor == 0 {
					monkeys[monkey.t1].items = append(monkeys[monkey.t1].items, item)
				} else {
					monkeys[monkey.t2].items = append(monkeys[monkey.t2].items, item)
				}
			}

			monkeys[j].items = nil
			monkeys[j].ctr += len(thrownItems)
		}
	}

	sort.Slice(monkeys, func(a, z int) bool {
		return monkeys[a].ctr > monkeys[z].ctr
	})

	m1 := monkeys[0].ctr
	m2 := monkeys[1].ctr

	return m1 * m2
}

func getMonkeys(monkeyLines []string) []monkey {
	monkeys := make([]monkey, len(monkeyLines))
	for i, monkeyLine := range monkeyLines {
		lines := strings.Split(monkeyLine, "\n")

		divisor := getNum(lines[3])
		items := getItems(lines[1])
		opr, operand := getFormula(lines[2])
		t1 := getNum(lines[4])
		t2 := getNum(lines[5])

		monkeys[i] = monkey{
			items:   items,
			divisor: divisor,
			opr:     opr,
			opd:     operand,
			t1:      t1,
			t2:      t2,
		}
	}

	return monkeys
}

func main() {
	txt := readInput("./day11/input.txt")

	monkeyLines := strings.Split(txt, "\n\n")

	fmt.Println("Part 1:", simulate(getMonkeys(monkeyLines), 20, calWorryLvl1))
	fmt.Println("Part 2:", simulate(getMonkeys(monkeyLines), 10000, calWorryLvl2))
}
