package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/diwasrimal/aoc-2024/utils"
)

func main() {
	part, file := utils.MustInitOptions()
	defer file.Close()

	sc := bufio.NewScanner(file)

	var reports [][]int
	for sc.Scan() {
		line := sc.Text()
		parts := strings.Split(line, " ")
		var levels []int
		for _, part := range parts {
			level, _ := strconv.Atoi(part)
			levels = append(levels, level)
		}
		reports = append(reports, levels)
	}

	if part == 1 {
		part1(reports)
	} else {
		part2(reports)
	}
}

func part1(reports [][]int) {
	safes := 0
	for _, levels := range reports {
		if unsafeIdx(levels) == -1 {
			safes++
		}
	}
	fmt.Println("part1", safes)
}

func part2(reports [][]int) {
	safes := 0
	for _, levels := range reports {
		unsafeAt := unsafeIdx(levels)
		if unsafeAt == -1 {
			safes++
			continue
		}
		replacement1, replacement2 := unsafeAt-1, unsafeAt
		change1Safe := unsafeIdx(deleteLevelAt(replacement1, levels)) == -1
		change2Safe := unsafeIdx(deleteLevelAt(replacement2, levels)) == -1
		if change1Safe || change2Safe {
			safes++
		}
	}
	fmt.Println("part2", safes)
}

// Returns the index of level that made record unsafe
// returns -1 if safe (no unsafe idx found)
func unsafeIdx(levels []int) int {
	if len(levels) <= 1 {
		return -1
	}
	increasing := levels[1] > levels[0]
	for i := 1; i < len(levels); i++ {
		diff := utils.Abs(levels[i] - levels[i-1])
		isSequential := (increasing && levels[i] > levels[i-1]) || (!increasing && levels[i] < levels[i-1])
		validDiff := 1 <= diff && diff <= 3
		if !isSequential || !validDiff {
			return i
		}
	}
	return -1
}

func deleteLevelAt(idx int, levels []int) []int {
	deleted := make([]int, len(levels)-1)
	copy(deleted[:idx], levels[:idx])
	copy(deleted[idx:], levels[idx+1:])
	return deleted
}
