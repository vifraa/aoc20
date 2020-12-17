package util

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInputAsInt(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func ReadInputAsString(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func EqualsStringSlice(i1, i2 []string) bool {
	if len(i1) != len(i2) {
		return false
	}

	for i := 0; i < len(i1); i++ {
		s1 := i1[i]
		s2 := i2[i]

		if s1 != s2 {
			return false
		}
	}
	return true
}
