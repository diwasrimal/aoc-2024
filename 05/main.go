package main

import (
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/diwasrimal/aoc-2024/utils"
)

var (
	befores = make(map[int][]int)
	afters  = make(map[int][]int)
)

func main() {
	part, file := utils.MustInitOptions()
	defer file.Close()

	data := string(utils.Must(io.ReadAll(file)))
	rules, ordering, _ := strings.Cut(data, "\n\n")

	for _, line := range strings.Split(rules, "\n") {
		var a, b int
		utils.Must(fmt.Sscanf(line, "%d|%d", &a, &b))
		befores[b] = append(befores[b], a)
		afters[a] = append(afters[a], b)
	}

	var orderings [][]int
	for _, line := range strings.Split(strings.TrimSpace(ordering), "\n") {
		parts := strings.Split(line, ",")
		nums := make([]int, len(parts))
		for i, s := range parts {
			nums[i] = utils.Must(strconv.Atoi(s))
		}
		orderings = append(orderings, nums)
	}

	if part == 1 {
		part1(orderings)
	} else {
		part2(orderings)
	}
}

func part1(orderings [][]int) {
	sum := 0

	for _, nums := range orderings {
		i1, i2 := orderingMistakes(nums)
		if i1 == -1 && i2 == -1 {
			sum += nums[len(nums)/2]
		}
	}
	fmt.Println("part1", sum)
}

func part2(orderings [][]int) {
	sum := 0

	for _, nums := range orderings {
		i1, i2 := orderingMistakes(nums)
		if !(i1 == -1 && i2 == -1) {
			// Keep swapping util ordering becomes right
			for {
				nums[i1], nums[i2] = nums[i2], nums[i1]
				i1, i2 = orderingMistakes(nums)
				if i1 == -1 && i2 == -1 {
					break
				}
			}
			sum += nums[len(nums)/2]
		}
	}
	fmt.Println("part2", sum)
}

func orderingMistakes(nums []int) (int, int) {
	for i := 0; i < len(nums); i++ {
		beforelist, afterlist := befores[nums[i]], afters[nums[i]]
		for r := i + 1; r <= len(nums)-1; r++ {
			if slices.Contains(beforelist, nums[r]) {
				return i, r
			}
		}
		for l := i - 1; l >= 0; l-- {
			if slices.Contains(afterlist, nums[l]) {
				return i, l
			}
		}
	}
	return -1, -1
}
