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

	instructions, aaaNode := formatInput(input)

	currentNode := aaaNode
	stepsSum := 0
	for currentNode.value != "ZZZ" {
		for _, instruction := range instructions {
			currentNode = currentNode.step(direction(instruction))
			stepsSum++
			if currentNode.value == "ZZZ" {
				break
			}
		}
	}

	fmt.Println(stepsSum)

}

func formatInput(input string) (instructions []string, aaaNode *node) {
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

		if source == "AAA" {
			aaaNode = nodes[source]
			continue
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

	return instructions, aaaNode
}

func (theNode node) step(dir direction) *node {
	switch dir {
	case leftDir:
		fmt.Println(dir, theNode.value, "->", theNode.left.value)
		return theNode.left
	case rightDir:
		fmt.Println(dir, theNode.value, "->", theNode.right.value)
		return theNode.right
	}
	panic("Invalid direction")
}
