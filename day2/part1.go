package main

import (
	"fmt"
	"strings"
)

func calculateNumberOfInvalid(ids IDRange) uint64 {
	var output uint64 = 0
	for id := ids.Start; id <= ids.End; id++ {
		idString := fmt.Sprintf("%d", id)
		halfLength := len(idString) / 2
		firstHalf := idString[:halfLength]
		secondHalf := idString[halfLength:]
		if strings.Compare(firstHalf, secondHalf) == 0 {
			output += id
		}
	}
	return output
}

func part1(input []IDRange) (uint64, error) {
	var output uint64 = 0
	for _, r := range input {
		output += calculateNumberOfInvalid(r)
	}
	return output, nil
}
