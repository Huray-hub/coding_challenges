// Package day02 provides solutions for aoc day 2 "Cube Conundrum" problem https://adventofcode.com/2023/day/2
package day02

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Huray-hub/coding_challenges/advent-of-code/aoc2023/helpers"
)

const (
	blue  = "blue"
	red   = "red"
	green = "green"
)

type cube struct {
	color    string
	quantity int
}

// part 1 solution
func solution1(inputFile string) int {
	text, err := helpers.ReadInputFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var idSum int
	for i, line := range text {
		sets := strings.Split(line, "; ")
		sets[0] = strings.TrimPrefix(sets[0], fmt.Sprintf("Game %d: ", i+1))

		valid, err := isValid(sets)
		if err != nil {
			log.Fatal(err)
		}
		if valid {
			idSum += i + 1
		}
	}

	return idSum
}

// part 2 solution
func solution2(inputFile string) int {
	text, err := helpers.ReadInputFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var powerSum int
	for i, line := range text {
		sets := strings.Split(line, "; ")
		sets[0] = strings.TrimPrefix(sets[0], fmt.Sprintf("Game %d: ", i+1))

		power, err := minimumPower(sets)
		if err != nil {
			log.Fatal(err)
		}

		powerSum += power
	}

	return powerSum
}

func isValid(sets []string) (bool, error) {
	const maxReds, maxGreens, maxBlues = 12, 13, 14

	for _, set := range sets {
		pairs := strings.Split(set, ", ")
		for _, pair := range pairs {
			var (
				c           cube
				quantityStr string
				err         error
			)

			quantityStr, c.color, _ = strings.Cut(pair, " ")
			c.quantity, err = strconv.Atoi(quantityStr)
			if err != nil {
				return false, err
			}

			switch {
			case c.color == blue && c.quantity > maxBlues:
				fallthrough
			case c.color == red && c.quantity > maxReds:
				fallthrough
			case c.color == green && c.quantity > maxGreens:
				return false, nil
			}
		}
	}

	return true, nil
}

func minimumPower(sets []string) (int, error) {
	minBlue, minGreens, minReds := 1, 1, 1 // minimum required for game in order to be valid
	for _, set := range sets {
		pairs := strings.Split(set, ", ")

		for _, pair := range pairs {
			var (
				c           cube
				quantityStr string
				err         error
			)

			quantityStr, c.color, _ = strings.Cut(pair, " ")
			c.quantity, err = strconv.Atoi(quantityStr)
			if err != nil {
				return 0, err
			}

			switch {
			case c.color == blue && c.quantity > minBlue:
				minBlue = c.quantity
			case c.color == red && c.quantity > minReds:
				minReds = c.quantity
			case c.color == green && c.quantity > minGreens:
				minGreens = c.quantity
			}
		}
	}

	return minBlue * minGreens * minReds, nil
}
