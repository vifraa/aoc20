package main

import (
	"fmt"
	"github.com/vifraa/aoc20/util"
	"strconv"
	"strings"
)

func main() {
	input, _ := util.ReadInputAsString("./day12/input.txt")

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}

func Part1(input []string) int {
	ship := Ship{
		90,
		0,
		0,
	}
	for _, s := range input {
		ship.Move(strings.TrimSpace(s))
	}

	return ship.Manhattan()
}

func Part2(input []string) int {
	return -1
}

type Ship struct {
	direction int
	x         int
	y         int
}

func (s *Ship) parseAction(action string) (command string, amount int) {
	command = string(action[0])
	amount, _ = strconv.Atoi(action[1:])
	return
}

func (s *Ship) Move(action string) {
	d, amount := s.parseAction(action)

	switch d {
	case "N":
		s.y += amount
	case "S":
		s.y -= amount
	case "E":
		s.x += amount
	case "W":
		s.x -= amount
	case "L":
		s.direction = moveDegrees(s.direction, -amount)
	case "R":
		s.direction = moveDegrees(s.direction, amount)
	case "F":
		if s.direction == 0 {
			s.Move("N" + strconv.Itoa(amount))
		} else if s.direction == 90 {
			s.Move("E" + strconv.Itoa(amount))
		} else if s.direction == 180 {
			s.Move("S" + strconv.Itoa(amount))
		} else if s.direction == 270 {
			s.Move("W" + strconv.Itoa(amount))
		} else {
			fmt.Println(s.direction)
		}
	}
}

func (s Ship) Manhattan() int {
	return abs(s.x) + abs(s.y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func moveDegrees(current, change int) int {
	t := current + change
	if t < 0 {
		t += 360
	}

	return t % 360
}
