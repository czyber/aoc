package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	aocdownloader "github.com/czyber/aoc-downloader"
)

func main() {
	inputString, err := aocdownloader.GetInput("2023", "4")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(inputString, "\n")

	sum := 0
	for _, line := range input {
		fmt.Println(line)
		reCard := regexp.MustCompile(`Card\s+\d+:`)
		reNumbers := regexp.MustCompile(`\d+`)
		trimmedLine := reCard.ReplaceAllString(line, "")
		numbers := strings.Split(trimmedLine, "|")
		winningNumbers := reNumbers.FindAllString(numbers[0], -1)
		myNumbers := reNumbers.FindAllString(numbers[1], -1)
		myNumbersSet := map[string]bool{}
		for _, v := range myNumbers {
			myNumbersSet[v] = true
		}

		winningNumbersSet := map[string]bool{}
		for _, v := range winningNumbers {
			winningNumbersSet[v] = true
		}

		winningNumbersInMyNumbers := map[string]bool{}

		for k, _ := range winningNumbersSet {
			if myNumbersSet[k] {
				winningNumbersInMyNumbers[k] = true
			}
		}

		if len(winningNumbersInMyNumbers) > 0 {
			partialSum := 1
			for i := 0; i < len(winningNumbersInMyNumbers)-1; i++ {
				partialSum *= 2
			}
			sum += partialSum
		}
	}
	fmt.Println("Sum: ", sum)
}
