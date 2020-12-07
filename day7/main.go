package main

import (
	"fmt"
	"github.com/vifraa/aoc20/util"
	"strconv"
	"strings"
)

func main() {
	input, _ := util.ReadInputAsString("./day7/input.txt")

	fmt.Printf("Part 1: %d\n", part1(input)) // Should be 259
	fmt.Printf("Part 2: %d\n", part2(input))
}


func part1(input []string) int {
	bagsMapped := mapBagContents(input)
	counter := 0
	for _, v := range bagsMapped {
		if contains(v, bagsMapped) {
			counter++
		}
	}
	return counter
}

func part2(input []string) int {
	bagsMapped := mapBagContents(input)
	return countInside(bagsMapped["shiny gold"], bagsMapped)
}


func mapBagContents(input []string) map[string]map[string]int {
	totalMapped := make(map[string]map[string]int)
	for _, s := range input {
		t := strings.Split(s, "bags contain")
		parent := strings.TrimSpace(t[0])

		bag := make(map[string]int)


		child := strings.TrimSpace(t[1])
		child = strings.ReplaceAll(child, "bags", "")
		child = strings.ReplaceAll(child, "bag", "")
		child = strings.ReplaceAll(child, ".", "")
		children := strings.Split(child, ",")

		for _, c := range children {
			c = strings.TrimSpace(c)
			cSplitted := strings.Split(c, " ")

			amount, err := strconv.Atoi(cSplitted[0])
			if err != nil {
				bag[strings.Join(cSplitted, " ")] = -1
			} else {
				bag[strings.Join(cSplitted[1:], " ")] = amount
			}
		}

		totalMapped[parent] = bag
	}
	return totalMapped
}


func countInside(current map[string]int, total map[string]map[string]int) int {
	totalCount := 0
	for k, v := range current {
		if k == "no other" {
			continue
		} else {
			totalCount += v * countInside(total[k], total) + v
		}
	}
	return totalCount
}

func contains(current map[string]int, total map[string]map[string]int) bool {
	found := make([]bool, 0)
	for k, _ := range current {
		if k == "shiny gold" {
			return true
		} else {
			found = append(found, contains(total[k], total))
		}
	}
	for _, b := range found {
		if b {
			return true
		}
	}
	return false
}
