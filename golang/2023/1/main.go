package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	aocdownloader "github.com/czyber/aoc-downloader"
)

func main() {
	input, err := aocdownloader.GetInput("2023", "1")
	if err != nil {
		log.Fatal("Error getting input: ", err)
	}

	inputStrings := strings.Split(input, "\n")

	re := regexp.MustCompile(`\d`)

	sum := 0
	for _, line := range inputStrings {
		numbers := re.FindAllString(line, -1)
		if len(numbers) > 0 {
			firstNumber := numbers[0]
			secondNumber := numbers[len(numbers)-1]
			integer, err := strconv.ParseInt(firstNumber+secondNumber, 10, 64)
			if err != nil {
				log.Fatal("Error parsing integer: ", err)
			}
			sum += int(integer)
		}
	}
	fmt.Println("Result: ", sum)

}
