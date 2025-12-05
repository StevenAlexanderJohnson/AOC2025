package main

import "fmt"

func part1(inventory Inventory) (uint64, error) {
	var output uint64 = 0

	for i, id := range inventory.ids {
		fmt.Printf("%d\r", i)
		for _, idRange := range inventory.freshMap {
			if id >= idRange.start && id <= idRange.end {
				output++
				break
			}
		}
	}

	return output, nil
}
