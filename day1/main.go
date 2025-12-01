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

	var output int64
	switch flags.Part {
	case 1:
		output = part1(input)
	case 2:
		output = part2(input)
	default:
		panic("invalid part or not implemented")
	}

	fmt.Println(output)
}
