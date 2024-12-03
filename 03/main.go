package main

import (
	"fmt"
	"io"
	"regexp"
	"strconv"

	"github.com/diwasrimal/aoc-2024/utils"
)

func main() {
	part, file := utils.MustInitOptions()
	defer file.Close()
	data, _ := io.ReadAll(file)
	input := string(data)

	if part == 1 {
		part1(input)
	} else {
		part2(input)
	}
}

func part1(input string) {
	rx := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := rx.FindAllStringSubmatch(input, -1)
	sum := 0
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += num1 * num2
	}
	fmt.Println("part1", sum)
}

func part2(input string) {
	rx := regexp.MustCompile(`(?:(do(?:n't)?\(\))|mul\((\d{1,3}),(\d{1,3})\))`)
	matches := rx.FindAllStringSubmatch(input, -1)
	sum := 0
	do := true
	for _, match := range matches {
		if match[0] == "do()" {
			do = true
		} else if match[0] == "don't()" {
			do = false
		} else if do {
			num1, _ := strconv.Atoi(match[2])
			num2, _ := strconv.Atoi(match[3])
			sum += num1 * num2
		}
	}
	fmt.Println("part2", sum)
}
