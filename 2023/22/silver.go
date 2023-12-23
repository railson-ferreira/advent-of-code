package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type block struct {
	min        [3]int
	max        [3]int
	supports   []*block
	supporting []*block
}

var ground = &block{
	min: [3]int{0, 0, 0},
	max: [3]int{0, 0, 0},
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

	blocks := formatInput(input)
	slices.SortFunc(blocks, func(a *block, b *block) int {
		return a.min[2] - b.min[2]
	})

	for _, currentBlock := range blocks {
		currentBlock.settle(blocks)
	}

	sum := 0
	for _, theBlock := range blocks {
		if theBlock.canBeSafelyDisintegrated() {
			sum++
		}
	}
	fmt.Println(sum)
}

func formatInput(input string) (blocks []*block) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, "~")
		theBlock := block{
			min: [3]int{},
			max: [3]int{},
		}
		aParts := strings.Split(parts[0], ",")
		for i, numberText := range aParts {
			theBlock.min[i], _ = strconv.Atoi(numberText)
		}
		bParts := strings.Split(parts[1], ",")
		for i, numberText := range bParts {
			theBlock.max[i], _ = strconv.Atoi(numberText)
		}
		blocks = append(blocks, &theBlock)
	}
	return
}

func (currentBlock *block) settle(blocks []*block) {
	zOffset := 1
	hitBlocks := getHitBlock(moveZ(zOffset, currentBlock.min), moveZ(zOffset, currentBlock.max), blocks, currentBlock)
	for len(hitBlocks) == 0 {
		zOffset++
		hitBlocks = getHitBlock(moveZ(zOffset, currentBlock.min), moveZ(zOffset, currentBlock.max), blocks, currentBlock)
	}
	currentBlock.min = moveZ(zOffset-1, currentBlock.min)
	currentBlock.max = moveZ(zOffset-1, currentBlock.max)
	for _, hitBlock := range hitBlocks {
		hitBlock.supporting = append(hitBlock.supporting, currentBlock)
		currentBlock.supports = append(currentBlock.supports, hitBlock)
	}
}

func getHitBlock(min [3]int, max [3]int, blocks []*block, ignoredBlock *block) (hitBoxes []*block) {
	for _, theBlock := range blocks {
		if theBlock == ignoredBlock {
			continue
		}
		if theBlock.max[0] >= min[0] && theBlock.min[0] <= max[0] &&
			theBlock.max[1] >= min[1] && theBlock.min[1] <= max[1] &&
			theBlock.max[2] >= min[2] && theBlock.min[2] <= max[2] {
			hitBoxes = append(hitBoxes, theBlock)
		}
	}
	if 0 >= min[2] && 0 <= max[2] {
		hitBoxes = append(hitBoxes, ground)
	}
	return
}

func moveZ(zOffset int, coordinates [3]int) [3]int {
	return [3]int{
		coordinates[0],
		coordinates[1],
		coordinates[2] - zOffset,
	}
}

func (currentBlock *block) canBeSafelyDisintegrated() bool {
	for _, supported := range currentBlock.supporting {
		safe := false
		for _, supportOfSupported := range supported.supports {
			if supportOfSupported != currentBlock {
				safe = true
				break
			}
		}
		if !safe {
			return false
		}
	}
	return true
}
