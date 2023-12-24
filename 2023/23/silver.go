package main

import (
	"fmt"
	"os"
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

	theMap := formatInput(input)

	startingPoint := [2]int{1, 0}
	endingPoint := [2]int{len(theMap[0]) - 2, len(theMap) - 1}

	undonePaths := [][]*[2]int{{&startingPoint}}
	var donePaths [][]*[2]int
	count := 0
	for len(undonePaths) > 0 {
		count++
		copyUndonePaths := undonePaths
		undonePaths = make([][]*[2]int, 0)
		for _, undonePath := range copyUndonePaths {
			nextPaths := getPaths(undonePath, theMap, &endingPoint)

			for _, np := range nextPaths {
				nexPath := np
				if *nexPath[len(nexPath)-1] == endingPoint {
					donePaths = append(donePaths, nexPath)
				} else {
					undonePaths = append(undonePaths, nexPath)
				}
			}
		}
	}

	fmt.Println("DONE:", len(donePaths))
	maxSteps := 0
	for _, donePath := range donePaths {
		steps := len(donePath) - 1
		if steps > maxSteps {
			maxSteps = steps
		}
	}
	fmt.Println("maxSteps", maxSteps)
}

func formatInput(input string) (theMap [][]rune) {
	lines := strings.Split(input, "\n")

	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		theMap = append(theMap, nil)
		for _, char := range line {
			theMap[y] = append(theMap[y], char)
		}
	}
	return
}

func getPaths(path []*[2]int, theMap [][]rune, endingPoint *[2]int) (nextPaths [][]*[2]int) {
	pathHead := path[len(path)-1]
	if *pathHead == *endingPoint {
		panic("Should never be reached")
	}
	var beforeHead *[2]int
	if len(path) > 1 {
		beforeHead = path[len(path)-2]
	}
	up := [2]int{pathHead[0], pathHead[1] - 1}
	right := [2]int{pathHead[0] + 1, pathHead[1]}
	down := [2]int{pathHead[0], pathHead[1] + 1}
	left := [2]int{pathHead[0] - 1, pathHead[1]}
	maxX := len(theMap[0]) - 1
	maxY := len(theMap) - 1
	for index, t := range [][2]int{up, right, down, left} {
		target := [2]int{t[0], t[1]}
		x := target[0]
		y := target[1]
		if x >= 0 && x <= maxX && y >= 0 && y <= maxY {
			if beforeHead == nil || target != *beforeHead {
				mapTile := theMap[y][x]
				if mapTile == '.' {
					nextPaths = append(nextPaths, append(path, &target))
				} else {
					switch theMap[y][x] {
					case '^':
						if index == 0 {
							nextPaths = append(nextPaths, append(append(make([]*[2]int, 0), path...), &target, &[2]int{target[0], target[1] - 1}))
						}
					case '>':
						if index == 1 {
							nextPaths = append(nextPaths, append(append(make([]*[2]int, 0), path...), &target, &[2]int{target[0] + 1, target[1]}))
						}
					case 'v':
						if index == 2 {
							nextPaths = append(nextPaths, append(append(make([]*[2]int, 0), path...), &target, &[2]int{target[0], target[1] + 1}))
						}
					case '<':
						if index == 3 {
							nextPaths = append(nextPaths, append(append(make([]*[2]int, 0), path...), &target, &[2]int{target[0] - 1, target[1]}))
						}

					}
				}
			}
		}
	}
	return
}
