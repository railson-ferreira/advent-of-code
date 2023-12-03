package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type schematicItem struct {
	char                   string
	adjacentSymbol         string
	adjacentSymbolPosition coordinates
}

type coordinates struct {
	x int
	y int
}

type partNumberDetails struct {
	partNumber     int
	symbol         string
	symbolPosition coordinates
}

type gear struct {
	adjacentPartNumbers []int
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

	partNumbersDetails := getPartNumbersDetails(matrix)

	gears := make(map[coordinates]gear)
	for _, thePartNumberDetails := range partNumbersDetails {
		if thePartNumberDetails.symbol == "*" {
			var adjacentPartNumbers = gears[thePartNumberDetails.symbolPosition].adjacentPartNumbers
			adjacentPartNumbers = append(adjacentPartNumbers, thePartNumberDetails.partNumber)
			// just putting stuff back to the map
			var theGear = gears[thePartNumberDetails.symbolPosition]
			theGear.adjacentPartNumbers = adjacentPartNumbers
			gears[thePartNumberDetails.symbolPosition] = theGear
		}
	}
	var gearRatios []int
	for _, value := range gears {
		if len(value.adjacentPartNumbers) == 2 {
			firstPartNumber := value.adjacentPartNumbers[0]
			secondPartNumber := value.adjacentPartNumbers[1]
			gearRatios = append(gearRatios, firstPartNumber*secondPartNumber)
		}
	}
	var sum int
	for _, value := range gearRatios {
		sum += value
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
				char: string(char),
			})
		}

	}
	for y := range matrix {
		row := matrix[y]
		for x := range row {
			matrix[y][x].adjacentSymbol, matrix[y][x].adjacentSymbolPosition = getAdjacentSymbol(matrix, x, y)
		}
	}

	return matrix
}

func getAdjacentSymbol(matrix [][]schematicItem, x int, y int) (symbol string, position coordinates) {
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
					return item.char, *theCoordinates
				}
			}
		}
	}
	return symbol, position
}

func getPartNumbersDetails(matrix [][]schematicItem) (partNumbersDetails []partNumberDetails) {
	for _, row := range matrix {
		partialNumber := ""
		var adjacentChar string
		var adjacentCharPosition coordinates
		for index, item := range row {
			if value, err0 := strconv.Atoi(item.char); err0 == nil {
				partialNumber += strconv.Itoa(value)
				if adjacentChar == "" {
					adjacentChar = item.adjacentSymbol
					adjacentCharPosition = item.adjacentSymbolPosition
				}
				if index < len(row)-1 {
					continue
				}
			}
			if len(partialNumber) > 0 && adjacentChar != "" {
				partNumber, err1 := strconv.Atoi(partialNumber)
				thePartNumberDetails := partNumberDetails{
					partNumber:     partNumber,
					symbol:         adjacentChar,
					symbolPosition: adjacentCharPosition,
				}
				check(err1)
				partNumbersDetails = append(partNumbersDetails, thePartNumberDetails)
			}
			partialNumber = ""
			adjacentChar = ""
			adjacentCharPosition = coordinates{}
		}
	}
	return partNumbersDetails
}
