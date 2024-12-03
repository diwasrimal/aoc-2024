package main

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/diwasrimal/aoc-2024/utils"
)

var list1, list2 []int
var list2freq = make(map[int]int)

func main() {
	part, file := utils.MustInitOptions()
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		first, second, _ := strings.Cut(line, "   ")
		val1, _ := strconv.Atoi(first)
		val2, _ := strconv.Atoi(second)
		list1 = append(list1, val1)
		list2 = append(list2, val2)
		list2freq[val2]++
	}

	if part == 1 {
		part1()
	} else {
		part2()
	}
}

func part1() {
	slices.Sort(list1)
	slices.Sort(list2)

	sum := 0
	for i := 0; i < len(list1); i++ {
		diff := list1[i] - list2[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}

	fmt.Println("part1", sum)
}

func part2() {
	sum := 0
	for _, val := range list1 {
		sum += val * list2freq[val]
	}
	fmt.Println("part2", sum)
}
