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

func getSize(sizeByDir map[string]int, dirListMap map[string][]string, path string) int {
	var sum int

	for _, line := range dirListMap[path] {
		// `line` is either
		// "dir <dir_name>" or
		// "<size> <file_name>"
		parts := strings.Split(line, " ")

		if parts[0] == "dir" {
			name := parts[1]
			if size, ok := sizeByDir[path]; ok {
				sum += size
			} else {
				var childPath string
				if path == "/" {
					childPath = path + name
				} else {
					childPath = path + "/" + name
				}
				getSize(sizeByDir, dirListMap, childPath)
				sum += sizeByDir[childPath]
			}
		} else {
			size, err := strconv.Atoi(parts[0])
			check(err)

			sum += size
		}
	}

	sizeByDir[path] = sum

	return sum
}

func getSizeByDir(dirListMap map[string][]string) map[string]int {
	sizeByDir := make(map[string]int)
	getSize(sizeByDir, dirListMap, "/")
	return sizeByDir
}

func pop(arr []string) (bool, string, []string) {
	if len(arr) == 0 {
		return false, "", arr
	}

	x, a := arr[len(arr)-1], arr[:len(arr)-1]

	return true, x, a
}

func isCommandLine(line string) bool {
	return line[:1] == "$"
}

func getCommand(commandLine string) string {
	return commandLine[2:4]
}

// Part 1
const Threshold = 100000

// Part 2
const UsableSize = 70000000
const RequiredSizeForUpdate = 30000000

func main() {
	txt := readInput("./day7/input.txt")

	lines := strings.Split(txt, "\n")

	var dirs []string
	var wdStack []string
	var path string
	dirListMap := make(map[string][]string)
	for _, line := range lines {
		// `line` is either
		// 1) "$ cd <dir_name>" or
		// 2) "$ ls" or
		// 3) "dir <dir_name>"
		// 4) "<size> <file_name>"
		if isCommandLine(line) {
			c := getCommand(line)
			if c == "cd" {
				dir := line[5:]

				if dir == ".." {
					_, _, wdStack = pop(wdStack)
				} else {
					wdStack = append(wdStack, dir)
				}
			} else {
				path = wdStack[0] + strings.Join(wdStack[1:], "/")
				dirs = append(dirs, path)
			}
		} else {
			dirListMap[path] = append(dirListMap[path], line)
		}
	}

	sizeByDir := getSizeByDir(dirListMap)

	var part1Ans int
	for _, dir := range dirs {
		if sizeByDir[dir] <= Threshold {

			part1Ans += sizeByDir[dir]
		}
	}

	totalSize := sizeByDir["/"]
	unusedSize := UsableSize - totalSize
	diff := RequiredSizeForUpdate - unusedSize

	min := totalSize
	for _, dir := range dirs {
		if sizeByDir[dir] > diff && min > sizeByDir[dir] {
			min = sizeByDir[dir]

		}
	}

	fmt.Println("Part 1", part1Ans)
	fmt.Println("Part 2", min)
}
