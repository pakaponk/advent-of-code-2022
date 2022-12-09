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

type point struct {
	x, y int
}

type command struct {
	direction string
	amount    int
}

func follow(head, tail point) point {
	diffX := head.x - tail.x
	diffY := head.y - tail.y

	moveX := cal(head.x, tail.x)
	moveY := cal(head.y, tail.y)

	if diffX != 0 && diffY != 0 {
		if moveX != 0 {
			tail.x += moveX
			tail.y += unit(diffY)
		} else if moveY != 0 {
			tail.x += unit(diffX)
			tail.y += moveY
		}
	} else {
		tail.x += moveX
		tail.y += moveY
	}

	return tail
}

func unit(num int) int {
	if num > 0 {
		return 1
	} else if num < 0 {
		return -1
	} else {
		return 0
	}
}

func cal(a, b int) int {
	if a > b && a-b != 1 {
		return 1
	} else if a < b && a-b != -1 {
		return -1
	} else {
		return 0
	}
}

func move(p point, d string) point {
	switch d {
	case "R":
		p.x += 1
		break
	case "L":
		p.x -= 1
		break
	case "U":
		p.y -= 1
		break
	case "D":
		p.y += 1
		break
	}
	return p
}

func solve(commands []command, size int) int {
	tailIndex := size - 1
	knots := make([]point, size)

	visited := make(map[point]bool)
	visited[knots[tailIndex]] = true

	for _, c := range commands {
		direction := c.direction
		amount := c.amount

		for i := 0; i < amount; i++ {
			knots[0] = move(knots[0], direction)

			for i := 0; i < len(knots)-1; i++ {
				leader := knots[i]
				follower := knots[i+1]
				next := follow(leader, follower)

				if follower.x != next.x || follower.y != next.y {
					knots[i+1] = next
				}
			}

			visited[knots[tailIndex]] = true
		}
	}

	return len(visited)
}

func main() {
	txt := readInput("./day9/input.txt")

	lines := strings.Split(txt, "\n")

	commands := make([]command, len(lines))
	for i, l := range lines {
		inputs := strings.Split(l, " ")

		direction := inputs[0]
		amount, err := strconv.Atoi(inputs[1])
		check(err)

		commands[i] = command{
			direction: direction,
			amount:    amount,
		}
	}

	fmt.Println("Part 1:", solve(commands, 2))
	fmt.Println("Part 2:", solve(commands, 10))
}
