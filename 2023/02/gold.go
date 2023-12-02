package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type color int8

const (
	_ color = iota
	red
	green
	blue
)

type game struct {
	id   int
	sets []map[color]int
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

	games := formatInput(input)

	var powerSum int
	for _, theGame := range games {
		minRed, minGreen, minBlue := theGame.fewestCubesNumber()
		power := minRed * minGreen * minBlue
		powerSum += power
	}
	fmt.Println(powerSum)
}

func formatInput(input string) []game {
	var games []game

	// performance here is not that important
	idRegex := regexp.MustCompile("Game (\\d+):")
	redRegex := regexp.MustCompile("(\\d+) red")
	greenRegex := regexp.MustCompile("(\\d+) green")
	blueRegex := regexp.MustCompile("(\\d+) blue")
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}

		gameId, _ := strconv.Atoi(idRegex.FindStringSubmatch(line)[1])
		theGame := game{
			id: gameId,
		}

		sets := strings.Split(strings.Split(line, ":")[1], ";")

		for _, set := range sets {
			redMatchResult := redRegex.FindStringSubmatch(set)
			greenMatchResult := greenRegex.FindStringSubmatch(set)
			blueMatchResult := blueRegex.FindStringSubmatch(set)
			colors := map[color]int{
				red:   0,
				green: 0,
				blue:  0,
			}
			if len(redMatchResult) > 1 {
				count, _ := strconv.Atoi(redMatchResult[1])
				colors[red] = count
			}
			if len(greenMatchResult) > 1 {
				count, _ := strconv.Atoi(greenMatchResult[1])
				colors[green] = count
			}
			if len(blueMatchResult) > 1 {
				count, _ := strconv.Atoi(blueMatchResult[1])
				colors[blue] = count
			}
			theGame.sets = append(theGame.sets, colors)
		}
		games = append(games, theGame)
	}

	return games
}

func (theGame game) fewestCubesNumber() (minRed int, minGreen int, minBlue int) {
	for _, set := range theGame.sets {
		if set[red] > minRed {
			minRed = set[red]
		}
		if set[green] > minGreen {
			minGreen = set[green]
		}
		if set[blue] > minBlue {
			minBlue = set[blue]
		}
	}
	return minRed, minGreen, minBlue
}
