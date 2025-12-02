package main

import (
	"fmt"
)

type IDRange struct {
	Start uint64
	End   uint64
}

func parseInput(input string) (IDRange, error) {
	var output IDRange
	if _, err := fmt.Sscanf(input, "%d-%d", &output.Start, &output.End); err != nil {
		return IDRange{}, fmt.Errorf("an invalid id range was provided: %v", err)
	}

	return output, nil
}
