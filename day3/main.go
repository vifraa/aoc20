package main

import (
	"fmt"
	"github.com/vifraa/aoc20/util"
	"log"
)

func main() {
	input, err := util.ReadInputAsString("./day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}


	fmt.Printf("Part 1: %d\n", treesOnPath(input, 3, 1))


	i1 := treesOnPath(input, 1, 1)
	i2 := treesOnPath(input, 3, 1)
	i3 := treesOnPath(input, 5, 1)
	i4 := treesOnPath(input, 7, 1)
	i5 := treesOnPath(input, 1, 2)

	fmt.Printf("Part 2: %d\n", i1 * i2 * i3 * i4 * i5)
}


func treesOnPath(input []string, stepRight, stepDown int) int {
	var amountOfTrees int
	col := stepRight

	for i := stepDown; i < len(input); i+=stepDown {
		row := input[i]

		if string(row[col]) == "#" {
			amountOfTrees++
		}

		col = (col + stepRight) % len(row)
	}

	return amountOfTrees
}
