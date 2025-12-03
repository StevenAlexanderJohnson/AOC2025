package main

import (
	"math"
)

func findLargeBattryValue(batteryBank []uint8) uint64 {
	outputParts := make([]uint8, 12)
	var output uint64

	startIndex := 0
	for i := range 12 {
		maxValue := uint8(0)
		maxIndex := uint8(0)
		for j := startIndex; j <= len(batteryBank)-12+i; j++ {
			if batteryBank[j] > maxValue {
				maxValue = batteryBank[j]
				maxIndex = uint8(j)
			}
		}
		outputParts[i] = maxValue
		startIndex = int(maxIndex) + 1
	}

	for i := 11; i >= 0; i-- {
		output += uint64(outputParts[i]) * uint64(math.Pow10(11-i))
	}
	return output
}

func part2(input [][]uint8) (uint64, error) {
	var output uint64
	for _, batteryBank := range input {
		output += findLargeBattryValue(batteryBank)
	}
	return output, nil
}
