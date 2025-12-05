package main

func part2(inventory Inventory) (uint64, error) {
	var output uint64 = 0
	for _, id := range inventory.freshMap {
		output += id.end - id.start + 1
	}
	return output, nil
}
