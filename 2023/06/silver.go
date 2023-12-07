package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type record struct {
	duration int
	distance int
}

type wayToWin record

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	bytes, err := os.ReadFile("input.txt")
	check(err)
	input := string(bytes)

	records := formatInput(input)

	totalWaysToWinProduct := 1
	for _, theRecord := range records {
		totalWaysToWinProduct *= len(theRecord.getWaysToWin())
	}

	fmt.Println(totalWaysToWinProduct)
}

func formatInput(input string) (records []record) {
	lines := strings.Split(input, "\n")

	spaceRegex := regexp.MustCompile(" +")
	durations := spaceRegex.Split(strings.Split(lines[0], ":")[1], -1)
	distances := spaceRegex.Split(strings.Split(lines[1], ":")[1], -1)
	for index, durationText := range durations {
		if len(durationText) == 0 {
			continue
		}
		duration, _ := strconv.Atoi(durationText)
		distance, _ := strconv.Atoi(distances[index])
		records = append(records, record{
			duration: duration,
			distance: distance,
		})
	}

	return records
}

func (theRecord record) getWaysToWin() (waysToWin []wayToWin) {
	halfDuration := theRecord.duration / 2

	if distance := calculateDistance(theRecord.duration, halfDuration); distance > theRecord.distance {
		waysToWin = append(waysToWin, wayToWin{
			duration: halfDuration,
			distance: distance,
		})
	}

	hasPrevious := true
	hasNext := true
	count := 1
	for hasPrevious && hasNext {
		previousDuration := halfDuration - count
		nextDuration := halfDuration + count
		if previousDuration >= 0 {
			distance := calculateDistance(theRecord.duration, previousDuration)
			if distance > theRecord.distance {
				waysToWin = append(waysToWin, wayToWin{
					duration: previousDuration,
					distance: distance,
				})
			} else {
				// could ignore the most previous at this point
			}
		} else {
			hasPrevious = false
		}
		if nextDuration < theRecord.duration {

			distance := calculateDistance(theRecord.duration, nextDuration)
			if distance > theRecord.distance {
				waysToWin = append(waysToWin, wayToWin{
					duration: nextDuration,
					distance: distance,
				})
			} else {
				// could ignore the most next at this point
			}
		} else {
			hasNext = false
		}

		count++
	}
	return waysToWin
}

func calculateDistance(maxDuration int, pressDuration int) (distance int) {
	timeLeft := maxDuration - pressDuration
	distance = timeLeft * pressDuration
	return distance
}
