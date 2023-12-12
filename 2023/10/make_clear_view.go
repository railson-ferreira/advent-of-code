package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

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

	matrix, x, y := formatInput(input)

	pipeCoordinates := findPipeCoordinates(matrix, x, y)

	fmt.Println(pipeCoordinates)

	_ = os.WriteFile("view.txt", []byte(highlightPipe(matrix, pipeCoordinates)), 0644)

}

func formatInput(input string) (matrix [][]string, initialX int, initialY int) {
	lines := strings.Split(input, "\n")

	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		matrix = append(matrix, []string{})

		for x, char := range line {
			letter := string(char)
			if letter == "S" {
				initialX = x
				initialY = y
			}
			matrix[y] = append(matrix[y], letter)
		}

	}

	return matrix, initialX, initialY
}

func findPipeCoordinates(matrix [][]string, x int, y int) (pipeCoordinates []coordinates) {
	directions := []int32{'t', 'r', 'b', 'l'}
	notAllowedDirection := ' '
	finished := false
	pipeCoordinates = append(pipeCoordinates, coordinates{
		x: x,
		y: y,
	})
	for !finished {
		for _, dir := range directions {
			if dir == notAllowedDirection {
				continue
			}
			currentSymbol := matrix[y][x]
			var nextSymbol string
			switch dir {
			case 't':
				if y == 0 {
					continue
				}
				try := matrix[y-1][x]
				if doesFit(dir, currentSymbol, try) {
					y--
					nextSymbol = try
					notAllowedDirection = 'b'
				}
			case 'r':
				if x == len(matrix[0])-1 {
					continue
				}
				try := matrix[y][x+1]
				if doesFit(dir, currentSymbol, try) {
					x++
					nextSymbol = try
					notAllowedDirection = 'l'
				}
			case 'b':
				if y == len(matrix)-1 {
					continue
				}
				try := matrix[y+1][x]
				if doesFit(dir, currentSymbol, try) {
					y++
					nextSymbol = try
					notAllowedDirection = 't'
				}
			case 'l':
				if x == 0 {
					continue
				}
				try := matrix[y][x-1]
				if doesFit(dir, currentSymbol, try) {
					x--
					nextSymbol = try
					notAllowedDirection = 'r'
				}
			}
			if nextSymbol == "S" {
				finished = true
			} else if nextSymbol != "" {
				pipeCoordinates = append(pipeCoordinates, coordinates{
					x: x,
					y: y,
				})
				break
			}
		}
	}
	return pipeCoordinates
}

func doesFit(direction int32, currentSymbol string, nextSymbol string) bool {
	var allowedCurrentSymbols []string
	var allowedNextSymbols []string
	switch direction {
	case 't':
		allowedCurrentSymbols = []string{"S", "|", "J", "L"}
		allowedNextSymbols = []string{"S", "|", "7", "F"}
	case 'r':
		allowedCurrentSymbols = []string{"S", "-", "L", "F"}
		allowedNextSymbols = []string{"S", "-", "7", "J"}
	case 'b':
		allowedCurrentSymbols = []string{"S", "|", "7", "F"}
		allowedNextSymbols = []string{"S", "|", "J", "L"}
	case 'l':
		allowedCurrentSymbols = []string{"S", "-", "7", "J"}
		allowedNextSymbols = []string{"S", "-", "L", "F"}
	}
	return slices.Contains(allowedCurrentSymbols, currentSymbol) && slices.Contains(allowedNextSymbols, nextSymbol)

}

func highlightPipe(matrix [][]string, pipeCoordinates []coordinates) (highlighted string) {
	for y, row := range matrix {
		for x, letter := range row {
			found := false
			for _, coord := range pipeCoordinates {
				if coord.x == x && coord.y == y {
					highlighted += letter
					found = true
					break
				}
			}
			if !found {
				highlighted += " "
			}
		}
		highlighted += "\n"
	}
	return highlighted
}
