package main

import "fmt"

func part1(homework Homework) (uint64, error) {
	var total uint64 = 0
	for _, question := range homework.Questions {
		var result uint64
		for _, number := range question.Numbers {
			switch question.Operation {
			case OperationAdd:
				result += number
			case OperationMul:
				if result == 0 {
					result = 1
				}
				result *= number
			default:
				return 0, fmt.Errorf("unknown operation for question: %s", question.Operation)
			}
		}
		total += result
	}
	return total, nil
}
