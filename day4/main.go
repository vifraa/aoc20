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
	passports := parsePassports(input)
	validPassports := make([]map[string]string, 0)
	for _, p := range passports {
		if passportRequiredFieldsExist(p, requiredFields) {
			validPassports = append(validPassports, p)
		}
	}
	return len(validPassports)
}

func part2(input []string) int {
	passports := parsePassports(input)
	validPassports := make([]map[string]string, 0)
	for _, p := range passports {
		if passportIsValid(p, requiredFields) {
			validPassports = append(validPassports, p)
		}
	}
	return len(validPassports)
}

func parsePassports(input []string) (passports []map[string]string) {
	passports = make([]map[string]string, 0)

	currentPassport := make(map[string]string)
	for _, line := range input {
		if line == "" {
			passports = append(passports, currentPassport)
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
	if len(currentPassport) > 0 {
		passports = append(passports, currentPassport)
	}

	return passports
}

func passportRequiredFieldsExist(passport map[string]string, requiredFields []FieldValidator) bool {
	for _, f := range requiredFields {
		if _, ok := passport[f.Name()]; !ok {
			return false
		}
	}
	return true
}

func passportIsValid(passport map[string]string, requiredFields []FieldValidator) bool {
	for _, f := range requiredFields {
		if item, ok := passport[f.Name()]; !ok || !f.IsValid(item) {
			return false
		}
	}
	return true
}
