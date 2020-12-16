package main

import (
	"fmt"
	"github.com/vifraa/aoc20/util"
	"sort"
)

func main() {
	input, _ := util.ReadInputAsInt("./day10/input.txt")

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}

func Part1(input []int) int {
	start := 0
	diff := make(map[int]int)

	sort.Ints(input)
	prevNum := start
	for _, i := range input {
		temp := i - prevNum

		diff[temp]++
		prevNum = i
	}

	diff[3]++

	return diff[1] * diff[3]
}

func Part2(input []int) int {

	return 0
}
