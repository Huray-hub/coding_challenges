package day02

import "testing"

func TestSolution1(t *testing.T) {
	t.Run("input 0", func(t *testing.T) {
		// Arrange
		inputFile := "input0.txt"
		expected := 8

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
		expected := 2727

		// Act
		actual := solution1(inputFile)

		// Actual
		if expected != actual {
			t.Fatalf("expected: %d, actual %d", expected, actual)
		}
	})
}

func TestSolution2(t *testing.T) {
	t.Run("input 0", func(t *testing.T) {
		// Arrange
		inputFile := "input0.txt"
		expected := 2286

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
		expected := 56580

		// Act
		actual := solution2(inputFile)

		// Actual
		if expected != actual {
			t.Fatalf("expected: %d, actual %d", expected, actual)
		}
	})
}
