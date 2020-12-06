package main

import "testing"

func TestBirthYear_IsValid(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"2002", true},
		{"2003", false},
		{"1920", true},
		{"1919", false},
		{"191a", false},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			s := birthYear.IsValid(tt.in)
			if s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}
		})
	}
}

func TestIssueYear_IsValid(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"2009", false},
		{"2010", true},
		{"2020", true},
		{"2021", false},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			s := issueYear.IsValid(tt.in)
			if s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}
		})
	}
}

func TestExpirationYear_IsValid(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"2019", false},
		{"2020", true},
		{"2030", true},
		{"2031", false},
		{"ffff", false},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			s := expirationYear.IsValid(tt.in)
			if s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}
		})
	}
}

func TestHeight_IsValid(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"149cm", false},
		{"150cm", true},
		{"156cm", true},
		{"193cm", true},
		{"194cm", false},
		{"58in", false},
		{"59in", true},
		{"76in", true},
		{"77in", false},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			s := height.IsValid(tt.in)
			if s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}
		})
	}
}

func TestHairColor_IsValid(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"#123456", true},
		{"#abcdef", true},
		{"#abc123", true},
		{"0123456", false},
		{"#1234567", false},
		{"#abc123", true},
		{"#abcdefghijk", false},
		{"#12345", false},
		{"#1234b", false},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			s := hairColor.IsValid(tt.in)
			if s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}
		})
	}
}

func TestEyeColor_IsValid(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"amb", true},
		{"blu", true},
		{"brn", true},
		{"gry", true},
		{"grn", true},
		{"hzl", true},
		{"oth", true},
		{"#12345", false},
		{"#1234b", false},
		{"abc", false},
		{"123", false},
		{"ambamb", false},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			s := eyeColor.IsValid(tt.in)
			if s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}
		})
	}
}

func TestPassportId_IsValid(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"000123456", true},
		{"123456789", true},
		{"000000001", true},
		{"000000000", false},
		{"12345678", false},
		{"12345678a", false},
		{"abcdefghijk", false},
		{"0123456789", false},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			s := passportID.IsValid(tt.in)
			if s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}
		})
	}
}
