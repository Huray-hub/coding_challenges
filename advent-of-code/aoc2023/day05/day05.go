package day05

import (
	"log"
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/Huray-hub/coding_challenges/advent-of-code/aoc2023/helpers"
)

type (
	source int
	value  struct {
		destination int
		steps       int
	}
)

// part 2 solution
func Solution2(inputFile string) int {
	text, err := helpers.ReadInputFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	seedsLine := text[0]
	seedsStrLine := strings.Split(seedsLine, "seeds: ")[1]
	seedsStr := strings.Fields(seedsStrLine)
	seeds := make([]int, len(seedsStr))
	for i, seedStr := range seedsStr {
		seeds[i], err = strconv.Atoi(seedStr)
		if err != nil {
			log.Fatal(err)
		}
	}

	maps := make([]map[source]value, 0, 7)
	i := 1
	for i < len(text) {
		if text[i] == "" {
			i += 2
			continue
		}
		var newMap map[source]value
		newMap, i = convertLinesToMap(text, i)
		maps = append(maps, newMap)
	}

	minimum := math.MaxInt32
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 0; i < len(seeds); i += 2 {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			localMin := math.MaxInt32
			length := seeds[i] + seeds[i+1]
			var result int
			for seed := seeds[i]; seed < length; seed++ {
				result = seed

				for _, m := range maps {
					for source, value := range m {
						if int(source) <= result && result <= int(source)+value.steps {
							result = value.destination + (result - int(source))
							break
						}
					}
				}
				if localMin > result {
					localMin = result
				}
				ch <- localMin
			}
		}(i, wg)
	}

	go func(ch chan int, wg *sync.WaitGroup) {
		wg.Wait()
		close(ch)
	}(ch, wg)

	for localMin := range ch {
		if localMin < minimum {
			minimum = localMin
		}
	}

	return minimum
}

// part 1 solution
func solution1(inputFile string) int {
	text, err := helpers.ReadInputFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	seedsLine := strings.Split(text[0], "seeds: ")[1]
	seedsStr := strings.Fields(seedsLine)
	seeds := make([]int, len(seedsStr))
	for i, seedStr := range seedsStr {
		seeds[i], err = strconv.Atoi(seedStr)
		if err != nil {
			log.Fatal(err)
		}
	}

	maps := make([]map[source]value, 0, 7)
	i := 1
	for i < len(text) {
		if text[i] == "" {
			i += 2
			continue
		}
		var newMap map[source]value
		newMap, i = convertLinesToMap(text, i)
		maps = append(maps, newMap)
	}

	minimum := math.MaxInt32
	for _, seed := range seeds {
		result := seed
		for _, m := range maps {
			for source, value := range m {
				if int(source) <= result && result <= int(source)+value.steps {
					result = value.destination + (result - int(source))
					break
				}
			}
		}
		if minimum > result {
			minimum = result
		}
	}

	return minimum
}

func convertLinesToMap(text []string, idx int) (map[source]value, int) {
	m := make(map[source]value)
	for idx < len(text) && text[idx] != "" {
		temp := strings.Fields(text[idx])
		dest, err := strconv.Atoi(temp[0])
		if err != nil {
			log.Fatalf("could not convert destination: %s", err)
		}

		src, err := strconv.Atoi(temp[1])
		if err != nil {
			log.Fatalf("could not convert source: %s", err)
		}

		steps, err := strconv.Atoi(temp[2])
		if err != nil {
			log.Fatalf("could not convert steps: %s", err)
		}

		m[source(src)] = value{
			destination: dest,
			steps:       steps,
		}

		idx++
	}
	return m, idx
}
