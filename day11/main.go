package main

import (
	"fmt"
	"github.com/vifraa/aoc20/util"
)

func main() {
	input, _ := util.ReadInputAsString("./day11/input.txt")

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}


func Part1(input []string) int {
	counter := 0
	for !util.EqualsStringSlice(input, modelSeatMovement(input)) {
		counter++
	}
	return counter
}

func Part2(input []string) int {
	return 0
}

func modelSeatMovement(input []string) []string {
	for x := 0; x < len(input); x++ {
		row := input[x]
		for y := 0; y < len(row); y++ {
			elem := row[y]
			switch elem {
			case 'L':
				if GetOccupiedAdjacentSeats(input, x, y) == 0 {
					input[x] = input[x][0:y] + "#" + input[x][y+1:]
				}
			case '#':
				if GetOccupiedAdjacentSeats(input, x, y) >= 4 {
					input[x] = input[x][0:y] + "L" + input[x][y+1:]
				}
			default:
				continue
			}
		}
	}

	return input
}


func GetOccupiedAdjacentSeats(seats []string, row,col int) int {
	directions := []struct {
		dx int
		dy int
	}{{-1,-1},{-1,0},{-1,1},{0,-1},{0,1},{1,-1},{1,0},{1,1}}

	found := 0
	for _, dir := range directions {
		x := row + dir.dx
		y := col + dir.dy

		if x < 0 || x >= len(seats) || y < 0 || y >= len(seats[0]) {
			continue
		}
		if seats[x][y] == '#' {
			found++
		}
	}

	return found
}
