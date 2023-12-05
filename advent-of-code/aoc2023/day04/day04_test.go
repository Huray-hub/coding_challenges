package day04

import "testing"

func TestSolution2(t *testing.T) {
	t.Run("input 0", func(t *testing.T) {
		// Arrange
		inputFile := "input0.txt"
		expected := 30

		// Act
		actual := solution2(inputFile)

		// Actual
		if expected != actual {
			t.Fatalf("expected: %d, actual %d", expected, actual)
		}
	})

	t.Run("input 1", func(t *testing.T) {
		// Arrange
		inputFile := "input1.txt"
		expected := 9236992

		// Act
		actual := solution2(inputFile)

		// Actual
		if expected != actual {
			t.Fatalf("expected: %d, actual %d", expected, actual)
		}
	})
}

func TestSolution1(t *testing.T) {
	t.Run("input 0", func(t *testing.T) {
		// Arrange
		inputFile := "input0.txt"
		expected := 13

		// Act
		actual := solution1(inputFile)

		// Actual
		if expected != actual {
			t.Fatalf("expected: %d, actual %d", expected, actual)
		}
	})

	t.Run("input 1", func(t *testing.T) {
		// Arrange
		inputFile := "input1.txt"
		expected := 23028

		// Act
		actual := solution1(inputFile)

		// Actual
		if expected != actual {
			t.Fatalf("expected: %d, actual %d", expected, actual)
		}
	})
}
