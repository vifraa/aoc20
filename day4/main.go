package main

import (
	"fmt"
	"github.com/vifraa/aoc20/util"
	"log"
	"strings"
)

const (
	birthYear      = "byr"
	issueYear      = "iyr"
	expirationYear = "eyr"
	height         = "hgt"
	hairColor      = "hcl"
	eyeColor       = "ecl"
	passportID     = "pid"
	countryID      = "cid"
)

var requiredFields = []string{birthYear, issueYear, expirationYear, height, hairColor, eyeColor, passportID}

func main() {
	input, err := util.ReadInputAsString("./day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	validPassports := make([]map[string]string, 0)
	invalidPassports := make([]map[string]string, 0)

	currentPassport := make(map[string]string, 0)
	for _, line := range input {
		if line == "" {
			if passportIsValid(currentPassport, requiredFields) {
				validPassports = append(validPassports, currentPassport)
			} else {
				invalidPassports = append(invalidPassports, currentPassport)
			}

			currentPassport = make(map[string]string, 0)
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

	fmt.Printf("Valid passports: %d", len(validPassports))
}

func passportIsValid(passport map[string]string, requiredFields []string) bool {
	for _, f := range requiredFields {
		if _, ok := passport[f]; !ok {
			return false
		}
	}
	return true
}
