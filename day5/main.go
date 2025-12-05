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
	inventory, err := parseInput(flags.InputPath)
	if err != nil {
		panic(err)
	}
	inventory.CombineOverlap()

	var output uint64

	switch flags.Part {
	case 1:
		output, err = utils.FuncRunner(part1, inventory)
	case 2:
		output, err = utils.FuncRunner(part2, inventory)
	default:
		panic("invalid part or not implemented")
	}
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
