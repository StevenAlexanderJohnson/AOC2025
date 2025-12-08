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

	input, err := parseInput(flags.InputPath, flags.Part)
	if err != nil {
		panic(err)
	}

	var output uint64
	switch flags.Part {
	case 1:
		output, err = utils.FuncRunner(part1, input)
	case 2:
		output, err = utils.FuncRunner(part2, input)
	default:
		panic("invalid part or not implemented")
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part %d: %d\n", flags.Part, output)
}
