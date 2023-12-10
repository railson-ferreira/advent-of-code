package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type handCategory int8

const (
	fiveOfAKind handCategory = 7 - iota
	fourOfAKind
	fullHouse
	threeOfAKind
	twoPair
	onePair
	highCard
	uncategorized
)

type card struct {
	label    string
	strength int
}

type hand struct {
	category handCategory
	text     string
	cards    [5]card
	bid      int
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

	hands := formatInput(input)

	for index, theHand := range hands {
		hands[index].category = theHand.findCategory()
	}

	orderByCategory(hands)

	orderedByRank := getOrderedByRank(hands)

	sum := 0
	for index, theHand := range orderedByRank {
		rank := index + 1
		//fmt.Println(rank, theHand.category, theHand.text, theHand.bid)
		sum += rank * theHand.bid
	}

	fmt.Println(sum)
}

func formatInput(input string) (hands []hand) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, " ")
		cardsText := parts[0]
		bidText := parts[1]
		bid, _ := strconv.Atoi(bidText)

		theHand := hand{
			category: uncategorized,
			text:     cardsText,
			bid:      bid,
		}

		for index, letter := range cardsText {
			label := string(letter)
			theHand.cards[index] = card{
				label:    label,
				strength: getLabelStrength(label),
			}
		}
		hands = append(hands, theHand)
	}

	return hands
}

func getLabelStrength(label string) int {
	switch label {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 1
	case "T":
		return 10
	case "9":
		return 9
	case "8":
		return 8
	case "7":
		return 7
	case "6":
		return 6
	case "5":
		return 5
	case "4":
		return 4
	case "3":
		return 3
	case "2":
		return 2
	case "1":
		return 1
	}

	panic("Invalid label")
}

func (theHand hand) findCategory() (category handCategory) {
	groupedCards := make(map[string]int)
	hasJoker := false
	for _, theCard := range theHand.cards {
		if theCard.label == "J" {
			hasJoker = true
		}
		groupedCards[theCard.label]++
	}
	distinctLabels := len(groupedCards)
	if hasJoker && distinctLabels > 1 {
		distinctLabels--
	}
	switch distinctLabels {
	case 1:
		return fiveOfAKind
	case 4:
		return onePair
	case 5:
		return highCard
	}
	var repetitions []int
	jokerQuantity := 0
	for label, quantity := range groupedCards {
		if label == "J" {
			jokerQuantity += quantity
			continue
		}
		repetitions = append(repetitions, quantity)
	}
	sort.Ints(repetitions)
	if hasJoker {
		if len(repetitions) == 0 {
			repetitions = append(repetitions, jokerQuantity)
		} else {
			repetitions[len(repetitions)-1] += jokerQuantity
		}
	}
	switch distinctLabels {
	case 2:
		if repetitions[1] == 4 {
			return fourOfAKind
		} else {
			return fullHouse
		}
	case 3:
		if repetitions[2] == 3 {
			return threeOfAKind
		} else {
			return twoPair
		}
	}
	panic("Invalid Category")
}

func orderByCategory(hands []hand) {
	slices.SortFunc(hands, func(a, b hand) int {
		return cmp.Compare(a.category, b.category)
	})
}

func getOrderedByRank(orderedHands []hand) (orderedByRank []hand) {
	groupedHands := make(map[handCategory][]hand)
	for _, theHand := range orderedHands {
		hands := groupedHands[theHand.category]
		hands = append(hands, theHand)
		groupedHands[theHand.category] = hands
	}
	var categories []int
	for category := range groupedHands {
		categories = append(categories, int(category))
	}
	sort.Ints(categories)
	for _, category := range categories {
		hands := groupedHands[handCategory(category)]
		slices.SortFunc(hands, func(a, b hand) int {
			for i := 0; i < 5; i++ {
				if comparison := cmp.Compare(a.cards[i].strength, b.cards[i].strength); comparison != 0 {
					return comparison
				}
			}
			panic("repeated hands")
		})

		orderedByRank = append(orderedByRank, hands...)
	}

	return orderedByRank
}
