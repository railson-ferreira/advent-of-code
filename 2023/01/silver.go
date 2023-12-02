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
	lastDigit := -1

	for _, letter := range line {
		if digit, err := strconv.Atoi(string(letter)); err == nil {
			firstDigit = digit
			break
		}
	}

	for _, letter := range reverse(line) {
		if digit, err := strconv.Atoi(string(letter)); err == nil {
			lastDigit = digit
			break
		}
	}

	valueAsString := fmt.Sprintf("%v%v", firstDigit, lastDigit)
	var value, err = strconv.Atoi(valueAsString)
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
