package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInputFile[T any](path string, processFunc func(string) (T, error)) ([]T, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read input file: %v", err)
	}
	defer file.Close()

	var output []T
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result, err := processFunc(line)
		if err != nil {
			return nil, fmt.Errorf("failed to process line %q: %v", line, err)
		}
		output = append(output, result)
	}
	return output, nil
}
