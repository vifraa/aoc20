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
	for true {
		tmp := modelSeatMovement(input)

		if util.EqualsStringSlice(input, tmp) {
			break
		}

		counter++
		input = tmp
	}

	return countOccupiedSeats(input)
}

func Part2(input []string) int {
	for true {
		tmp := modelSeatMovementPart2(input)

		if util.EqualsStringSlice(input, tmp) {
			break
		}
		input = tmp
	}

	return countOccupiedSeats(input)
}

func modelSeatMovement(input []string) []string {
	tmp := make([]string, len(input))
	copy(tmp, input)

	for x := 0; x < len(input); x++ {
		row := input[x]
		for y := 0; y < len(row); y++ {
			elem := row[y]
			switch elem {
			case 'L':
				if GetOccupiedAdjacentSeats(input, x, y) == 0 {
					tmp[x] = tmp[x][0:y] + "#" + tmp[x][y+1:]
				}
			case '#':
				if GetOccupiedAdjacentSeats(input, x, y) >= 4 {
					tmp[x] = tmp[x][0:y] + "L" + tmp[x][y+1:]
				}
			default:
				continue
			}
		}
	}

	return tmp
}

func modelSeatMovementPart2(input []string) []string {
	tmp := make([]string, len(input))
	copy(tmp, input)

	for x := 0; x < len(input); x++ {
		row := input[x]
		for y := 0; y < len(row); y++ {
			elem := row[y]
			switch elem {
			case 'L':
				if GetOccupiedVisibleSeats(input, x, y) == 0 {
					tmp[x] = tmp[x][0:y] + "#" + tmp[x][y+1:]
				}
			case '#':
				if GetOccupiedVisibleSeats(input, x, y) >= 5 {
					tmp[x] = tmp[x][0:y] + "L" + tmp[x][y+1:]
				}
			default:
				continue
			}
		}
	}

	return tmp
}

func GetOccupiedAdjacentSeats(seats []string, row, col int) int {
	directions := []struct {
		dx int
		dy int
	}{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

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

func GetOccupiedVisibleSeats(seats []string, row, col int) int {
	directions := []struct {
		dx int
		dy int
	}{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	found := 0
	for _, dir := range directions {
		x := row + dir.dx
		y := col + dir.dy

		for !(x < 0 || x >= len(seats) || y < 0 || y >= len(seats[0])) {
			if seats[x][y] == '#' {
				found++
				break
			} else if seats[x][y] == 'L' {
				break
			}
			x += dir.dx
			y += dir.dy
		}
	}

	return found
}

func countOccupiedSeats(seats []string) int {
	res := 0
	for _, s := range seats {
		for _, ss := range s {
			if ss == '#' {
				res++
			}
		}
	}
	return res
}
