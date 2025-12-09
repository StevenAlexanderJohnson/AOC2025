package main

import "fmt"

func part2(diagram *diagram) (uint64, error) {
	memo := make(map[position]uint64)

	var dfs func(pos position) uint64
	dfs = func(pos position) uint64 {

		if val, exists := memo[pos]; exists {
			return val
		}

		if pos.y >= len(diagram.board) {
			return 1
		}
		if pos.x < 0 || pos.x >= len(diagram.board[0]) {
			return 0
		}

		nextPositions := []position{}

		currentChar := diagram.board[pos.y][pos.x]
		switch currentChar {
		case 'S', '.':
			nextPositions = append(nextPositions, position{x: pos.x, y: pos.y + 1})
		case '^':
			nextPositions = append(nextPositions, position{x: pos.x + 1, y: pos.y}, position{x: pos.x - 1, y: pos.y})
		}

		var result uint64 = 0
		for _, nextPos := range nextPositions {
			result += dfs(nextPos)
		}
		memo[pos] = result

		return result
	}

	if len(diagram.stack) == 0 {
		return 0, fmt.Errorf("no starting position 'S' found in the diagram")
	}

	output := dfs(diagram.stack[0])
	return output, nil
}
