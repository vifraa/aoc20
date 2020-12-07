package main

import (
	"fmt"
	"github.com/vifraa/aoc20/util"
)

func main() {
	input, _ := util.ReadInputAsString("./day6/input.txt")

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	total := 0
	seenChars := make(map[int32]bool)
	for _, s := range input {
		if s == "" {
			total += len(seenChars)
			seenChars = make(map[int32]bool)
		}

		for _, ss := range s {
			if _, ok := seenChars[ss]; !ok {
				seenChars[ss] = true
			}
		}
	}

	total += len(seenChars)
	return total
}

func part2(input []string) int {
	total := 0
	seenChars := make(map[int32]int)
	peopleInGroup := 0

	for _, s := range input {
		if s == "" {
			for _, v := range seenChars {
				if v == peopleInGroup {
					total++
				}
			}
			seenChars = make(map[int32]int)
			peopleInGroup = 0
			continue
		}

		peopleInGroup++
		for _, ss := range s {
			seenChars[ss] += 1
		}
	}

	for _, v := range seenChars {
		if v == peopleInGroup {
			total++
		}
	}
	return total
}
