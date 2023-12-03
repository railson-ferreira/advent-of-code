package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type schematicItem struct {
	char                string
	isAdjacentToASymbol bool
}

type coordinates struct {
	x int
	y int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	bytes, err := os.ReadFile("input.txt")
	check(err)
	input := string(bytes)

	matrix := formatInput(input)

	partNumbers := getPartNumbers(matrix)

	var sum int
	for _, partNumber := range partNumbers {
		sum += partNumber
	}
	fmt.Println(sum)
}

func formatInput(input string) [][]schematicItem {
	lines := strings.Split(input, "\n")
	var matrix [][]schematicItem

	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		matrix = append(matrix, []schematicItem{})

		for x := range line {
			char := line[x]
			matrix[y] = append(matrix[y], schematicItem{
				char:                string(char),
				isAdjacentToASymbol: false,
			})
		}

	}
	for y := range matrix {
		row := matrix[y]
		for x := range row {
			matrix[y][x].isAdjacentToASymbol = isThereAdjacentSymbol(matrix, x, y)
		}
	}

	return matrix
}

func isThereAdjacentSymbol(matrix [][]schematicItem, x int, y int) bool {
	maxY := len(matrix) - 1
	maxX := len(matrix[0]) - 1

	dotChar := "."

	var lt, t, rt, r, rb, b, lb, l *coordinates
	if x > 0 {
		l = &coordinates{x: x - 1, y: y}
		if y > 0 {
			lt = &coordinates{x: x - 1, y: y - 1}
		}
		if y < maxY {
			lb = &coordinates{x: x - 1, y: y + 1}
		}
	}
	if x < maxX {
		r = &coordinates{x: x + 1, y: y}
		if y > 0 {
			rt = &coordinates{x: x + 1, y: y - 1}
		}
		if y < maxY {
			rb = &coordinates{x: x + 1, y: y + 1}
		}
	}
	if y > 0 {
		b = &coordinates{x: x, y: y - 1}
	}
	if y < maxY {
		t = &coordinates{x: x, y: y + 1}
	}
	for _, theCoordinates := range []*coordinates{lt, t, rt, r, rb, b, lb, l} {
		if theCoordinates != nil {
			item := matrix[theCoordinates.y][theCoordinates.x]

			if item.char != dotChar {
				if _, err := strconv.Atoi(item.char); err != nil {
					return true
				}
			}
		}
	}
	return false
}

func getPartNumbers(matrix [][]schematicItem) (partNumbers []int) {
	for _, row := range matrix {
		partialNumber := ""
		isAdjacent := false
		for index, item := range row {
			if value, err0 := strconv.Atoi(item.char); err0 == nil {
				partialNumber += strconv.Itoa(value)
				isAdjacent = isAdjacent || item.isAdjacentToASymbol
				if index < len(row)-1 {
					continue
				}
			}
			if len(partialNumber) > 0 && isAdjacent {
				partNumber, err1 := strconv.Atoi(partialNumber)
				check(err1)
				partNumbers = append(partNumbers, partNumber)
			}
			partialNumber = ""
			isAdjacent = false
		}
	}
	return partNumbers
}
