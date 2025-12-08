package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type operation string

const (
	OperationAdd    operation = "+"
	OperationMul    operation = "*"
	OperationNotSet operation = ""
)

type column struct {
	startIndex int
	endIndex   int
}

type HomeworkQuestion struct {
	Numbers   []uint64
	Operation operation
}

func (hq HomeworkQuestion) LargestNumberLength() int {
	var maxLength int
	for _, number := range hq.Numbers {
		length := int(math.Log10(float64(number))) + 1
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}

func (hq HomeworkQuestion) CephalopodNumber() []uint64 {
	output := make([]uint64, hq.LargestNumberLength())
	for i := hq.LargestNumberLength() - 1; i >= 0; i-- {
		var column uint64
		for j, number := range hq.Numbers {
			x := ((number / uint64(math.Pow10(i))) % 10) * uint64(math.Pow10(hq.LargestNumberLength()-1-j))
			if x == 0 {
				column /= 10
				continue
			}
			column += x
		}
		output[i] = column
	}
	return output
}

func newHomeworkQuestion() HomeworkQuestion {
	return HomeworkQuestion{
		Numbers:   make([]uint64, 0),
		Operation: OperationNotSet,
	}
}

type Homework struct {
	Questions []HomeworkQuestion
}

func newHomework(questionCount int) Homework {
	questions := make([]HomeworkQuestion, questionCount)
	for i := range questions {
		questions[i] = newHomeworkQuestion()
	}
	return Homework{
		Questions: questions,
	}
}

func processLine(line string, questions []HomeworkQuestion) ([]HomeworkQuestion, error) {
	parts := strings.Fields(strings.TrimSpace(line))
	for i, part := range parts {
		if len(part) == 1 && (part[0] == '+' || part[0] == '*') {
			switch part[0] {
			case '+':
				questions[i].Operation = OperationAdd
			case '*':
				questions[i].Operation = OperationMul
			}
		} else {
			var number uint64
			if _, err := fmt.Sscanf(part, "%d", &number); err != nil {
				return questions, fmt.Errorf("failed to scan question part into number: %w: %s", err, part)
			}
			questions[i].Numbers = append(questions[i].Numbers, number)
		}
	}
	return questions, nil
}

func getColumnLengths(operationLine string) []column {
	start := 0
	output := make([]column, 0)
	for i, char := range operationLine {
		if char == '+' || char == '*' {
			if i == 0 {
				continue
			}
			output = append(output, column{
				startIndex: start,
				endIndex:   i - 1,
			})
			start = i
		}
	}
	output = append(output, column{
		startIndex: start,
		endIndex:   len(operationLine),
	})
	return output
}

func processLine2(line string, columns []column, questions []HomeworkQuestion) ([]HomeworkQuestion, error) {
	for i, col := range columns {
		if line[0] == '+' || line[0] == '*' {
			switch line[col.startIndex] {
			case '+':
				questions[i].Operation = OperationAdd
			case '*':
				questions[i].Operation = OperationMul
			}
		} else {
			window := line[col.startIndex:col.endIndex]
			window = strings.TrimLeft(window, " ")
			window = strings.ReplaceAll(window, " ", "0")
			var number uint64
			if _, err := fmt.Sscanf(window, "%d", &number); err != nil {
				return []HomeworkQuestion{}, fmt.Errorf("failed to scan question part into number: %w: %s", err, window)
			}
			questions[i].Numbers = append(questions[i].Numbers, number)
		}
	}
	return questions, nil
}
func parseInput(inputPath string, part int) (Homework, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return Homework{}, fmt.Errorf("failed to open input file: %w", err)
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return Homework{}, fmt.Errorf("failed to read input file: %w", err)
	}

	homework := newHomework(len(strings.Fields(lines[0])))
	for _, line := range lines {
		if part == 1 {
			homework.Questions, err = processLine(line, homework.Questions)
		} else {
			columns := getColumnLengths(lines[len(lines)-1])
			homework.Questions, err = processLine2(line, columns, homework.Questions)
		}
		if err != nil {
			return Homework{}, fmt.Errorf("failed to process line: %w", err)
		}
	}

	return homework, nil
}
