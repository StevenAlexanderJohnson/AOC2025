package main

import (
	"math"
)

func getTotalRotations(rotationCount int) int64 {
	rotationCount = rotationCount / 100
	return int64(math.Abs(float64(rotationCount)))
}

func part2(input []rotationDirection) int64 {
	position := 50
	var output int64 = 0
	for _, rotation := range input {
		var zero_passes int64 = getTotalRotations(rotation.count)
		switch rotation.direction {
		case DirectionRight:
			position += rotation.count % 100
			if position >= 100 {
				position -= 100
				zero_passes += 1
			}
		case DirectionLeft:
			// Necessary in case we start at 0, starting at zero does not count as passing it.
			// Checking before we change it
			shouldCount := position != 0
			position -= rotation.count % 100
			// If it's now at or below zero and we weren't already at zero that's when we want to count
			if shouldCount && (position == 0 || position < 0) {
				zero_passes += 1
			}
			if position < 0 {
				position += 100
			}
		}
		output += zero_passes
	}
	return output
}
