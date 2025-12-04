package main

type board [][]rune

func parseInput(input string) ([]rune, error) {
	return []rune(input), nil
}

func (b board) testRoll(r int, c int) bool {
	rows := len(b)
	columns := len(b[0])
	if b[r][c] == '.' {
		return false
	}
	var surrounding int = 0
	// Check the surrounding 8 cells to see if there is a @ ignoring out of bounds and the current r, c
	if r > 0 && c > 0 && b[r-1][c-1] == '@' {
		surrounding++
	}
	if r > 0 && b[r-1][c] == '@' {
		surrounding++
	}
	if r > 0 && c < columns-1 && b[r-1][c+1] == '@' {
		surrounding++
	}
	if c > 0 && b[r][c-1] == '@' {
		surrounding++
	}
	if c < columns-1 && b[r][c+1] == '@' {
		surrounding++
	}
	if r < rows-1 && c > 0 && b[r+1][c-1] == '@' {
		surrounding++
	}
	if r < rows-1 && b[r+1][c] == '@' {
		surrounding++
	}
	if r < rows-1 && c < columns-1 && b[r+1][c+1] == '@' {
		surrounding++
	}
	if surrounding < 4 {
		return true
	}
	return false
}

func (b board) removeRoll(r int, c int) {
	b[r][c] = '.'
}
