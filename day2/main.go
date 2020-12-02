package main

import (
	"fmt"
	"github.com/vifraa/aoc20/util"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := util.ReadInputAsString("./day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1 := 0
	part2 := 0
	for _, i := range input {
		if isValidPartOne(i) {
			part1++
		}
		if isValidPartTwo(i) {
			part2++
		}
	}

	fmt.Printf("Part1: %d \n", part1)
	fmt.Printf("Part2: %d \n", part2)
}

func splitInput(input string) (int, int, string, string) {
	s := strings.Split(input, " ")

	numbers := strings.Split(s[0], "-")
	low, _ := strconv.Atoi(numbers[0])
	high, _ := strconv.Atoi(numbers[1])

	character := string(s[1][0])
	password := s[2]

	return low, high, character, password
}

func isValidPartOne(input string) bool {
	low, high, c, password := splitInput(input)

	seenAmount := 0
	for _, char := range password {
		if string(char) == c {
			seenAmount++
		}
	}

	return seenAmount >= low && seenAmount <= high
}

func isValidPartTwo(input string) bool {
	low, high, c, password := splitInput(input)

	lowChar := password[low-1]
	highCar := password[high-1]

	return (string(lowChar) == c || string(highCar) == c) && string(lowChar) != string(highCar)
}
