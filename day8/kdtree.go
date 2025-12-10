package main

import (
	"math"
	"slices"
)

type kdTreeNode struct {
	*point3D
	left  *kdTreeNode
	right *kdTreeNode
}

func (n *kdTreeNode) insert(point point3D, depth int) {
	if n.point3D == nil {
		n.point3D = &point
		return
	}
	switch depth % 3 {
	case 0:
		if point.x < n.x {
			if n.left == nil {
				n.left = &kdTreeNode{point3D: &point}
			} else {
				n.left.insert(point, depth+1)
			}
		} else {
			if n.right == nil {
				n.right = &kdTreeNode{point3D: &point}
			} else {
				n.right.insert(point, depth+1)
			}
		}
	case 1:
		if point.y < n.y {
			if n.left == nil {
				n.left = &kdTreeNode{point3D: &point}
			} else {
				n.left.insert(point, depth+1)
			}
		} else {
			if n.right == nil {
				n.right = &kdTreeNode{point3D: &point}
			} else {
				n.right.insert(point, depth+1)
			}
		}
	case 2:
		if point.z < n.z {
			if n.left == nil {
				n.left = &kdTreeNode{point3D: &point}
			} else {
				n.left.insert(point, depth+1)
			}
		} else {
			if n.right == nil {
				n.right = &kdTreeNode{point3D: &point}
			} else {
				n.right.insert(point, depth+1)
			}
		}
	}
}

func (n *kdTreeNode) nearestSearch(target point3D, depth int, bestPoint **point3D, bestDistSqrt *float64, exclude map[int]bool) {
	if n == nil {
		return
	}

	isSelf := (n.point3D.id == target.id)
	_, isExcluded := exclude[n.point3D.id]

	if !isSelf && !isExcluded {
		currentDistSq := euclideanDistance(*n.point3D, target)
		if currentDistSq < *bestDistSqrt {
			*bestDistSqrt = currentDistSq
			*bestPoint = n.point3D
		}
	}

	axis := depth % 3
	var primaryBranch, secondaryBranch *kdTreeNode
	var queryVal, nodeVal float64
	switch axis {
	case 0:
		queryVal = target.x
		nodeVal = n.x
	case 1:
		queryVal = target.y
		nodeVal = n.y
	case 2:
		queryVal = target.z
		nodeVal = n.z
	}

	if queryVal < nodeVal {
		primaryBranch = n.left
		secondaryBranch = n.right
	} else {
		primaryBranch = n.right
		secondaryBranch = n.left
	}

	primaryBranch.nearestSearch(target, depth+1, bestPoint, bestDistSqrt, exclude)

	planeDist := queryVal - nodeVal
	planeDistSq := planeDist * planeDist

	if planeDistSq < *bestDistSqrt {
		secondaryBranch.nearestSearch(target, depth+1, bestPoint, bestDistSqrt, exclude)
	}
}

type kdTree struct {
	root *kdTreeNode
}

func (t *kdTree) nearestNeighbor(target point3D, exclude map[int]bool) *point3D {
	if t.root == nil || t.root.point3D == nil {
		return nil
	}
	bestPoint := t.root.point3D
	bestDistSqrt := math.MaxFloat64

	t.root.nearestSearch(target, 0, &bestPoint, &bestDistSqrt, exclude)
	return bestPoint
}

func newKdTree(points []point3D) *kdTree {
	tree := &kdTree{root: &kdTreeNode{}}
	getSortFunction := func(plane int) func(i, j point3D) int {
		switch plane {
		case 0:
			return func(i, j point3D) int {
				if i.x < j.x {
					return -1
				} else if i.x > j.x {
					return 1
				}
				return 0
			}
		case 1:
			return func(i, j point3D) int {
				if i.y < j.y {
					return -1
				} else if i.y > j.y {
					return 1
				}
				return 0
			}
		case 2:
			return func(i, j point3D) int {
				if i.z < j.z {
					return -1
				} else if i.z > j.z {
					return 1
				}
				return 0
			}
		default:
			return nil
		}
	}
	for i := range points {
		plain := i % 3
		slices.SortFunc(points, getSortFunction(plain))
		medianPoint := points[len(points)/2]
		points = append(points[:len(points)/2], points[len(points)/2+1:]...)
		tree.root.insert(medianPoint, 0)
	}
	return tree
}
