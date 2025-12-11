package main

import (
	"fmt"
	"slices"
)

type heapItem struct {
	pointA, pointB point3D
	distance       float64
}

type minHeap struct {
	heap       []heapItem
	excludeMap map[int]map[int]bool
	tree       *kdTree
}

func newMinHeap(tree *kdTree, points []point3D) *minHeap {
	h := &minHeap{
		heap:       make([]heapItem, 0, len(points)),
		excludeMap: make(map[int]map[int]bool),
		tree:       tree,
	}
	for _, point := range points {
		h.excludeMap[point.id] = map[int]bool{point.id: true}
	}
	for _, point := range points {
		closestNeighbor := tree.nearestNeighbor(point, h.excludeMap[point.id])
		if closestNeighbor == nil {
			continue
		}
		dist := euclideanDistance(point, *closestNeighbor)
		h.push(heapItem{
			pointA:   point,
			pointB:   *closestNeighbor,
			distance: dist,
		})
		h.excludeMap[point.id][closestNeighbor.id] = true
		h.excludeMap[closestNeighbor.id][point.id] = true
	}
	h.minify()
	return h
}

func (h *minHeap) pop() heapItem {
	item := h.heap[0]
	h.heap = h.heap[1:]

	newNeighborA := h.tree.nearestNeighbor(item.pointA, h.excludeMap[item.pointA.id])
	if newNeighborA != nil {
		h.excludeMap[item.pointA.id][newNeighborA.id] = true
		h.excludeMap[newNeighborA.id][item.pointA.id] = true
		distA := euclideanDistance(item.pointA, *newNeighborA)
		h.push(heapItem{
			pointA:   item.pointA,
			pointB:   *newNeighborA,
			distance: distA,
		})
	}

	h.minify()
	fmt.Printf("Heap Size: %d\r", len(h.heap))
	return item
}

func (h *minHeap) len() int { return len(h.heap) }

func (h *minHeap) minify() {
	slices.SortFunc(h.heap, func(i, j heapItem) int {
		if i.distance < j.distance {
			return -1
		} else if i.distance > j.distance {
			return 1
		}
		return 0
	})
}

func (h *minHeap) push(item heapItem) bool {
	h.heap = append(h.heap, item)
	h.minify()
	return true
}
