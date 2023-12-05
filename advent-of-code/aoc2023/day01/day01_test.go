package day01

import "testing"

func TestSolution2(t *testing.T) {
	// Arrange
	type args struct {
		inputFile string
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{
			name: "input 02",
			args: args{
				inputFile: "input02.txt",
			},
			expected: 281,
		},
		{
			name: "input 1",
			args: args{
				inputFile: "input1.txt",
			},
			expected: 55093,
		},
		{
			name: "input alexaki",
			args: args{
				inputFile: "input-alexaki.txt",
			},
			expected: 53221,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			actual := solution2(tt.args.inputFile)
			// Assert
			if tt.expected != actual {
				t.Fatalf("expected: %d, actual: %d", tt.expected, actual)
			}
		})
	}
}

func TestSolution1(t *testing.T) {
	// Arrange
	type args struct {
		inputFile string
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{
			name: "input 01",
			args: args{
				inputFile: "input01.txt",
			},
			expected: 142,
		},
		{
			name: "input 1",
			args: args{
				inputFile: "input1.txt",
			},
			expected: 55002,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			actual := solution1(tt.args.inputFile)
			// Assert
			if tt.expected != actual {
				t.Fatalf("expected: %q, actual: %q", tt.expected, actual)
			}
		})
	}
}
