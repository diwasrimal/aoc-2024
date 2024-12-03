package main

import (
	"bufio"

	"github.com/diwasrimal/aoc-2024/utils"
)

func main() {
	part, file := utils.MustInitOptions()
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		_ = line
	}

	if part == 1 {
		part1()
	} else {
		part2()
	}
}

func part1() {
}

func part2() {
}
