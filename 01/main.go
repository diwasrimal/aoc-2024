package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var list1, list2 []int
var list2freq = make(map[int]int)

func main() {
	part := flag.Int("part", 0, "part of the problem (1 or 2)")
	input := flag.String("input", "", "input file")
	flag.Parse()

	if (*part != 1 && *part != 2) || *input == "" {
		flag.Usage()
		os.Exit(1)
	}

	file, err := os.Open(*input)
	if err != nil {
		fmt.Printf("Error opening %s: %s\n", *input, err)
		os.Exit(1)
	}
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

	if *part == 1 {
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
