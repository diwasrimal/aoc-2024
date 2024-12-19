package main

import (
	"bufio"
	"fmt"

	"github.com/diwasrimal/aoc-2024/utils"
)

func main() {
	part, file := utils.MustInitOptions()
	defer file.Close()

	sc := bufio.NewScanner(file)
	var lines []string
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}

	if part == 1 {
		part1(lines)
	} else {
		part2(lines)
	}
}

func part1(lines []string) {
	// 0--1--2
	// 3-----4
	// 5--6--7
	vectors := [8][2]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	count := 0

	isvalid := func(x, y int) bool {
		return 0 <= x && x < len(lines[0]) && 0 <= y && y < len(lines)
	}

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			for _, v := range vectors {
				found := true
				xnow, ynow := x, y
				for i := 0; i < 4; i++ {
					if !isvalid(xnow, ynow) || lines[ynow][xnow] != "XMAS"[i] {
						found = false
						break
					}
					xnow += v[0]
					ynow += v[1]
				}
				if found {
					count++
				}
			}
		}
	}

	fmt.Println("part1", count)
}

func part2(lines []string) {
	vectors := [4][2]int{
		{-1, -1}, //  0------1
		{1, -1},  //  |      |
		{1, 1},   //  |      |
		{-1, 1},  //  3------2
	}
	count := 0
	isvalid := func(x, y int) bool {
		return 0 <= x && x < len(lines[0]) && 0 <= y && y < len(lines)
	}

	for y := 0; y < len(lines); y++ {
	next:
		for x := 0; x < len(lines[0]); x++ {
			if lines[y][x] != 'A' {
				continue
			}
			s := ""
			for _, vec := range vectors {
				xnow, ynow := x+vec[0], y+vec[1]
				if !isvalid(xnow, ynow) {
					continue next
				}
				s += string(lines[y+vec[1]][x+vec[0]])
			}
			a := string(s[0]) + string(s[2])
			b := string(s[1]) + string(s[3])

			if (a == "MS" || a == "SM") && (b == "MS" || b == "SM") {
				count++
			}
		}
	}
	fmt.Println("part2", count)
}
