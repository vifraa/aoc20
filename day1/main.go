package main

import (
	"fmt"
	"github.com/vifraa/aoc20/util"
	"log"
)

func main() {
	input, err := util.ReadInputAsInt("./day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part1: %d \n", part1(input))
	fmt.Printf("Part2: %d \n", part2(input))
}

func part1(input []int) int {
	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {
			if input[i]+input[j] == 2020 {
				return input[i] * input[j]
			}

		}
	}
	return 0
}

func part2(input []int) int {
	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {
			for k := j; k < len(input); k++ {
				if input[i]+input[j]+input[k] == 2020 {
					return input[i] * input[j] * input[k]
				}
			}
		}
	}
	return 0
}
