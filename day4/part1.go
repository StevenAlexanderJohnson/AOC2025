package main

func part1(input [][]rune) (uint64, error) {
	board := board(input)
	var output uint64 = 0

	for r := range input {
		for c := range input[r] {
			if board.testRoll(r, c) {
				output++
			}
		}
	}

	return output, nil
}
