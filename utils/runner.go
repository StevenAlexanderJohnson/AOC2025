package utils

func Runner[T, K any](part1 func([]T) (K, error), part2 func([]T) (K, error), parseInputLine func(string) (T, error)) (K, error) {
	flags, err := ParseFlags()
	if err != nil {
		var zero K
		return zero, err
	}

	input, err := ReadInputFile(flags.InputPath, parseInputLine)
	if err != nil {
		var zero K
		return zero, err
	}

	var output K
	switch flags.Part {
	case 1:
		output, err = part1(input)
	case 2:
		output, err = part2(input)
	default:
		panic("invalid part or not implemented")
	}

	if err != nil {
		var zero K
		return zero, err
	}

	return output, nil
}
