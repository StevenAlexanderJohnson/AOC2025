package main

func part1(input []rotationDirection) int64 {
	position := 50
	var output int64 = 0
	for _, rotation := range input {
		switch rotation.direction {
		case DirectionRight:
			position += rotation.count % 100
			if position >= 100 {
				position -= 100
			}
		case DirectionLeft:
			position -= rotation.count % 100
			if position < 0 {
				position += 100
			}
		}

		if position == 0 {
			output++
		}
	}
	return output
}
