package main

import (
	"fmt"
	"github.com/vifraa/aoc20/util"
	"log"
	"strings"
)

func main() {
	input, err := util.ReadInputAsString("./day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	valid, _ := parsePassports(input)
	return len(valid)
}

func part2(input []string) int {
	valid, _ := parsePassports(input)
	return len(valid)
}

func parsePassports(input []string) (validPassports, invalidPassports []map[string]string) {
	validPassports = make([]map[string]string, 0)
	invalidPassports = make([]map[string]string, 0)

	currentPassport := make(map[string]string)
	for _, line := range input {
		if line == "" {
			if passportIsValid(currentPassport, requiredFields) {
				validPassports = append(validPassports, currentPassport)
			} else {
				invalidPassports = append(invalidPassports, currentPassport)
			}

			currentPassport = make(map[string]string)
			continue
		}

		valuePairs := strings.Split(line, " ")
		for _, p := range valuePairs {
			sp := strings.Split(p, ":")
			key := sp[0]
			value := sp[1]
			currentPassport[key] = value
		}
	}
	if passportIsValid(currentPassport, requiredFields) {
		validPassports = append(validPassports, currentPassport)
	} else {
		invalidPassports = append(invalidPassports, currentPassport)
	}
	return validPassports, invalidPassports
}

func passportIsValid(passport map[string]string, requiredFields []FieldValidator) bool {
	for _, f := range requiredFields {
		if item, ok := passport[f.Name()]; !ok || !f.IsValid(item) {
			return false
		}
	}
	return true
}
