package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type step struct {
	label       string
	operation   rune
	focalLength int8
}

type lens struct {
	label       string
	focalLength int8
}

type box struct {
	lensList []lens
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

	steps := formatInput(input)

	boxes := make([]box, 256)

	for _, theStep := range steps {
		currentValue := 0
		for _, letter := range strings.Split(theStep.label, "") {
			currentValue += int(letter[0])
			currentValue *= 17
			currentValue %= 256
		}
		theBox := boxes[currentValue]
		switch theStep.operation {
		case '-':
			boxes[currentValue] = theBox.removeLens(theStep.label)
		case '=':
			boxes[currentValue] = theBox.replaceOrInsertLens(lens{
				label:       theStep.label,
				focalLength: theStep.focalLength,
			})
		default:
			panic("Invalid operation")
		}

	}
	sum := 0
	for i, theBox := range boxes {
		for slotIndex, theLens := range theBox.lensList {
			focusingPower := (i + 1) * (slotIndex + 1) * int(theLens.focalLength)
			sum += focusingPower
		}
	}
	fmt.Println(sum)
}

func formatInput(input string) (steps []step) {
	lines := strings.Split(input, ",")

	for i, stepText := range lines {
		if len(stepText) == 0 {
			continue
		}
		if i == len(lines)-1 {
			stepText = strings.Replace(stepText, "\n", "", -1)
		}
		if strings.Contains(stepText, "-") {
			steps = append(steps, step{
				label:     stepText[:len(stepText)-1],
				operation: '-',
			})
		} else {
			split := strings.Split(stepText, "=")
			focalLength, _ := strconv.Atoi(split[1])
			steps = append(steps, step{
				label:       split[0],
				operation:   '=',
				focalLength: int8(focalLength),
			})
		}
	}

	return steps
}

func (theBox box) removeLens(label string) box {
	for index, theLens := range theBox.lensList {
		if theLens.label == label {
			theBox.lensList = removeLens(theBox.lensList, index)
			break
		}
	}
	return theBox
}

func (theBox box) replaceOrInsertLens(theLens lens) box {
	for index, _theLens := range theBox.lensList {
		if _theLens.label == theLens.label {
			theBox.lensList[index] = theLens
			return theBox
		}
	}
	theBox.lensList = append(theBox.lensList, theLens)
	return theBox
}

// https://stackoverflow.com/a/37335777
func removeLens(lensList []lens, lensIndex int) []lens {
	return append(lensList[:lensIndex], lensList[lensIndex+1:]...)
}
