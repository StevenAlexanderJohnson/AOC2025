package main

import (
	"flag"
	"fmt"
	"oac2025/utils"
	"time"
)

func main() {
	var iterations int
	flag.IntVar(&iterations, "iteration", 1000, "How many iterations it should do before returning circuit information.")
	flags, err := utils.ParseFlags()
	if err != nil {
		panic(err)
	}

	setupStart := time.Now()
	input, err := parseInput(flags.InputPath)
	if err != nil {
		panic(err)
	}

	inputCopy := make([]point3D, len(input))
	copy(inputCopy, input)
	kdTree := newKdTree(inputCopy)
	pq := newMinHeap(kdTree, input)
	setupEnd := time.Now()
	fmt.Printf("Setup took %v\n", setupEnd.Sub(setupStart))

	result := uint64(0)
	switch flags.Part {
	case 1:
		result, err = utils.FuncRunner(func(heap *minHeap) (uint64, error) {
			return part1(heap, iterations)
		}, pq)
	case 2:
		panic("part 2 not implemented yet")
	default:
		panic("invalid part specified")
	}

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
