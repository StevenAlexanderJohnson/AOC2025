package main

import (
	"fmt"
	"oac2025/utils"
)

func main() {
	if result, err := utils.Runner(part1, part2, parseInput); err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}
}
