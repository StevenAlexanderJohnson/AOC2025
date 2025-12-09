package main

import "fmt"

func part2(homework Homework) (uint64, error) {
	var output uint64
	for _, question := range homework.Questions {
		var result uint64
		for _, cephalopodNumber := range question.CephalopodNumber() {
			switch question.Operation {
			case OperationAdd:
				result += cephalopodNumber
			case OperationMul:
				if result == 0 {
					result = 1
				}
				result *= cephalopodNumber
			default:
				return 0, fmt.Errorf("invalid operation for question: %s", question.Operation)
			}
		}
		output += result
	}

	return output, nil
}
