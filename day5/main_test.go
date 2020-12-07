package main

import "testing"

func TestFindRow(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"FBFBBFF", 44},
		{"FFFBBBF", 14},
		{"BBFFBBF", 102},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			s := FindRow(tt.in)
			if s != tt.out {
				t.Errorf("got %d, want %d", s, tt.out)
			}
		})
	}
}

func TestFindCol(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"RRR", 7},
		{"RLL", 4},
		{"LLL", 0},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			s := FindCol(tt.in)
			if s != tt.out {
				t.Errorf("got %d, want %d", s, tt.out)
			}
		})
	}
}
