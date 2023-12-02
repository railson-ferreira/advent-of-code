package main

import (
	"fmt"
	"os"
	"strconv"
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

	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}

		sum += getCalibrationValue(line)
	}

	fmt.Println(sum)
}

func getCalibrationValue(line string) int {
	firstDigit := -1
	firstDigitIndex := -1
	lastDigit := -1
	lastDigitIndex := -1

	for index, letter := range line {
		if digit, err := strconv.Atoi(string(letter)); err == nil {
			firstDigit = digit
			firstDigitIndex = index
			break
		}
	}

	for index, letter := range reverse(line) {
		if digit, err := strconv.Atoi(string(letter)); err == nil {
			lastDigit = digit
			lastDigitIndex = len(line) - 1 - index
			break
		}
	}

	firstTextDigit, firstTextDigitIndex := findFirstTextDigit(line)
	lastTextDigit, lastTextDigitIndex := findLastTextDigit(line)

	if firstTextDigitIndex != -1 && firstTextDigitIndex < firstDigitIndex {
		firstDigit = firstTextDigit
	}
	if lastTextDigitIndex != -1 && lastTextDigitIndex > lastDigitIndex {
		lastDigit = lastTextDigit
	}
	valueAsString := fmt.Sprintf("%v%v", firstDigit, lastDigit)
	value, err := strconv.Atoi(valueAsString)
	if err != nil {
		fmt.Println(valueAsString)
	}
	check(err)
	return value
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func findFirstTextDigit(text string) (digit int, index int) {
	digit = -1
	index = -1
	if lIndex := strings.Index(text, "zero"); lIndex != -1 {
		if lIndex < index || index == -1 {
			digit = 0
			index = lIndex
		}
	}
	if lIndex := strings.Index(text, "one"); lIndex != -1 {
		if lIndex < index || index == -1 {
			digit = 1
			index = lIndex
		}
	}
	if lIndex := strings.Index(text, "two"); lIndex != -1 {
		if lIndex < index || index == -1 {
			digit = 2
			index = lIndex
		}
	}
	if lIndex := strings.Index(text, "three"); lIndex != -1 {
		if lIndex < index || index == -1 {
			digit = 3
			index = lIndex
		}
	}
	if lIndex := strings.Index(text, "four"); lIndex != -1 {
		if lIndex < index || index == -1 {
			digit = 4
			index = lIndex
		}
	}
	if lIndex := strings.Index(text, "five"); lIndex != -1 {
		if lIndex < index || index == -1 {
			digit = 5
			index = lIndex
		}
	}
	if lIndex := strings.Index(text, "six"); lIndex != -1 {
		if lIndex < index || index == -1 {
			digit = 6
			index = lIndex
		}
	}
	if lIndex := strings.Index(text, "seven"); lIndex != -1 {
		if lIndex < index || index == -1 {
			digit = 7
			index = lIndex
		}
	}
	if lIndex := strings.Index(text, "eight"); lIndex != -1 {
		if lIndex < index || index == -1 {
			digit = 8
			index = lIndex
		}
	}
	if lIndex := strings.Index(text, "nine"); lIndex != -1 {
		if lIndex < index || index == -1 {
			digit = 9
			index = lIndex
		}
	}

	return digit, index
}

func findLastTextDigit(text string) (digit int, index int) {
	digit = -1
	index = -1
	if lIndex := strings.LastIndex(text, "zero"); lIndex != -1 {
		if lIndex > index {
			digit = 0
			index = lIndex
		}
	}
	if lIndex := strings.LastIndex(text, "one"); lIndex != -1 {
		if lIndex > index {
			digit = 1
			index = lIndex
		}
	}
	if lIndex := strings.LastIndex(text, "two"); lIndex != -1 {
		if lIndex > index {
			digit = 2
			index = lIndex
		}
	}
	if lIndex := strings.LastIndex(text, "three"); lIndex != -1 {
		if lIndex > index {
			digit = 3
			index = lIndex
		}
	}
	if lIndex := strings.LastIndex(text, "four"); lIndex != -1 {
		if lIndex > index {
			digit = 4
			index = lIndex
		}
	}
	if lIndex := strings.LastIndex(text, "five"); lIndex != -1 {
		if lIndex > index {
			digit = 5
			index = lIndex
		}
	}
	if lIndex := strings.LastIndex(text, "six"); lIndex != -1 {
		if lIndex > index {
			digit = 6
			index = lIndex
		}
	}
	if lIndex := strings.LastIndex(text, "seven"); lIndex != -1 {
		if lIndex > index {
			digit = 7
			index = lIndex
		}
	}
	if lIndex := strings.LastIndex(text, "eight"); lIndex != -1 {
		if lIndex > index {
			digit = 8
			index = lIndex
		}
	}
	if lIndex := strings.LastIndex(text, "nine"); lIndex != -1 {
		if lIndex > index {
			digit = 9
			index = lIndex
		}
	}

	return digit, index
}
