package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

type block struct {
	char rune
	x    int
	y    int
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

	fmt.Println("Total load before tilt:", getLoad(matrix))
	start := time.Now()
	var endOfFirstCycleMatrices [][][]rune
	var totalCycles = 1000_000_000
	alreadySkipped := false
	for count := 0; count < totalCycles; count++ {
		tilt('t', matrix)
		tilt('l', matrix)
		tilt('b', matrix)
		tilt('r', matrix)
		fmt.Println(count, alreadySkipped)
		if alreadySkipped {
			continue
		}
		if index := indexOfMatrix(endOfFirstCycleMatrices, matrix); index != -1 {
			distance := count - index
			remaining := (totalCycles - 1) - count
			count += distance * (remaining / distance)
			alreadySkipped = true
		}
		endOfFirstCycleMatrices = append(endOfFirstCycleMatrices, copyMatrix(matrix))
	}
	fmt.Println("Elapsed: ", time.Now().Sub(start))
	fmt.Println("Total load after cycles:", getLoad(matrix))
	//_ = os.WriteFile("view.txt", []byte(getOutput(matrix)), 0644)
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

func tilt(dir rune, matrix [][]rune) {
	if !slices.Contains([]rune{'t', 'r', 'b', 'l'}, dir) {
		panic("Invalid direction")
	}
	maxX := len(matrix[0]) - 1
	maxY := len(matrix) - 1

	for yy := 0; yy <= maxY; yy++ {
		for xx := 0; xx <= maxX; xx++ {
			x := xx
			y := yy
			if dir == 'r' || dir == 'b' {
				x = maxX - xx
				y = maxY - yy
			}
			char := matrix[y][x]
			if char == 'O' {
				near := getNear(dir, x, y, matrix)
				var replacementX int
				var replacementY int
				if near != nil {
					replacementX = near.x
					replacementY = near.y
					switch dir {
					case 't':
						replacementY++
					case 'r':
						replacementX--
					case 'b':
						replacementY--
					case 'l':
						replacementX++
					}
				} else {
					replacementX = x
					replacementY = y
					switch dir {
					case 't':
						replacementY = 0
					case 'r':
						replacementX = maxX
					case 'b':
						replacementY = maxY
					case 'l':
						replacementX = 0
					}
				}
				// do not need to check if it's at the same position
				matrix[y][x] = '.'
				matrix[replacementY][replacementX] = 'O'
			}
		}
	}
}

func getNear(dir rune, x int, y int, matrix [][]rune) *block {
	if !slices.Contains([]rune{'t', 'r', 'b', 'l'}, dir) {
		panic("Invalid direction")
	}
	maxX := len(matrix[0]) - 1
	maxY := len(matrix) - 1

	currentX := x
	currentY := y
	switch dir {
	case 't':
		currentY--
	case 'r':
		currentX++
	case 'b':
		currentY++
	case 'l':
		currentX--
	}

	for currentX >= 0 && currentX <= maxX && currentY >= 0 && currentY <= maxY {
		char := matrix[currentY][currentX]
		if char != '.' {
			return &block{
				char: char,
				x:    currentX,
				y:    currentY,
			}
		}
		switch dir {
		case 't':
			currentY--
		case 'r':
			currentX++
		case 'b':
			currentY++
		case 'l':
			currentX--
		}
	}
	return nil
}

func getLoad(matrix [][]rune) int {
	totalLoad := 0
	height := len(matrix[0])
	for y, row := range matrix {
		for _, char := range row {
			if char == 'O' {
				totalLoad += height - y
			}
		}
	}

	return totalLoad
}

//func getOutput(matrix [][]rune) (output string) {
//	for _, row := range matrix {
//		for _, char := range row {
//			output += string(char)
//		}
//		output += "\n"
//	}
//	return output
//}

func copyMatrix(matrix [][]rune) (copiedMatrix [][]rune) {
	for y, row := range matrix {
		copiedMatrix = append(copiedMatrix, nil)
		for _, char := range row {
			copiedMatrix[y] = append(copiedMatrix[y], char)
		}
	}
	return copiedMatrix
}

func indexOfMatrix(matrices [][][]rune, matrix [][]rune) int {
	for i, matrix2 := range matrices {
		if isEquals(matrix, matrix2) {
			return i
		}
	}
	return -1
}

func isEquals(matrix1 [][]rune, matrix2 [][]rune) bool {
	if len(matrix1) != len(matrix2) || len(matrix1[0]) != len(matrix2[0]) {
		return false
	}
	for y, row := range matrix1 {
		for x, char := range row {
			if char != matrix2[y][x] {
				return false
			}
		}
	}
	return true
}
