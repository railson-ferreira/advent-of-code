package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type direction int8

const (
	up direction = iota + 1
	right
	down
	left
)

type pathNode struct {
	previous     *pathNode
	x            int
	y            int
	dir          direction
	heatLossSum  int
	sameDirCount int
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

	path := getOptimizedPath(matrix)
	if path == nil {
		fmt.Printf("Path is nil")
	} else {
		currentHead := path
		var fullPath []pathNode
		for currentHead != nil {
			fullPath = append(fullPath, *currentHead)
			currentHead = currentHead.previous
		}
		slices.Reverse(fullPath)
		fmt.Println()
		fmt.Println()
		lastLossSum := 0
		for _, node := range fullPath {
			fmt.Print(node.heatLossSum - lastLossSum)
			lastLossSum = node.heatLossSum
		}
		fmt.Println(" ", "Sum:", path.heatLossSum)
		for _, node := range fullPath {
			switch node.dir {
			case up:
				fmt.Print("↑")
			case right:
				fmt.Print("→")
			case down:
				fmt.Print("↓")
			case left:
				fmt.Print("←")
			}
		}
		fmt.Println()
	}

}

func formatInput(input string) (matrix [][]int) {
	lines := strings.Split(input, "\n")

	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		matrix = append(matrix, nil)
		for _, char := range line {
			v, _ := strconv.Atoi(string(char))
			matrix[y] = append(matrix[y], v)
		}
	}

	return
}

var initialPath = pathNode{
	heatLossSum: 9999999999999999,
}
var bestPath = &initialPath

func getOptimizedPath(matrix [][]int) *pathNode {

	currentPaths := []*pathNode{
		{
			dir:          right,
			x:            0,
			y:            0,
			sameDirCount: 1,
			heatLossSum:  0,
		},
		{
			dir:          down,
			x:            0,
			y:            0,
			sameDirCount: 1,
			heatLossSum:  0,
		},
	}

	handlePaths(currentPaths, matrix)
	if *bestPath != initialPath {
		return bestPath
	}
	return nil
}

func getPathsToTheEnd(pathHead *pathNode, matrix [][]int) *pathNode {
	width, height := len(matrix[0]), len(matrix)
	heatLossSum, sameDirCount := pathHead.heatLossSum, pathHead.sameDirCount
	if pathHead.x == width-1 && pathHead.y == height-1 {
		return pathHead
	}

	if sameDirCount < 3 {
		x, y := move(pathHead.dir, pathHead.x, pathHead.y)
		if x >= 0 && x < width && y >= 0 && y < height {
			newNode := &pathNode{
				previous:     pathHead,
				dir:          pathHead.dir,
				x:            x,
				y:            y,
				sameDirCount: sameDirCount + 1,
				heatLossSum:  matrix[y][x] + heatLossSum,
			}
			if !isCrossingPath(newNode) {
				if newNode.heatLossSum < bestPath.heatLossSum {
					pathToTheEnd := getPathsToTheEnd(newNode, matrix)
					if pathToTheEnd != nil && pathToTheEnd.heatLossSum < bestPath.heatLossSum {
						fmt.Println("new best: ", newNode)
						bestPath = newNode
					}
				}
			}
		}
	}
	for _, dir := range getPerpendicularDirections(pathHead.dir) {
		x, y := move(dir, pathHead.x, pathHead.y)

		if x >= 0 && x < width && y >= 0 && y < height {
			newNode := &pathNode{
				previous:     pathHead,
				dir:          dir,
				x:            x,
				y:            y,
				sameDirCount: 1,
				heatLossSum:  matrix[y][x] + heatLossSum,
			}
			if !isCrossingPath(newNode) {
				if newNode.heatLossSum < bestPath.heatLossSum {
					pathToTheEnd := getPathsToTheEnd(newNode, matrix)
					if pathToTheEnd != nil && pathToTheEnd.heatLossSum < bestPath.heatLossSum {
						fmt.Println("new best: ", newNode)
						bestPath = newNode
					}
				}
			}
		}
	}

	return nil
}

func getPerpendicularDirections(dir direction) []direction {
	switch dir {
	case up, down:
		return []direction{right, left}
	case right, left:
		return []direction{down, up}
	}
	panic("Invalid direction")
}

func move(dir direction, x int, y int) (_x int, _y int) {
	switch dir {
	case up:
		return x, y - 1
	case right:
		return x + 1, y
	case down:
		return x, y + 1
	case left:
		return x - 1, y
	}
	panic("Invalid direction")
}

func handlePaths(pathHeads []*pathNode, matrix [][]int) (bestPath *pathNode) {

	for _, pathHead := range pathHeads {
		getPathsToTheEnd(pathHead, matrix)
	}
	return
}

func isCrossingPath(pathHead *pathNode) bool {
	currentPath := pathHead.previous
	for currentPath != nil {
		if currentPath.x == pathHead.x && currentPath.y == pathHead.y {
			return true
		}
		currentPath = currentPath.previous
	}
	return false
}
