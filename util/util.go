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
