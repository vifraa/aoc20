package main

import (
	"regexp"
	"strconv"
)

var (
	birthYear      = BirthYear{"byr"}
	issueYear      = IssueYear{"iyr"}
	expirationYear = ExpirationYear{"eyr"}
	height         = Height{"hgt"}
	hairColor      = HairColor{"hcl"}
	eyeColor       = EyeColor{"ecl"}
	passportID     = PassportId{"pid"}
)

var requiredFields = []FieldValidator{birthYear, issueYear, expirationYear, height, hairColor, eyeColor, passportID}

type FieldValidator interface {
	IsValid(f string) bool
	Name() string
}

type BirthYear struct {
	name string
}

func (br BirthYear) Name() string {
	return br.name
}

func (_ BirthYear) IsValid(f string) bool {
	if len(f) != 4 {
		return false
	}

	fc, err := strconv.Atoi(f)
	if err != nil {
		return false
	}

	return fc >= 1920 && fc <= 2002
}

type IssueYear struct {
	name string
}

func (iy IssueYear) Name() string {
	return iy.name
}

func (_ IssueYear) IsValid(f string) bool {
	if len(f) != 4 {
		return false
	}

	fc, err := strconv.Atoi(f)
	if err != nil {
		return false
	}

	return fc >= 2010 && fc <= 2020
}

type ExpirationYear struct {
	name string
}

func (ey ExpirationYear) Name() string {
	return ey.name
}

func (_ ExpirationYear) IsValid(f string) bool {
	if len(f) != 4 {
		return false
	}

	fc, err := strconv.Atoi(f)
	if err != nil {
		return false
	}

	return fc >= 2020 && fc <= 2030
}

type Height struct {
	name string
}

func (h Height) Name() string {
	return h.name
}

func (_ Height) IsValid(f string) bool {
	cmValid, _ := regexp.Match("^1[5-8][0-9]cm|19[0-3]cm$", []byte(f))
	inValid, _ := regexp.Match("^59in|6\\din|7[0-6]in$", []byte(f))
	return cmValid || inValid
}

type HairColor struct {
	name string
}

func (hc HairColor) Name() string {
	return hc.name
}

func (_ HairColor) IsValid(f string) bool {
	valid, _ := regexp.Match("^#[0-9a-f]{6}$", []byte(f))
	return valid
}

type EyeColor struct {
	name string
}

func (ec EyeColor) Name() string {
	return ec.name
}

func (_ EyeColor) IsValid(f string) bool {
	switch f {
	case "amb":
		return true
	case "blu":
		return true
	case "brn":
		return true
	case "gry":
		return true
	case "grn":
		return true
	case "hzl":
		return true
	case "oth":
		return true
	default:
		return false
	}
}

type PassportId struct {
	name string
}

func (pID PassportId) Name() string {
	return pID.name
}

func (_ PassportId) IsValid(f string) bool {
	valid, _ := regexp.Match("^[0-9]{8}[1-9]$", []byte(f))
	return valid
}
