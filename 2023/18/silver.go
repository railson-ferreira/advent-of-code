package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type coordinates struct {
	x int
	y int
}

type direction rune

const (
	up    direction = 'U'
	down  direction = 'D'
	left  direction = 'L'
	right direction = 'R'
)

type instruction struct {
	dir   direction
	steps int
	color int
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

	instructions := formatInput(input)
	points := []coordinates{
		{x: 0, y: 0},
	}
	var lastDirection direction
	angle := 0
	for index, theInstruction := range instructions {
		lastPoint := points[len(points)-1]
		points = append(points, lastPoint.move(theInstruction.dir, theInstruction.steps)...)
		if index != 0 {
			angle += getAngle(lastDirection, theInstruction.dir)
		}
		lastDirection = theInstruction.dir
	}
	firstPoint := points[0]
	lastPoint := points[len(points)-1]
	difference := coordinates{
		x: lastPoint.x - firstPoint.x,
		y: lastPoint.y - firstPoint.y,
	}
	fmt.Println("first point:", firstPoint)
	fmt.Println("last point:", lastPoint)
	fmt.Println("difference:", difference)
	fmt.Println("direction:", string(lastDirection))
	fmt.Println("angle:", angle)
	fmt.Println("edge count:", len(points)-1)

	if angle < 0 || lastDirection != up {
		panic("Expected angle to be greater than 0 and lastDirection == 'U'")
	}
	preLastPoint := points[len(points)-2]
	startingPoint := coordinates{
		x: preLastPoint.x + 1,
		y: preLastPoint.y,
	}
	fill(startingPoint, &points)
	fmt.Println("edge count:", len(points)-1)

}

func formatInput(input string) (instructions []instruction) {
	lines := strings.Split(input, "\n")

	regex := regexp.MustCompile("(\\w) (\\d+) \\(#(.+)\\)")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		results := regex.FindStringSubmatch(line)
		dirText := results[1]
		stepsText := results[2]
		colorText := results[3]

		dir := direction(dirText[0])
		steps, _ := strconv.Atoi(stepsText)
		color, _ := strconv.ParseInt(colorText, 16, 32)

		instructions = append(instructions, instruction{
			dir:   dir,
			steps: steps,
			color: int(color),
		})

	}

	return
}

func (theCoordinates coordinates) move(dir direction, steps int) (result []coordinates) {
	x := theCoordinates.x
	y := theCoordinates.y
	switch dir {
	case up:
		for i := 0; i < steps; i++ {
			y--
			result = append(result, coordinates{
				x: x,
				y: y,
			})
		}
	case down:
		for i := 0; i < steps; i++ {
			y++
			result = append(result, coordinates{
				x: x,
				y: y,
			})
		}
	case left:
		for i := 0; i < steps; i++ {
			x--
			result = append(result, coordinates{
				x: x,
				y: y,
			})
		}
	case right:
		for i := 0; i < steps; i++ {
			x++
			result = append(result, coordinates{
				x: x,
				y: y,
			})
		}
	default:
		panic("Invalid direction")
	}
	return
}

func (theCoordinates coordinates) moveOne(dir direction) coordinates {
	x := theCoordinates.x
	y := theCoordinates.y
	switch dir {
	case up:
		y--
	case down:
		y++
	case left:
		x--
	case right:
		x++
	default:
		panic("Invalid direction")
	}
	return coordinates{
		x: x,
		y: y,
	}
}

func getAngle(from direction, to direction) int {
	if from == up && to == right {
		return 90
	}
	if from == up && to == left {
		return -90
	}
	if from == right && to == down {
		return 90
	}
	if from == right && to == up {
		return -90
	}
	if from == down && to == left {
		return 90
	}
	if from == down && to == right {
		return -90
	}
	if from == left && to == up {
		return 90
	}
	if from == left && to == down {
		return -90
	}

	panic(fmt.Sprintln("Invalid directions from:", string(from), "to:", string(to)))
}

func fill(startingPoint coordinates, points *[]coordinates) {
	canAdd := true
	for _, point := range *points {
		if point == startingPoint {
			canAdd = false
		}
	}
	if canAdd {
		*points = append(*points, startingPoint)
	} else {
		return
	}

	fourNewPositions := []coordinates{
		startingPoint.moveOne(up),
		startingPoint.moveOne(down),
		startingPoint.moveOne(left),
		startingPoint.moveOne(right),
	}

	for _, position := range fourNewPositions {
		fill(position, points)
	}
}
