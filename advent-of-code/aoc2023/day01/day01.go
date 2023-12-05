// Package day01 provides solutions for aoc day 1 "Trebuchet?!" problem https://adventofcode.com/2023/day/1
package day01

import (
	"log"
	"strings"
	"unicode"

	"github.com/Huray-hub/coding_challenges/advent-of-code/aoc2023/helpers"
)

type number int

const (
	one number = iota + 1
	two
	three
	four
	five
	six
	seven
	eight
	nine
)

func (n number) String() string {
	return [...]string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}[n]
}

// part 2 solution
func solution2(inputFile string) int {
	text, err := helpers.ReadInputFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var calibrationSum int
	for _, line := range text {
		calibrationSum += extractCalibration(line)
	}
	return calibrationSum
}

// part 1 solution
func solution1(inputFile string) int {
	text, err := helpers.ReadInputFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var calibrationSum int
	for _, line := range text {
		calibrationSum += extractCalibrationWithoutLetters(line)
	}
	return calibrationSum
}

func extractCalibration(line string) int {
	var first, second rune

	var (
		i    int
		char rune
	)

	for i, char = range line {
		if unicode.IsDigit(char) {
			first = char
			break
		}

		if digitStr, ok := containsDigitLetter(line[i:]); ok {
			first = digitStr
			break
		}
	}

	for j := len(line) - 1; j >= i+1; j-- {
		char := rune(line[j])
		if unicode.IsDigit(char) {
			second = char
			break
		}

		if digitStr, ok := containsDigitLetter(line[j:]); ok {
			second = digitStr
			break
		}
	}

	if second == 0 {
		second = first
	}

	number := (int(first)-'0')*10 + (int(second) - '0')
	return number
}

func extractCalibrationWithoutLetters(line string) int {
	var first, second rune

	for _, char := range line {
		if unicode.IsDigit(char) {
			first = char
			break
		}
	}

	for j := len(line) - 1; j >= 0; j-- {
		char := rune(line[j])
		if unicode.IsDigit(char) {
			second = char
			break
		}
	}

	number := (int(first)-'0')*10 + (int(second) - '0')
	return number
}

func containsDigitLetter(substr string) (rune, bool) {
	switch {
	case strings.HasPrefix(substr, one.String()):
		return '1', true
	case strings.HasPrefix(substr, two.String()):
		return '2', true
	case strings.HasPrefix(substr, three.String()):
		return '3', true
	case strings.HasPrefix(substr, four.String()):
		return '4', true
	case strings.HasPrefix(substr, five.String()):
		return '5', true
	case strings.HasPrefix(substr, six.String()):
		return '6', true
	case strings.HasPrefix(substr, seven.String()):
		return '7', true
	case strings.HasPrefix(substr, eight.String()):
		return '8', true
	case strings.HasPrefix(substr, nine.String()):
		return '9', true
	}
	return '0', false
}
