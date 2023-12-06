package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type card struct {
	id           int
	leftNumbers  []int
	rightNumbers []int
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

	cards := formatInput(input)

	additionalCards := 0
	for i, card := range cards {
		additionalCards += card.processWin(i, cards)
	}
	cardsLength := len(cards)
	fmt.Printf("cards = %v; additional = %v; total = %v", cardsLength, additionalCards, cardsLength+additionalCards)
}

func formatInput(input string) (cards []card) {
	lines := strings.Split(input, "\n")

	cardRegex := regexp.MustCompile("Card +(\\d+): +(.+) +[|] +(.+)")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		result := cardRegex.FindStringSubmatch(line)
		id := result[1]
		left := result[2]
		right := result[3]

		var leftNumbers []int
		var rightNumbers []int
		for _, value := range strings.Split(left, " ") {
			if len(value) == 0 {
				continue
			}
			number, _ := strconv.Atoi(value)
			leftNumbers = append(leftNumbers, number)
		}
		for _, value := range strings.Split(right, " ") {
			if len(value) == 0 {
				continue
			}
			number, _ := strconv.Atoi(value)
			rightNumbers = append(rightNumbers, number)
		}

		idNumber, _ := strconv.Atoi(id)
		cards = append(cards, card{
			id:           idNumber,
			leftNumbers:  leftNumbers,
			rightNumbers: rightNumbers,
		})

	}

	return cards
}
func (theCard card) processWin(index int, cards []card) (additionalCards int) {
	matchQuantity := theCard.getMatchQuantity()

	nextCards := cards[index+1 : index+1+matchQuantity]
	additionalCards = len(nextCards)
	for i, otherCard := range nextCards {
		additionalCards += otherCard.processWin(index+1+i, cards)
	}
	return additionalCards
}

func (theCard card) getMatchQuantity() (matchQuantity int) {
	for _, number := range theCard.rightNumbers {
		if slices.Contains(theCard.leftNumbers, number) {
			matchQuantity++
		}
	}
	return matchQuantity
}
