package main

import (
	"fmt"
	"oac2025/utils"
)

func main() {
	flags, err := utils.ParseFlags()
	if err != nil {
		panic(err)
	}

	input, err := utils.ReadInputFile(flags.InputPath, parseInput)
	if err != nil {
		panic(err)
	}

	var result uint64
	switch flags.Part {
	case 1:
		result = part1(input)
	case 2:
		result = part2(input)
	default:
		panic("invalid part provided or not implemented")
	}

	fmt.Println(result)
}
