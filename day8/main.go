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

	input, err := parseInput(flags.InputPath)
	if err != nil {
		panic(err)
	}

	inputCopy := make([]point3D, len(input))
	copy(inputCopy, input)
	kdTree := newKdTree(inputCopy)
	pq := newMinHeap(kdTree, input)

	result, err := part1(pq)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
