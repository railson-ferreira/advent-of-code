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

type vector struct {
	x int8
	y int8
}

type pipe struct {
	coords          coordinates
	directionVector vector
	outsideVector   vector
	previousPipe    *pipe
	nextPipe        *pipe
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

	pipes := findPipes(matrix, x, y)

	removeAdditionalPipes(matrix, pipes)

	var isOutsideVectorClockwise bool
	for _, thePipe := range pipes {
		if thePipe.coords.y == 6 && thePipe.coords.x == 46 {
			if (thePipe.directionVector == vector{x: 1, y: -1}) {
				isOutsideVectorClockwise = false
			} else if (thePipe.directionVector == vector{x: -1, y: 1}) {
				isOutsideVectorClockwise = true
			}
		}
	}
	fillTheOutsideVector(pipes, isOutsideVectorClockwise)
	highLightOutsidePoints(matrix, pipes)
	propagateAtThoughtTheDots(matrix)

	_ = os.WriteFile("gold.txt", []byte(convertToString(matrix)), 0644)

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

func findPipes(matrix [][]string, x int, y int) (pipes []*pipe) {
	directions := []int32{'t', 'r', 'b', 'l'}
	notAllowedDirection := ' '
	finished := false
	pipes = append(pipes, &pipe{
		coords: coordinates{
			x: x,
			y: y,
		},
		directionVector: vector{},
		outsideVector:   vector{},
		nextPipe:        nil,
		previousPipe:    nil,
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
				firstPipe := pipes[0]
				lastPipe := pipes[len(pipes)-1]
				firstPipe.previousPipe = lastPipe
				lastPipe.nextPipe = pipes[0]
				lastPipe.directionVector = getVector(currentSymbol, string(dir))
			} else if nextSymbol != "" {
				previousPipe := pipes[len(pipes)-1]
				if currentSymbol != "S" {
					previousPipe.directionVector = getVector(currentSymbol, string(dir))
				}
				nextPipe := &pipe{
					coords: coordinates{
						x: x,
						y: y,
					},
					previousPipe: previousPipe,
				}
				previousPipe.nextPipe = nextPipe
				pipes = append(pipes, nextPipe)
				break
			}
		}
	}
	return pipes
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

func getVector(previousSymbol string, direction string) vector {

	switch previousSymbol + direction {
	case "|t":
		return vector{x: 0, y: -1}
	case "Jt":
		return vector{x: 1, y: -1}
	case "Lt":
		return vector{x: -1, y: -1}
	case "|b":
		return vector{x: 0, y: +1}
	case "7b":
		return vector{x: 1, y: +1}
	case "Fb":
		return vector{x: -1, y: +1}
	case "-l":
		return vector{x: -1, y: 0}
	case "Jl":
		return vector{x: -1, y: 1}
	case "7l":
		return vector{x: -1, y: -1}
	case "-r":
		return vector{x: +1, y: 0}
	case "Fr":
		return vector{x: +1, y: -1}
	case "Lr":
		return vector{x: +1, y: 1}
	}
	panic("Invalid combination" + fmt.Sprint(previousSymbol, direction))
}

func fillTheOutsideVector(pipes []*pipe, clockWise bool) {
	for _, thePipe := range pipes {
		if clockWise {
			panic("unimplemented")
		} else {
			if (thePipe.directionVector == vector{x: 1, y: -1}) {
				thePipe.outsideVector = vector{x: -1, y: -1}
			}
			if (thePipe.directionVector == vector{x: 1, y: 0}) {
				thePipe.outsideVector = vector{x: 0, y: -1}
			}
			if (thePipe.directionVector == vector{x: 1, y: 1}) {
				thePipe.outsideVector = vector{x: 1, y: -1}
			}
			if (thePipe.directionVector == vector{x: 0, y: 1}) {
				thePipe.outsideVector = vector{x: 1, y: 0}
			}
			if (thePipe.directionVector == vector{x: -1, y: 1}) {
				thePipe.outsideVector = vector{x: 1, y: 1}
			}
			if (thePipe.directionVector == vector{x: -1, y: 0}) {
				thePipe.outsideVector = vector{x: 0, y: 1}
			}
			if (thePipe.directionVector == vector{x: -1, y: -1}) {
				thePipe.outsideVector = vector{x: -1, y: 1}
			}
			if (thePipe.directionVector == vector{x: 0, y: -1}) {
				thePipe.outsideVector = vector{x: -1, y: 0}
			}
		}
	}
}

func removeAdditionalPipes(matrix [][]string, pipes []*pipe) {
	for y, row := range matrix {
		for x := range row {
			found := false
			for _, thePipe := range pipes {
				coords := thePipe.coords
				if coords.x == x && coords.y == y {
					found = true
					break
				}
			}
			if !found {
				matrix[y][x] = "."
			}
		}
	}
}
func highLightOutsidePoints(matrix [][]string, pipes []*pipe) {
	for _, thePipe := range pipes {
		outsideVector := thePipe.outsideVector
		coords := thePipe.coords
		targetX := coords.x + int(outsideVector.x)
		targetY := coords.y

		if targetY < 0 || targetY > len(matrix)-1 {
			continue
		}
		if targetX < 0 || targetX > len(matrix[0])-1 {
			continue
		}
		if matrix[targetY][targetX] == "." {
			matrix[targetY][targetX] = "@"
		}

		targetX = coords.x
		targetY = coords.y + int(outsideVector.y)
		if targetY < 0 || targetY > len(matrix)-1 {
			continue
		}
		if targetX < 0 || targetX > len(matrix[0])-1 {
			continue
		}
		if matrix[targetY][targetX] == "." {
			matrix[targetY][targetX] = "@"
		}

		targetX = coords.x + int(outsideVector.x)
		targetY = coords.y + int(outsideVector.y)
		if targetY < 0 || targetY > len(matrix)-1 {
			continue
		}
		if targetX < 0 || targetX > len(matrix[0])-1 {
			continue
		}
		if matrix[targetY][targetX] == "." {
			matrix[targetY][targetX] = "@"
		}
	}
}

func convertToString(matrix [][]string) (text string) {
	for _, row := range matrix {
		for _, letter := range row {
			text += letter
		}
		text += "\n"
	}
	return text
}

func propagateAtThoughtTheDots(matrix [][]string) {
	count := 1
	for count > 0 {
		count = 0
		for y, row := range matrix {
			for x := range row {
				if matrix[y][x] == "@" {
					nearbyDotCoords := getNearbyDotCoords(x, y, matrix)
					count += len(nearbyDotCoords)
					for _, coords := range nearbyDotCoords {
						matrix[coords.y][coords.x] = "@"
					}
				}
			}
		}

		fmt.Println(count)
	}
}

func getNearbyDotCoords(x int, y int, matrix [][]string) (neabyDotCoords []coordinates) {
	maxY := len(matrix) - 1
	maxX := len(matrix[0]) - 1

	for lx := x - 1; lx <= maxX && lx <= x+1; lx++ {
		if lx < 0 {
			continue
		}
		for ly := y - 1; ly <= maxY && ly <= y+1; ly++ {
			if ly < 0 {
				continue
			}
			letter := matrix[ly][lx]
			if letter == "." {
				neabyDotCoords = append(neabyDotCoords, coordinates{x: lx, y: ly})
			}
		}
	}
	return neabyDotCoords
}
