package main

func part2(input [][]rune) (uint64, error) {
	board := board(input)
	var output uint64 = 0

	removeIndicies := make([][2]int, 0)
	for {
		removeIndicies = removeIndicies[:0]

		for r := range input {
			for c := range input[r] {
				if board.testRoll(r, c) {
					output++
					removeIndicies = append(removeIndicies, [2]int{r, c})
				}
			}
		}

		if len(removeIndicies) == 0 {
			break
		}

		for _, indicies := range removeIndicies {
			board.removeRoll(indicies[0], indicies[1])
		}
	}
	return output, nil
}
