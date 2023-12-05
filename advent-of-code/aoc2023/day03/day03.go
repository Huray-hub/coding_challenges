// Package day03 provides solutions for aoc day 3 "" problem https://adventofcode.com/2023/day/3
package day03

import (
	"log"
	"strconv"
	"unicode"

	"github.com/Huray-hub/coding_challenges/advent-of-code/aoc2023/helpers"
)

// TODO: refactor solutions

// part 2 solution
func solution2(inputFile string) int {
	text, err := helpers.ReadInputFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var sum int
	for i, line := range text {
		for j := 0; j < len(line); j++ {
			if char := rune(line[j]); char != '*' {
				continue
			}
			gearRatio, isGear := findGearRatio(text, i, j)
			if isGear {
				sum += gearRatio
			}
		}
	}

	return sum
}

func findGearRatio(text []string, lineIdx, charIdx int) (int, bool) {
	var oneNumberFound bool
	var twoNumbersFound bool
	ratio := 1

	var upleft bool
	if lineIdx-1 >= 0 && charIdx-1 >= 0 && unicode.IsDigit(rune(text[lineIdx-1][charIdx-1])) {
		if oneNumberFound {
			twoNumbersFound = true
		} else {
			oneNumberFound = true
		}
		ratio *= extractDigitV2(text, lineIdx-1, charIdx-1)
		upleft = true
	}

	var up bool
	if lineIdx-1 >= 0 && unicode.IsDigit(rune(text[lineIdx-1][charIdx])) {
		if !upleft {
			if oneNumberFound {
				twoNumbersFound = true
			} else {
				oneNumberFound = true
			}
			ratio *= extractDigitV2(text, lineIdx-1, charIdx)
		}
		up = true
	}

	if lineIdx-1 >= 0 && charIdx+1 < len(text[lineIdx]) &&
		unicode.IsDigit(rune(text[lineIdx-1][charIdx+1])) {
		if !up {
			if oneNumberFound {
				twoNumbersFound = true
			} else {
				oneNumberFound = true
			}
			ratio *= extractDigitV2(text, lineIdx-1, charIdx+1)
		}
	}

	// left
	if charIdx-1 >= 0 && unicode.IsDigit(rune(text[lineIdx][charIdx-1])) {
		if oneNumberFound {
			twoNumbersFound = true
		} else {
			oneNumberFound = true
		}
		ratio *= extractDigitV2(text, lineIdx, charIdx-1)
	}

	// right
	if charIdx+1 < len(text[lineIdx]) && unicode.IsDigit(rune(text[lineIdx][charIdx+1])) {
		if oneNumberFound {
			twoNumbersFound = true
		} else {
			oneNumberFound = true
		}
		ratio *= extractDigitV2(text, lineIdx, charIdx+1)
	}

	var downleft bool
	if lineIdx+1 < len(text) && charIdx-1 >= 0 &&
		unicode.IsDigit(rune(text[lineIdx+1][charIdx-1])) {
		if oneNumberFound {
			twoNumbersFound = true
		} else {
			oneNumberFound = true
		}
		ratio *= extractDigitV2(text, lineIdx+1, charIdx-1)
		downleft = true
	}

	var down bool
	if lineIdx+1 < len(text) && unicode.IsDigit(rune(text[lineIdx+1][charIdx])) {
		if !downleft {
			if oneNumberFound {
				twoNumbersFound = true
			} else {
				oneNumberFound = true
			}
			ratio *= extractDigitV2(text, lineIdx+1, charIdx)
		}
		down = true
	}
	if lineIdx+1 < len(text) && charIdx+1 < len(text[lineIdx]) &&
		unicode.IsDigit(rune(text[lineIdx+1][charIdx+1])) {
		if !down {
			if oneNumberFound {
				twoNumbersFound = true
			} else {
				oneNumberFound = true
			}
			ratio *= extractDigitV2(text, lineIdx+1, charIdx+1)
		}
	}

	return ratio, twoNumbersFound
}

func extractDigitV2(text []string, lineIdx, startIdx int) (number int) {
	/* var temp rune */
	line := text[lineIdx]
	charIdx := startIdx
	char := rune(line[charIdx])

	for charIdx-1 >= 0 && unicode.IsDigit(rune(line[charIdx-1])) {
		charIdx--
	}

	var tempNumber string
	for unicode.IsDigit(char) {
		tempNumber += string(line[charIdx])
		charIdx++
		if charIdx == len(line) {
			break
		}
		char = rune(line[charIdx])
	}

	var err error
	number, err = strconv.Atoi(tempNumber)
	if err != nil {
		log.Fatalf("could not parse strign %q to int: %s", tempNumber, err)
	}

	return
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// part 1 solution
func solution1(inputFile string) int {
	text, err := helpers.ReadInputFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var sum int
	for i, line := range text {
		j := 0
		for j < len(line) {
			if char := rune(line[j]); unicode.IsDigit(char) {
				var number int
				var isvalid bool
				if j, number, isvalid = extractDigit(text, i, j); isvalid {
					sum += number
				}
				_ = number
			} else {
				j++
			}
		}
	}

	return sum
}

func extractDigit(text []string, lineIdx, startIdx int) (idx, number int, isValid bool) {
	/* var temp rune */
	line := text[lineIdx]
	var tempNumber string
	charIdx := startIdx
	char := rune(line[charIdx])
	for unicode.IsDigit(char) {
		tempNumber += string(line[charIdx])
		if !isValid {
			isValid = isPart(text, lineIdx, charIdx)
		}
		charIdx++
		if charIdx == len(line) {
			break
		}
		char = rune(line[charIdx])
	}

	var err error
	number, err = strconv.Atoi(tempNumber)
	if err != nil {
		log.Fatalf("could not parse strign %q to int: %s", tempNumber, err)
	}
	idx = charIdx

	return
}

func isPart(text []string, lineIdx, charIdx int) bool {
	if charIdx+1 < len(text[lineIdx]) && isSymbol(rune(text[lineIdx][charIdx+1])) {
		return true
	}
	if charIdx-1 >= 0 && isSymbol(rune(text[lineIdx][charIdx-1])) {
		return true
	}
	if lineIdx-1 >= 0 && isSymbol(rune(text[lineIdx-1][charIdx])) {
		return true
	}
	if lineIdx+1 < len(text) && isSymbol(rune(text[lineIdx+1][charIdx])) {
		return true
	}

	if lineIdx-1 >= 0 && charIdx-1 >= 0 && isSymbol(rune(text[lineIdx-1][charIdx-1])) {
		return true
	}
	if lineIdx-1 >= 0 && charIdx+1 < len(text[lineIdx]) &&
		isSymbol(rune(text[lineIdx-1][charIdx+1])) {
		return true
	}
	if lineIdx+1 < len(text) && charIdx-1 >= 0 && isSymbol(rune(text[lineIdx+1][charIdx-1])) {
		return true
	}
	if lineIdx+1 < len(text) && charIdx+1 < len(text[lineIdx]) &&
		isSymbol(rune(text[lineIdx+1][charIdx+1])) {
		return true
	}
	return false
}

func isSymbol(char rune) bool {
	return char != '.' && !unicode.IsDigit(char)
}
