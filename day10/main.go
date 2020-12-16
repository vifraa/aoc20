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
	input = append(input, 0)
	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)

	steps := map[int]int{}

	for i := len(input) - 1; i >= 0; i-- {
		num := input[i]

		possibilities := 0
		if v, ok := steps[num+1]; ok && v > 0 {
			possibilities += v
		}
		if v, ok := steps[num+2]; ok && v > 0 {
			possibilities += v
		}
		if v, ok := steps[num+3]; ok && v > 0 {
			possibilities += v
		}
		if possibilities == 0 {
			possibilities++
		}

		steps[num] = possibilities
	}

	return steps[0]
}
