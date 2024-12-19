package utils

import (
	"flag"
	"fmt"
	"os"
)

type numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Abs[T numeric](val T) T {
	if val < 0 {
		return -val
	}
	return val
}

func Tern[T any](condition bool, a T, b T) T {
	if condition {
		return a
	}
	return b
}

func MustInitOptions() (int, *os.File) {
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

	return *part, file
}

func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
