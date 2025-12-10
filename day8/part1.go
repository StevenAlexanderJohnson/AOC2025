package main

import "fmt"

func part1(pq *minHeap) (int, error) {
	seen := make(map[int]int)
	circuitIds := make(map[int][]int)
	newCircuitId := 0

	for i := 0; i < 10; i++ {
		if pq.len() == 0 {
			return 0, fmt.Errorf("ran out of elements in pq")
		}
		item := pq.pop()
		pointA := item.pointA
		pointB := item.pointB

		idA, okA := seen[pointA.id]
		idB, okB := seen[pointB.id]

		if okA && okB && idA == idB {
			i--
			continue
		}

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
		} else { // Both are found in existing circuits, need to merge
			if idA == idB {
				continue
			}
			for _, pid := range circuitIds[idB] {
				seen[pid] = idA
			}
			delete(circuitIds, idB)
		}
	}
	return newCircuitId, nil
}
