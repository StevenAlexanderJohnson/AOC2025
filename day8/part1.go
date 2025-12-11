package main

import (
	"fmt"
	"slices"
)

func part1(pq *minHeap) (uint64, error) {
	// This tracks if the point ID has been seen and which circuit it belongs to
	seen := make(map[int]int)
	// This tracks the circuit IDs and their associated point IDs
	circuitIds := make(map[int][]int)
	newCircuitId := 0

	for range 10 {
		if pq.len() == 0 {
			return 0, fmt.Errorf("ran out of elements in pq")
		}
		item := pq.pop()
		pointA := item.pointA
		pointB := item.pointB

		idA, okA := seen[pointA.id]
		idB, okB := seen[pointB.id]

		// This happens when both points are already on the same circuit
		if !okA && !okB { // Neither are found in an existing circuit
			newCircuitId++
			seen[pointA.id] = newCircuitId
			seen[pointB.id] = newCircuitId
			circuitIds[newCircuitId] = []int{pointA.id, pointB.id}
		} else if okA && !okB { // One is found in an existing circuit
			seen[pointB.id] = idA
			circuitIds[idA] = append(circuitIds[idA], pointB.id)
		} else if !okA && okB { // The other is found in an existing circuit
			seen[pointA.id] = idB
			circuitIds[idB] = append(circuitIds[idB], pointA.id)
		} else if okA && okB && idA != idB { // Both are found in existing circuits, need to merge
			for _, pid := range circuitIds[idB] {
				seen[pid] = idA
				circuitIds[idA] = append(circuitIds[idA], pid)
			}
			delete(circuitIds, idB)
		}
	}
	largestIds := make([]int, 0, len(circuitIds))
	for _, points := range circuitIds {
		largestIds = append(largestIds, len(points))
	}
	slices.Sort(largestIds)
	n := len(largestIds)
	output := uint64(largestIds[n-1]) * uint64(largestIds[n-2]) * uint64(largestIds[n-3])
	fmt.Println(largestIds)
	return output, nil
}
