package main

import "fmt"

type directionType string

const (
	DirectionLeft  directionType = "L"
	DirectionRight directionType = "R"
)

type rotationDirection struct {
	direction directionType
	count     int
}

func (r rotationDirection) String() string {
	return fmt.Sprintf("%v%d", r.direction, r.count)
}

func parseInput(input string) (rotationDirection, error) {
	var r rotationDirection
	var d rune
	if _, err := fmt.Sscanf(input, "%c%d", &d, &r.count); err != nil {
		return r, fmt.Errorf("the input line is invalid format: %s", input)
	}

	switch d {
	case 'R':
		r.direction = DirectionRight
	case 'L':
		r.direction = DirectionLeft
	default:
		return rotationDirection{}, fmt.Errorf("invalid direction string provided in input: %c", d)
	}
	if r.String() != input {
		return r, fmt.Errorf("parse failed and resulted in invalid string: %s -> %s", input, r.String())
	}
	return r, nil
}
