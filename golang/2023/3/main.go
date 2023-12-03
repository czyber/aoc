package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	aocdownloader "github.com/czyber/aoc-downloader"
)

func partsFromLocs(locs [][]int, line string, index int) []Part {
	var parts []Part
	for _, loc := range locs {
		value, err := strconv.Atoi(line[loc[0]:loc[1]])
		if err != nil {
			log.Fatal("Error converting string to integer: ", err)
		}
		var indices [][]int
		for i := loc[0]; i < loc[1]; i++ {
			indices = append(indices, []int{index, i})
		}
		part := Part{
			Value:   value,
			Indices: indices,
		}
		parts = append(parts, part)
	}
	return parts
}

func symbolsFromLocs(locs [][]int, line string, index int) []Symbol {
	var symbols []Symbol
	for _, loc := range locs {
		symbol := Symbol{
			Value:   line[loc[0]:loc[1]],
			Indices: []int{index, loc[0]},
		}
		symbols = append(symbols, symbol)
	}
	return symbols
}

func buildEngineSchematic(input []string) EngineSchematic {
	engineSchematic := EngineSchematic{}
	re := regexp.MustCompile(`\d+`)
	reSymbol := regexp.MustCompile(`[^0-9.]`)
	for index, line := range input {
		locs := re.FindAllStringIndex(line, -1)
		locsSymbols := reSymbol.FindAllStringIndex(line, -1)
		parts := partsFromLocs(locs, line, index)
		symbols := symbolsFromLocs(locsSymbols, line, index)
		engineSchematic.addParts(parts)
		engineSchematic.addSymbols(symbols)

	}
	return engineSchematic
}

func part1() {
	inputString, err := aocdownloader.GetInput("2023", "3")

	if err != nil {
		log.Fatal("Error retrieving input: ", err)
	}

	input := strings.Split(inputString, "\n")

	engineSchematic := buildEngineSchematic(input)
	sum := engineSchematic.buildPartSum()
	fmt.Println(sum)
}

func part2() {
	inputString, err := aocdownloader.GetInput("2023", "3")

	if err != nil {
		log.Fatal("Error retrieving input: ", err)
	}

	input := strings.Split(inputString, "\n")

	engineSchematic := buildEngineSchematic(input)
	sum := engineSchematic.buildGearSum()
	fmt.Println(sum)
}

func main() {
	part1()
	part2()
}
