package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type mapper func(source int) (destination *int)

type almanac struct {
	seedsStarts                  []int
	seedsLengths                 []int
	seedToSoilMappers            []mapper
	soilToFertilizerMappers      []mapper
	fertilizerToWaterMappers     []mapper
	waterToLightMappers          []mapper
	lightToTemperatureMappers    []mapper
	temperatureToHumidityMappers []mapper
	humidityToLocationMappers    []mapper
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

	theAlmanac := formatInput(input)

	minLocation := -1

	fmt.Println("Started: ", time.Now().String())
	for i, startSeed := range theAlmanac.seedsStarts {
		length := theAlmanac.seedsLengths[i]
		for seedNumber := startSeed; seedNumber < startSeed+length; seedNumber++ {
			soil := theAlmanac.findSoil(seedNumber)
			fertilizer := theAlmanac.findFertilizer(soil)
			water := theAlmanac.findWater(fertilizer)
			light := theAlmanac.findLight(water)
			temperature := theAlmanac.findTemperature(light)
			humidity := theAlmanac.findHumidity(temperature)
			location := theAlmanac.findLocation(humidity)
			if minLocation == -1 {
				minLocation = location
			} else if location < minLocation {
				minLocation = location
			}
		}
		fmt.Println("Current date and time is: ", time.Now().String())
		fmt.Printf("%v/%v\n", i+1, len(theAlmanac.seedsStarts))
	}
	fmt.Println("Finished: ", time.Now().String())
	fmt.Println(minLocation)
}

func formatInput(input string) (theAlmanac almanac) {
	lines := strings.Split(input, "\n")

	type step int
	const (
		seedToSoil step = iota
		soilToFertilizer
		fertilizerToWater
		waterToLight
		lightToTemperature
		temperatureToHumidity
		humidityToLocation
	)

	currentStep := seedToSoil

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		if i == 0 {
			seedsTexts := strings.Split(strings.Split(line, ": ")[1], " ")
			for seedPairIndex := 0; seedPairIndex < len(seedsTexts)/2; seedPairIndex++ {
				startSeedText := seedsTexts[seedPairIndex*2]
				lengthText := seedsTexts[seedPairIndex*2+1]
				startSeed, _ := strconv.Atoi(startSeedText)
				length, _ := strconv.Atoi(lengthText)

				theAlmanac.seedsStarts = append(theAlmanac.seedsStarts, startSeed)
				theAlmanac.seedsLengths = append(theAlmanac.seedsLengths, length)
			}
			continue
		}

		if strings.Contains(line, "seed-to-soil") {
			currentStep = seedToSoil
			continue
		}

		if strings.Contains(line, "soil-to-fertilizer") {
			currentStep = soilToFertilizer
			continue
		}

		if strings.Contains(line, "fertilizer-to-water") {
			currentStep = fertilizerToWater
			continue
		}

		if strings.Contains(line, "water-to-light") {
			currentStep = waterToLight
			continue
		}

		if strings.Contains(line, "light-to-temperature") {
			currentStep = lightToTemperature
			continue
		}
		if strings.Contains(line, "temperature-to-humidity") {
			currentStep = temperatureToHumidity
			continue
		}
		if strings.Contains(line, "humidity-to-location") {
			currentStep = humidityToLocation
			continue
		}

		parts := strings.Split(line, " ")
		destinationText := parts[0]
		sourceText := parts[1]
		lengthText := parts[2]
		destination, _ := strconv.Atoi(destinationText)
		source, _ := strconv.Atoi(sourceText)
		length, _ := strconv.Atoi(lengthText)

		theMapper := func(src int) (dst *int) {
			if src >= source && src < source+length {
				dst = new(int)
				*dst = destination + (src - source)
			}

			return dst
		}

		switch currentStep {
		case seedToSoil:
			theAlmanac.seedToSoilMappers = append(theAlmanac.seedToSoilMappers, theMapper)
			break
		case soilToFertilizer:
			theAlmanac.soilToFertilizerMappers = append(theAlmanac.soilToFertilizerMappers, theMapper)
			break
		case fertilizerToWater:
			theAlmanac.fertilizerToWaterMappers = append(theAlmanac.fertilizerToWaterMappers, theMapper)
			break
		case waterToLight:
			theAlmanac.waterToLightMappers = append(theAlmanac.waterToLightMappers, theMapper)
			break
		case lightToTemperature:
			theAlmanac.lightToTemperatureMappers = append(theAlmanac.lightToTemperatureMappers, theMapper)
			break
		case temperatureToHumidity:
			theAlmanac.temperatureToHumidityMappers = append(theAlmanac.temperatureToHumidityMappers, theMapper)
			break
		case humidityToLocation:
			theAlmanac.humidityToLocationMappers = append(theAlmanac.humidityToLocationMappers, theMapper)
			break
		}

	}

	return theAlmanac
}

func (theAlmanac almanac) findSoil(seed int) int {
	for _, theMapper := range theAlmanac.seedToSoilMappers {
		if valuePointer := theMapper(seed); valuePointer != nil {
			return *valuePointer
		}
	}
	return seed
}

func (theAlmanac almanac) findFertilizer(soil int) int {
	for _, theMapper := range theAlmanac.soilToFertilizerMappers {
		if valuePointer := theMapper(soil); valuePointer != nil {
			return *valuePointer
		}
	}
	return soil
}

func (theAlmanac almanac) findWater(fertilizer int) int {
	for _, theMapper := range theAlmanac.fertilizerToWaterMappers {
		if valuePointer := theMapper(fertilizer); valuePointer != nil {
			return *valuePointer
		}
	}
	return fertilizer
}

func (theAlmanac almanac) findLight(water int) int {
	for _, theMapper := range theAlmanac.waterToLightMappers {
		if valuePointer := theMapper(water); valuePointer != nil {
			return *valuePointer
		}
	}
	return water
}

func (theAlmanac almanac) findTemperature(light int) int {
	for _, theMapper := range theAlmanac.lightToTemperatureMappers {
		if valuePointer := theMapper(light); valuePointer != nil {
			return *valuePointer
		}
	}
	return light
}

func (theAlmanac almanac) findHumidity(temperature int) int {
	for _, theMapper := range theAlmanac.temperatureToHumidityMappers {
		if valuePointer := theMapper(temperature); valuePointer != nil {
			return *valuePointer
		}
	}
	return temperature
}

func (theAlmanac almanac) findLocation(humidity int) int {
	for _, theMapper := range theAlmanac.humidityToLocationMappers {
		if valuePointer := theMapper(humidity); valuePointer != nil {
			return *valuePointer
		}
	}
	return humidity
}
