package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	aocdownloader "github.com/czyber/aoc-downloader"
)

func readInput() []string {
	input, err := aocdownloader.GetInput("2023", "2")
	if err != nil {
		log.Fatal("Error getting input: ", err)
	}

	inputStrings := strings.Split(input, "\n")

	return inputStrings
}

func extractGameID(line string) string {
	re := regexp.MustCompile(`\d+`)
	game := strings.Split(line, ": ")
	if len(game) == 2 {
		gameID := re.FindString(game[0])
		return gameID
	}
	return ""
}

func extractGameData(line string) string {
	game := strings.Split(line, ": ")
	if len(game) == 2 {
		return game[1]
	}
	return ""
}

func extractSamples(line string) []string {
	return strings.Split(line, "; ")
}

func compareToMaxSize(size int, color string) bool {
	maxSizeForColor := map[string]int{
		"green": 13,
		"red":   12,
		"blue":  14,
	}

	return size <= maxSizeForColor[color]
}

func validateSamples(samples []string) bool {
	re := regexp.MustCompile(`\b\d+\s(green|blue|red)\b`)
	reDigits := regexp.MustCompile(`\d+`)
	reColors := regexp.MustCompile(`\b(green|blue|red)\b`)
	gameIsValid := true
	for _, sample := range samples {
		compiledSamples := re.FindAllString(sample, -1)
		for _, compiledSample := range compiledSamples {
			sampleSize := reDigits.FindString(compiledSample)
			sampleColor := reColors.FindString(compiledSample)
			iSampleSize, err := strconv.Atoi(sampleSize)

			if err != nil {
				log.Fatal("String to integer conversion failed: ", err)
			}

			sampleIsValid := compareToMaxSize(iSampleSize, sampleColor)
			if !sampleIsValid {
				gameIsValid = false
				break
			}

		}
	}
	return gameIsValid
}

func part1() {
	input := readInput()
	sum := 0
	for _, line := range input {
		gameID := extractGameID(line)
		gameData := extractGameData(line)
		samples := extractSamples(gameData)
		gameIsValid := validateSamples(samples)
		if gameIsValid {
			iGameID, err := strconv.Atoi(gameID)
			if err != nil {
				log.Fatal("String to integer conversion failed: ", err)
			}
			sum += iGameID
		}
	}
	fmt.Println("Result:", sum)
}

func getPowerOfSample(samples []string) int {
	maxCountOfColor := map[string]int{
		"green": 0,
		"red":   0,
		"blue":  0,
	}
	re := regexp.MustCompile(`\b\d+\s(green|blue|red)\b`)
	reDigits := regexp.MustCompile(`\d+`)
	reColors := regexp.MustCompile(`\b(green|blue|red)\b`)
	for _, sample := range samples {
		compiledSamples := re.FindAllString(sample, -1)
		for _, compiledSample := range compiledSamples {
			sampleSize := reDigits.FindString(compiledSample)
			sampleColor := reColors.FindString(compiledSample)

			iSampleSize, err := strconv.Atoi(sampleSize)
			if err != nil {
				log.Fatal("String to integer conversion failed: ", err)
			}

			if iSampleSize > maxCountOfColor[sampleColor] {
				maxCountOfColor[sampleColor] = iSampleSize
			}
		}
	}
	product := 1
	for _, value := range maxCountOfColor {
		product *= value
	}
	return product
}

func part2() {
	input := readInput()
	sum := 0
	for _, line := range input {
		gameData := extractGameData(line)
		samples := extractSamples(gameData)
		powerOfSample := getPowerOfSample(samples)
		sum += powerOfSample

	}
	fmt.Println("Result:", sum)

}
func main() {
	part1()
	part2()
}
