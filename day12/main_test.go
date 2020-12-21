package main

import "testing"

var input = []string{
	"F10",
	"N3 ",
	"F7 ",
	"R90",
	"F11",
}

func TestPart1(t *testing.T) {
	expected := 25
	actual := Part1(input)
	if expected != actual {
		t.Errorf("Expected: %d, got: %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := -1
	actual := Part2(input)
	if expected != actual {
		t.Errorf("Expected: %d, got: %d", expected, actual)
	}
}
