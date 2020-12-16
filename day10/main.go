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
	oneDiff := 0
	twoDiff := 0
	threeDiff := 0

	sort.Ints(input)

	prevNum := start
	for _, i := range input {
		temp := i - prevNum

		if temp == 1 {
			oneDiff++
		} else if temp == 2 {
			twoDiff++
		} else if temp == 3 {
			threeDiff++
		} else {
			return 0
		}
		prevNum = i
	}

	threeDiff++

	return oneDiff * threeDiff
}

func Part2(input []int) int {


	return 0
}
