package main

import (
	"oac2025/utils"
)

func main() {
	if output, err := utils.Runner(part1, part2, parseInput); err != nil {
		panic(err)
	} else {
		println(output)
	}
}
