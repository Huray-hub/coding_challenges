package day04

import (
	"log"
	"slices"
	"strings"

	"github.com/Huray-hub/coding_challenges/advent-of-code/aoc2023/helpers"
)

// part 2 solution
func solution2(inputFile string) int {
	text, err := helpers.ReadInputFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	instances := make([]int, len(text))
	for i := range instances {
		instances[i] = 1 // original cards
	}

	var scratchcards int
	for i, line := range text {
		scratchcards += instances[i]

		lineNumbersStr := strings.Split(line, ": ")[1]
		winningNumbersStr, numbersYouHaveStr, _ := strings.Cut(lineNumbersStr, " | ")
		winningNumbers := strings.Fields(winningNumbersStr)
		numbersYouHave := strings.Fields(numbersYouHaveStr)

		var matchingNumbers int
		for _, number := range numbersYouHave {
			if slices.Contains(winningNumbers, number) {
				matchingNumbers++
			}
		}

		for j := i + 1; j < i+1+matchingNumbers && j < len(text); j++ {
			instances[j] += instances[i]
		}
	}

	return scratchcards
}

// part 1 solution
func solution1(inputFile string) int {
	text, err := helpers.ReadInputFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var totalPoints int
	for _, line := range text {
		lineNumbersStr := strings.Split(line, ": ")[1]
		winningNumbersStr, numbersYouHaveStr, _ := strings.Cut(lineNumbersStr, " | ")
		winningNumbers := strings.Fields(winningNumbersStr)
		numberYouHave := strings.Fields(numbersYouHaveStr)

		var cardSum int
		for _, number := range numberYouHave {
			if !slices.Contains(winningNumbers, number) {
				continue
			}

			if cardSum == 0 {
				cardSum++
				continue
			}
			cardSum *= 2
		}
		totalPoints += cardSum
	}

	return totalPoints
}
