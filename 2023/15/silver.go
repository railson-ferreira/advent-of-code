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

	steps := formatInput(input)

	sum := 0
	for _, step := range steps {
		currentValue := 0
		for _, letter := range strings.Split(step, "") {
			currentValue += int(letter[0])
			currentValue *= 17
			currentValue %= 256
		}
		sum += currentValue
	}
	fmt.Println(sum)
}

func formatInput(input string) (steps []string) {
	lines := strings.Split(input, ",")

	for i, step := range lines {
		if len(step) == 0 {
			continue
		}
		if i == len(lines)-1 {
			steps = append(steps, strings.Replace(step, "\n", "", -1))
		} else {
			steps = append(steps, step)
		}
	}

	return steps
}
