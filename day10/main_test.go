package main

import "testing"

var testInput = []int{
	28,
	33,
	18,
	42,
	31,
	14,
	46,
	20,
	48,
	47,
	24,
	23,
	49,
	45,
	19,
	38,
	39,
	11,
	1,
	32,
	25,
	35,
	8,
	17,
	7,
	9,
	4,
	2,
	34,
	10,
	3,
}

func TestPart1(t *testing.T) {
	expected := 220
	actual := Part1(testInput)

	if expected != actual {
		t.Errorf("Expected: %d, Got: %d", expected, actual)
	}
}


func TestPart2(t *testing.T) {
	expected := 19208
	actual := Part2(testInput)

	if expected != actual {
		t.Errorf("Expected: %d, Got: %d", expected, actual)
	}
}
