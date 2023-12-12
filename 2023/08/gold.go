package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type direction string

const (
	leftDir  direction = "L"
	rightDir direction = "R"
)

type node struct {
	value string
	left  *node
	right *node
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

	instructions, endsWithANodes := formatInput(input)

	currentNodes := endsWithANodes
	stepsSum := 0
	for i, v := range currentNodes {
		if v == nil {
			fmt.Print(i, "-", "___", " ")
		} else {
			fmt.Print(i, "-", v.value, " ")
		}
	}
	fmt.Println(" -> ", stepsSum)

	lastSteps := []int{0, 0, 0, 0, 0, 0}

	for !allDifferentOfZero(lastSteps) {
		for _, instruction := range instructions {
			var oldOnes []string
			for index, currentNode := range currentNodes {
				oldOnes = append(oldOnes, currentNode.value)
				currentNodes[index] = currentNode.step(direction(instruction))
				if finishWithZ(currentNodes[index]) {
					fmt.Println(index, stepsSum+1-lastSteps[index])
					lastSteps[index] = stepsSum + 1
				}
			}
			stepsSum++
			if allDifferentOfZero(lastSteps) {
				break
			}
		}
	}

	fmt.Println("LCM:", LCM(lastSteps[0], lastSteps[1], lastSteps[2:]...))

}

func formatInput(input string) (instructions []string, endsWithANodes []*node) {
	lines := strings.Split(input, "\n")

	instructions = strings.Split(lines[0], "")

	elementsRegex := regexp.MustCompile("(.+) = \\((.+), (.+)\\)")

	nodes := make(map[string]*node)
	for _, line := range lines[2:] {
		if len(line) == 0 {
			continue
		}
		result := elementsRegex.FindStringSubmatch(line)
		source := result[1]

		nodes[source] = &node{
			value: source,
		}

		if source[2] == 'A' {
			endsWithANodes = append(endsWithANodes, nodes[source])
		}
	}

	for _, line := range lines[2:] {
		if len(line) == 0 {
			continue
		}
		result := elementsRegex.FindStringSubmatch(line)
		source := result[1]
		left := result[2]
		right := result[3]

		nodes[source].left = nodes[left]
		nodes[source].right = nodes[right]
	}

	return instructions, endsWithANodes
}

func (theNode node) step(dir direction) *node {
	switch dir {
	case leftDir:
		return theNode.left
	case rightDir:
		return theNode.right
	}
	panic("Invalid direction")
}
func finishWithZ(theNode *node) bool {
	return theNode.value[2] == 'Z'
}

func allDifferentOfZero(numbers []int) bool {
	return all(numbers, func(number int) bool {
		return number != 0
	})
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// https://stackoverflow.com/a/75435478
func all[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if !pred(t) {
			return false
		}
	}
	return true
}
