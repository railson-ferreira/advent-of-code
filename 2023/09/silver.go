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

	sequences := formatInput(input)

	sum := 0
	for _, sequence := range sequences {
		sum += findTheExtrapolateNumber(sequence)
	}

	fmt.Println(sum)

}

func formatInput(input string) (sequences [][]int) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		var sequence []int
		for _, value := range strings.Split(line, " ") {
			valueInt, _ := strconv.Atoi(value)
			sequence = append(sequence, valueInt)
		}

		sequences = append(sequences, sequence)

	}

	return sequences
}

func getDiffSequence(sequence []int) (diffSequence []int) {
	for index := range sequence {
		if index == 0 {
			continue
		}
		diffSequence = append(diffSequence, sequence[index]-sequence[index-1])
	}
	return diffSequence
}

func areAllZeros(sequence []int) bool {
	for _, value := range sequence {
		if value != 0 {
			return false
		}
	}
	return true
}

func findTheExtrapolateNumber(sequence []int) int {
	diffSequence := getDiffSequence(sequence)
	if areAllZeros(diffSequence) {
		return sequence[len(sequence)-1]
	}
	diffExtrapolateNumber := findTheExtrapolateNumber(diffSequence)

	return sequence[len(sequence)-1] + diffExtrapolateNumber
}
