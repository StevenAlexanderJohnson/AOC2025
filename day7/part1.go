package main

func part1(diagram *diagram) (uint64, error) {
	for len(diagram.stack) > 0 {
		pos := diagram.stack[len(diagram.stack)-1]
		diagram.stack = diagram.stack[:len(diagram.stack)-1]

		if pos.y >= len(diagram.board) {
			continue
		}
		if pos.x < 0 || pos.x >= len(diagram.board[pos.y]) {
			continue
		}

		char := diagram.board[pos.y][pos.x]
		// Don't include '|' as a valid path because that means a path has already takken it.
		switch char {
		case 'S':
			diagram.stack = append(diagram.stack, position{x: pos.x, y: pos.y + 1})
		case '^':
			diagram.splits++
			diagram.stack = append(diagram.stack, position{x: pos.x - 1, y: pos.y})
			diagram.stack = append(diagram.stack, position{x: pos.x + 1, y: pos.y})
		case '.':
			diagram.board[pos.y][pos.x] = '|'
			diagram.stack = append(diagram.stack, position{x: pos.x, y: pos.y + 1})
		}
	}
	return diagram.splits, nil
}
