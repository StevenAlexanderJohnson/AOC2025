package main

import (
	"bufio"
	"fmt"
	"os"
)

func euclideanDistance(p1, p2 point3D) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	dz := p1.z - p2.z
	return (dx * dx) + (dy * dy) + (dz * dz)
}

type point3D struct {
	x, y, z float64
	id      int
}

func parseInput(inputPath string) ([]point3D, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var points []point3D
	id := 0
	for scanner.Scan() {
		line := scanner.Text()
		var point point3D
		if _, err := fmt.Sscanf(line, "%f,%f,%f", &point.x, &point.y, &point.z); err != nil {
			return nil, err
		}
		point.id = id
		id++
		points = append(points, point)
	}
	return points, nil
}
