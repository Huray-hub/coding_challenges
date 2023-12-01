package day01

import "testing"

func TestDay01(t *testing.T) {
	t.Run("input 0", func(t *testing.T) {
		// Arrange
		inputFile := "input0.txt"
		expected := 281

		// Act
		actual := day01(inputFile)

		// Assert
		if expected != actual {
			t.Fatalf("expected: %d, actual: %d", expected, actual)
		}
	})

	t.Run("input 1", func(t *testing.T) {
		// Arrange
		inputFile := "input1.txt"
		expected := 55093

		// Act
		actual := day01(inputFile)

		// Assert
		if expected != actual {
			t.Fatalf("expected: %d, actual: %d", expected, actual)
		}
	})
}

func BenchmarkDay01(b *testing.B) {
	// Arrange
	inputFile := "input1.txt"
	/* expected := 55093 */

	for i := 0; i < b.N; i++ {
		// Act
		_ = day01(inputFile)
		b.ReportAllocs()
	}
}
