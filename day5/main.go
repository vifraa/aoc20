package main

import (
	"fmt"
	"github.com/vifraa/aoc20/util"
	"math"
	"sort"
)

func main() {
	input, _ := util.ReadInputAsString("./day5/input.txt")

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	seats := findSeatIds(input)
	return seats[len(seats)-1]
}

func part2(input []string) int {
	seats := findSeatIds(input)

	prev := seats[0]
	for _, x := range seats {
		if x-prev == 2 {
			return x - 1
		}
		prev = x
	}
	return -1
}

func findSeatIds(input []string) []int {
	result := make([]int, 0, len(input))

	for _, i := range input {
		row := FindRow(i[:7])
		col := FindCol(i[7:])

		seatId := row*8 + col
		result = append(result, seatId)
	}
	sort.Ints(result)
	return result
}

func FindRow(input string) int {
	var low float64 = 0
	var high float64 = 127

	for _, c := range input {
		switch c {
		case 'F':
			high = math.Floor((high + low) / 2)
		case 'B':
			low = math.Ceil((high + low) / 2)
		}
	}
	return int(low)
}

func FindCol(input string) int {
	var low float64 = 0
	var high float64 = 7

	for _, c := range input {
		switch c {
		case 'L':
			high = math.Floor((high + low) / 2)
		case 'R':
			low = math.Ceil((high + low) / 2)
		}
	}
	return int(low)
}
