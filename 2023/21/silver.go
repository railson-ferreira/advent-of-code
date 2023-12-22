package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	bytes, err := os.ReadFile("input.txt")
	check(err)
	input := string(bytes)

	startingX, startingY, theMap := formatInput(input)

	previousPoints := [][2]int{{startingX, startingY}}

	goal := 64
	for i := 0; i < goal; i++ {
		points := append(previousPoints)
		previousPoints = make([][2]int, 0)
		for _, point := range points {
			x := point[0]
			y := point[1]
			targets := tryToMoveAllDirections(x, y, theMap)
			for _, targetRef := range targets {
				target := *targetRef
				if !slices.Contains(previousPoints, target) {
					previousPoints = append(previousPoints, target)
				}
			}
		}
	}
	fmt.Println(len(previousPoints))
}

func formatInput(input string) (startingX int, startingY int, theMap [][]bool) {
	lines := strings.Split(input, "\n")

	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		theMap = append(theMap, nil)
		for x, char := range line {
			if char == '#' {
				theMap[y] = append(theMap[y], false)
			} else {
				theMap[y] = append(theMap[y], true)
			}
			if char == 'S' {
				startingX = x
				startingY = y
			}
		}
	}
	return
}

func tryToMoveAllDirections(x int, y int, theMap [][]bool) (result []*[2]int) {
	maxX := len(theMap[0]) - 1
	maxY := len(theMap) - 1

	up := [2]int{x, y - 1}
	down := [2]int{x, y + 1}
	left := [2]int{x - 1, y}
	right := [2]int{x + 1, y}

	for _, point := range [][2]int{up, down, left, right} {
		copyPoint := point
		targetX := point[0]
		targetY := point[1]
		if targetX >= 0 && targetX <= maxX && targetY >= 0 && targetY <= maxY {
			if theMap[targetY][targetX] { // if it's an allowed position
				result = append(result, &copyPoint)
			}
		}
	}
	return
}
