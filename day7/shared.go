package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	x int
	y int
}

type diagram struct {
	board  [][]rune
	splits uint64
	stack  []position
}

func parseInput(inputPath string) (*diagram, error) {
	var diagram diagram

	file, err := os.Open(inputPath)
	if err != nil {
		return &diagram, fmt.Errorf("failed to open input file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lien := scanner.Text()
		row := []rune(lien)
		diagram.board = append(diagram.board, row)
	}
	if err := scanner.Err(); err != nil {
		return &diagram, fmt.Errorf("failed to read input file: %w", err)
	}

	for y, row := range diagram.board {
		for x, char := range row {
			if char == 'S' {
				diagram.stack = append(diagram.stack, position{x: x, y: y})
			}
		}
	}

	return &diagram, nil
}
