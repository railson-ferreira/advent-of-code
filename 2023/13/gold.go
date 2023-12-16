package main

import (
	"fmt"
	"math"
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

	groupsInputs := formatInput(input)

	sum := 0
	for _, groupInput := range groupsInputs {
		group := formatGroupInput(groupInput)
		groupValue := getGroupValue(group)
		sum += groupValue
	}

	fmt.Println(sum)

}

func formatInput(input string) (groups []string) {
	lines := strings.Split(input, "\n")

	count := 0
	groups = append(groups, "")
	for i, line := range lines {
		if len(line) == 0 {
			if i < len(lines)-1 {
				groups = append(groups, "")
				count++
			}
			continue
		}
		groups[count] += line + "\n"
	}

	return groups
}

func formatGroupInput(groupInput string) (matrix [][]rune) {

	lines := strings.Split(groupInput, "\n")

	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		matrix = append(matrix, nil)
		for _, char := range line {
			if len(line) == 0 {
				continue
			}
			matrix[y] = append(matrix[y], char)
		}
	}
	return matrix
}

func getGroupValue(group [][]rune) (result int) {
	offset := 9999.0
	count := 0
	for count < len(group)*len(group[0]) {
		modified := getModifiedGroup(count, group)
		count++

		// look for horizontal lines
		for y := range modified {
			if y == 0 {
				continue
			}
			if isBothEquals(modified[y-1], modified[y]) {
				if ensureHorizontalLineReflection(y, modified) {
					lOffset := math.Abs(float64(y - len(modified)/2))
					if lOffset < offset {
						if isBothEquals(group[y-1], group[y]) && ensureHorizontalLineReflection(y, group) {
							continue
						}
						offset = lOffset
						result = y * 100
					}
				}
			}
		}

		// look for vertical lines
		for x := range modified[0] {
			if x == 0 {
				continue
			}
			array1 := getColumn(x-1, modified)
			array2 := getColumn(x, modified)
			if isBothEquals(array1, array2) {
				if ensureVerticalLineReflection(x, modified) {
					lOffset := math.Abs(float64(x - len(modified[0])/2))
					if lOffset < offset {
						if isBothEquals(getColumn(x-1, group), getColumn(x, group)) && ensureVerticalLineReflection(x, group) {
							continue
						}
						offset = lOffset
						result = x
					}
				}
			}
		}
	}
	if result == 0 {
		panic("No equals rows or columns were found")
	}
	return result
}

func isBothEquals(array1 []rune, array2 []rune) bool {
	if len(array1) != len(array2) {
		return false
	}
	for i, v := range array1 {
		if v != array2[i] {
			return false
		}
	}
	return true
}

func getColumn(columnIndex int, group [][]rune) []rune {
	var array []rune
	for _, row := range group {
		array = append(array, row[columnIndex])
	}
	return array
}

func ensureHorizontalLineReflection(reflexIndex int, group [][]rune) bool {
	y1 := reflexIndex - 1
	y2 := reflexIndex
	for y1 >= 0 && y2 <= len(group)-1 {
		if !isBothEquals(group[y1], group[y2]) {
			return false
		}
		y1--
		y2++
	}
	return true
}

func ensureVerticalLineReflection(reflexIndex int, group [][]rune) bool {
	x1 := reflexIndex - 1
	x2 := reflexIndex
	for x1 >= 0 && x2 <= len(group[0])-1 {
		if !isBothEquals(getColumn(x1, group), getColumn(x2, group)) {
			return false
		}
		x1--
		x2++
	}
	return true
}

func getModifiedGroup(index int, group [][]rune) (copy [][]rune) {
	for y, row := range group {
		copy = append(copy, nil)
		for _, v := range row {
			copy[y] = append(copy[y], v)
		}
	}
	y := 0
	length := len(group[0])
	// It could be O(1) with a simple calculation
	for index >= length {
		index -= length
		y++
	}
	char := group[y][index]
	if char == '.' {
		copy[y][index] = '#'
	} else {
		copy[y][index] = '.'
	}
	return copy
}
