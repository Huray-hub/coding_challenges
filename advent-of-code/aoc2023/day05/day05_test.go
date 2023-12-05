package day05

import "testing"

func TestSolution2(t *testing.T) {
	t.Run("input 0", func(t *testing.T) {
		// Arrange
		inputFile := "input0.txt"
		expected := 46

		// Act
		actual := Solution2(inputFile)

		// Actual
		if expected != actual {
			t.Fatalf("expected: %d, actual %d", expected, actual)
		}
	})

	t.Run("input 1", func(t *testing.T) {
		// Arrange
		inputFile := "input1.txt"
		expected := 57451709

		// Act
		actual := Solution2(inputFile)

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
		expected := 35

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
		expected := 551761867

		// Act
		actual := solution1(inputFile)

		// Actual
		if expected != actual {
			t.Fatalf("expected: %d, actual %d", expected, actual)
		}
	})
}
