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
	expanded := expand(matrix)
	galaxiesCoordinatesList := getGalaxiesCoordinates(expanded)
	pairs := getAllPairCombinations(galaxiesCoordinatesList)

	fmt.Printf("Original: (%v, %v)\n", len(matrix[0]), len(matrix))
	fmt.Printf("Expanded: (%v, %v)\n", len(expanded[0]), len(expanded))
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

func expand(matrix [][]rune) (expanded [][]rune) {
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

	for y := 0; y < (len(matrix)-len(nonEmptyRows))+len(matrix); y++ {
		expanded = append(expanded, nil)
		for x := 0; x < (len(matrix[0])-len(nonEmptyColumns))+len(matrix[0]); x++ {
			expanded[y] = append(expanded[y], '.')
		}
	}

	additionalY := 0
	for y, row := range matrix {
		if !slices.Contains(nonEmptyRows, y) {
			additionalY++
		}
		expandedY := y + additionalY
		additionalX := 0
		for x, char := range row {
			if !slices.Contains(nonEmptyColumns, x) {
				additionalX++
			}
			expandedX := x + additionalX
			if char != '.' {
				expanded[expandedY][expandedX] = char
			}
		}
	}

	return expanded
}

func getGalaxiesCoordinates(matrix [][]rune) (coordinatesList []coordinates) {
	for y, row := range matrix {
		for x, char := range row {
			if char == '#' {
				coordinatesList = append(coordinatesList, coordinates{
					x: x,
					y: y,
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
