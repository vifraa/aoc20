package main

import "testing"

var input = []string{
	"L.LL.LL.LL",
	"LLLLLLL.LL",
	"L.L.L..L..",
	"LLLL.LL.LL",
	"L.LL.LL.LL",
	"L.LLLLL.LL",
	"..L.L.....",
	"LLLLLLLLLL",
	"L.LLLLLL.L",
	"L.LLLLL.LL",
}

func TestPart1(t *testing.T) {
	expected := 37
	actual := Part1(input)
	if expected != actual {
		t.Errorf("Expected: %d, Got: %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 0
	actual := Part2(input)
	if expected != actual {
		t.Errorf("Expected: %d, Got: %d", expected, actual)
	}
}

func TestGetOccupiedAdjacentSeats(t *testing.T) {
	expected := 5
	actual := GetOccupiedAdjacentSeats(input, 3, 4)
	if expected != actual {
		t.Errorf("Expected: %d, Got: %d", expected, actual)
	}
}
