package main

import (
	"fmt"
	"os"
	"strings"
)

type coordinates struct {
	x int
	y int
}

type roundedRock coordinates
type cubeRock coordinates

type scene struct {
	width        int
	height       int
	roundedRocks []*roundedRock
	cubeRocks    []*cubeRock
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

	theScene := formatInput(input)

	fmt.Println("Total load before tilt:", theScene.getLoad())
	theScene.tilt()
	fmt.Println("Total load after tilt:", theScene.getLoad())
	//_ = os.WriteFile("view.txt", []byte(theScene.getOutput()), 0644)
}

func formatInput(input string) (theScene scene) {
	lines := strings.Split(input, "\n")

	theScene.width = len(lines[0])
	theScene.height = len(lines) - 1
	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		for x, char := range line {
			switch char {
			case '.':
				break
			case '#':
				theScene.cubeRocks = append(theScene.cubeRocks, &cubeRock{
					x: x,
					y: y,
				})
			case 'O':
				theScene.roundedRocks = append(theScene.roundedRocks, &roundedRock{
					x: x,
					y: y,
				})
			default:
				panic("Invalid character")
			}
		}
	}

	return theScene
}

func (theScene scene) tilt() {
	// Considering that it's already ordered from the top most rock to the top lesser rock
	for _, rock := range theScene.roundedRocks {
		aboveRocks := theScene.getAllRocksFromColumnAndAboveLine(rock.x, rock.y)
		if len(aboveRocks) == 0 {
			rock.y = 0
			continue
		}
		topLesserOfAboveRocks := getTopLesser(aboveRocks)
		distance := rock.y - topLesserOfAboveRocks.y
		rock.y = rock.y - (distance - 1)
	}
}

func (theScene scene) getAllRocksFromColumnAndAboveLine(column int, line int) (rocks []coordinates) {
	for _, rock := range theScene.cubeRocks {
		if rock.x == column && rock.y == line {
			continue
		}
		if rock.x == column && rock.y < line {
			rocks = append(rocks, coordinates(*rock))
		}
	}
	for _, rock := range theScene.roundedRocks {
		if rock.x == column && rock.y == line {
			continue
		}
		if rock.x == column && rock.y < line {
			rocks = append(rocks, coordinates(*rock))
		}
	}
	return rocks
}

func getTopLesser(rocks []coordinates) coordinates {
	topLesser := rocks[0]
	for _, rock := range rocks {
		if rock.y > topLesser.y {
			topLesser = rock
		}
	}
	return topLesser
}

func (theScene scene) getLoad() int {
	totalLoad := 0
	for _, rock := range theScene.roundedRocks {
		totalLoad += theScene.height - rock.y
	}

	return totalLoad
}

//func (theScene scene) getOutput() (output string) {
//	for y := 0; y < theScene.height; y++ {
//		for x := 0; x < theScene.width; x++ {
//			output += string(theScene.getCharacter(x, y))
//		}
//		output += "\n"
//	}
//	return output
//}
//
//func (theScene scene) getCharacter(x int, y int) rune {
//	for _, rock := range theScene.roundedRocks {
//		if rock.x == x && rock.y == y {
//			return 'O'
//		}
//	}
//	for _, rock := range theScene.cubeRocks {
//		if rock.x == x && rock.y == y {
//			return '#'
//		}
//	}
//
//	return '.'
//}
