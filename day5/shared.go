package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type idRange struct {
	start uint64
	end   uint64
}

type Inventory struct {
	ids      []uint64
	freshMap []idRange
}

func (inv *Inventory) CombineOverlap() {
	if len(inv.freshMap) <= 1 {
		return
	}

	sort.Slice(inv.freshMap, func(a, b int) bool {
		return inv.freshMap[a].start < inv.freshMap[b].start
	})

	for i := 1; i < len(inv.freshMap); i++ {
		previous := &inv.freshMap[i-1]
		current := &inv.freshMap[i]
		if current.start <= previous.end {
			if current.end >= previous.end {
				previous.end = current.end
				// remove current from slice
				inv.freshMap = append(inv.freshMap[:i], inv.freshMap[i+1:]...)
				i-- // stay at the same index for next iteration
			} else if current.end < previous.end {
				// current is fully contained within previous, just remove it
				inv.freshMap = append(inv.freshMap[:i], inv.freshMap[i+1:]...)
				i-- // stay at the same index for next iteration
			}
		}
	}
}

func processIdRange(rangeString string, freshMap *[]idRange) error {
	var start uint64
	var end uint64
	if _, err := fmt.Sscanf(rangeString, "%d-%d", &start, &end); err != nil {
		return fmt.Errorf("error occurred while parsing id range %s: %v", rangeString, err)
	}

	*freshMap = append(*freshMap, idRange{start: start, end: end})

	return nil
}

func processId(idString string, ids *[]uint64) error {
	var id uint64
	if _, err := fmt.Sscanf(idString, "%d", &id); err != nil {
		return fmt.Errorf("error occurred while parsing id %s: %v", idString, err)
	}

	*ids = append(*ids, id)
	return nil
}

func parseInput(inputPath string) (Inventory, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return Inventory{}, fmt.Errorf("error occurred while opening the input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ids []uint64
	freshMap := make([]idRange, 0)

	i := 0
	processingFresh := true
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			processingFresh = false
			continue
		}
		fmt.Printf("Processing line: %d\r", i)
		switch processingFresh {
		case true:
			if err := processIdRange(line, &freshMap); err != nil {
				return Inventory{}, err
			}
		case false:
			if err := processId(line, &ids); err != nil {
				return Inventory{}, err
			}
		}
	}
	return Inventory{ids: ids, freshMap: freshMap}, nil
}
