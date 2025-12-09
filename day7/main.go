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

	diagram, err := parseInput(flags.InputPath)
	if err != nil {
		panic(err)
	}

	var output uint64
	switch flags.Part {
	case 1:
		output, err = utils.FuncRunner(part1, diagram)
	case 2:
		output, err = utils.FuncRunner(part2, diagram)
	default:
		panic("invalid part")
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
