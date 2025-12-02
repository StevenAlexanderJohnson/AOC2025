package main

import (
	"fmt"
)

// This enforces that the string is split into k equal parts.
func splitIntoParts(str string, k int) ([]string, bool) {
	result := make([]string, k)

	if len(str)%k != 0 {
		return nil, false
	}

	partLength := len(str) / k

	for i := range k {
		start := i * partLength
		end := start + partLength
		result[i] = str[start:end]
	}

	return result, true
}

// Ensure all strings within the provided parts are equal.
func partsMatch(parts []string) bool {
	if len(parts) < 2 {
		return false
	}

	for _, part := range parts {
		if part != parts[0] {
			return false
		}
	}

	return true
}

// This manages the helper functions. It tries to find a window that does have equal windows.
func findWindow(ids uint64) uint64 {
	idString := fmt.Sprintf("%d", ids)
	for i := 1; i <= len(idString); i++ {
		parts, success := splitIntoParts(idString, i)
		if !success {
			continue
		}
		if partsMatch(parts) {
			var repeat uint64
			if _, err := fmt.Sscanf(parts[0], "%d", &repeat); err != nil {
				panic("error parsing repeat")
			}
			return ids
		}
	}
	return 0
}

func part2(input []IDRange) uint64 {
	var output uint64
	for _, ids := range input {
		for id := ids.Start; id <= ids.End; id++ {
			output += findWindow(id)
		}
	}
	return output
}
