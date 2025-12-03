package main

func parseInput(input string) ([]uint8, error) {
	output := make([]uint8, len(input))

	for i, num := range input {
		output[i] = uint8(num - '0')
	}

	return output, nil
}
