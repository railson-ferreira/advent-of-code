package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type coordinates struct {
	x int
	y int
}

type coordinatesPair struct {
	first  coordinates
	second coordinates
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
	galaxiesCoordinatesList := getGalaxiesCoordinates(matrix)
	pairs := getAllPairCombinations(galaxiesCoordinatesList)

	fmt.Printf("Original: (%v, %v)\n", len(matrix[0]), len(matrix))
	sum := 0
	for _, pair := range pairs {
		sum += pair.getMinDistance()
	}

	fmt.Println(sum)

}

func formatInput(input string) (matrix [][]rune) {
	lines := strings.Split(input, "\n")

	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		matrix = append(matrix, nil)
		for _, char := range line {
			matrix[y] = append(matrix[y], char)
		}
	}

	return matrix
}

func getGalaxiesCoordinates(matrix [][]rune) (coordinatesList []coordinates) {
	nonEmptyColumnsMap := make(map[int]bool)
	nonEmptyRowsMap := make(map[int]bool)

	for y, row := range matrix {
		for x, char := range row {
			if char != '.' {
				nonEmptyRowsMap[y] = true
				nonEmptyColumnsMap[x] = true
			}
		}
	}
	var nonEmptyColumns []int
	var nonEmptyRows []int

	for k, _ := range nonEmptyColumnsMap {
		nonEmptyColumns = append(nonEmptyColumns, k)
	}
	for k, _ := range nonEmptyRowsMap {
		nonEmptyRows = append(nonEmptyRows, k)
	}

	additionalY := 0
	for y, row := range matrix {
		if !slices.Contains(nonEmptyRows, y) {
			additionalY += 1000000 - 1
		}
		expandedY := y + additionalY
		additionalX := 0
		for x, char := range row {
			if !slices.Contains(nonEmptyColumns, x) {
				additionalX += 1000000 - 1
			}
			expandedX := x + additionalX
			if char == '#' {
				coordinatesList = append(coordinatesList, coordinates{
					x: expandedX,
					y: expandedY,
				})
			}
		}
	}
	return coordinatesList
}

func getAllPairCombinations(coordinatesList []coordinates) (pairs []coordinatesPair) {
	for i, first := range coordinatesList {
		for _, second := range coordinatesList[i+1:] {
			pairs = append(pairs, coordinatesPair{
				first:  first,
				second: second,
			})
		}
	}
	return pairs
}

func (pair coordinatesPair) getMinDistance() int {
	xDiff := pair.first.x - pair.second.x
	yDiff := pair.first.y - pair.second.y
	return int(math.Abs(float64(xDiff)) + math.Abs(float64(yDiff)))
}
