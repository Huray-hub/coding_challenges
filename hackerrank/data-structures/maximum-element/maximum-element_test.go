package maximumelement

import (
	"reflect"
	"testing"
)

func TestGetMax(t *testing.T) {
	// Arrange
	type args struct {
		// operationLen int
		operations []string
	}
	tests := []struct {
		name     string
		args     args
		expected []int32
	}{
		{
			name: "Test case 1",
			args: args{
				// operationLen: 10,
				operations: []string{
					"1 97",
					"2",
					"1 20",
					"2",
					"1 26",
					"1 20",
					"2",
					"3",
					"1 91",
					"3",
				},
			},
			expected: []int32{26, 91},
		},
		{
			name: "Test case 2",
			args: args{
				operations: []string{
					"1 16",
					"2",
				},
			},
			expected: []int32{},
		},
		{
			name: "Test case 3",
			args: args{
				operations: []string{
					"1 1",
					"1 44",
					"3",
					"3",
					"2",
					"3",
					"3",
					"1 3",
					"1 37",
					"2",
					"3",
					"1 29",
					"3",
					"1 73",
					"1 51",
					"3",
					"3",
					"3",
					"1 70",
					"3",
					"1 8",
					"2",
					"1 49",
					"1 56",
					"1 81",
					"2",
					"1 59",
					"1 44",
					"2",
					"3",
					"3",
					"2",
					"3",
					"3",
					"1 4",
					"3",
					"1 89",
					"2",
					"1 37",
					"1 50",
					"1 64",
					"2",
					"1 49",
					"1 35",
					"1 85",
					"3",
					"1 41",
					"2",
					"3",
					"3",
					"1 86",
					"3",
					"1 60",
					"1 8",
					"3",
					"1 100",
					"3",
					"1 83",
					"3",
					"1 47",
					"2",
					"1 78",
					"2",
					"1 55",
					"1 97",
					"2",
					"3",
					"1 40",
				},
			},
			expected: []int32{
				44,
				44,
				1,
				1,
				3,
				29,
				73,
				73,
				73,
				73,
				73,
				73,
				73,
				73,
				73,
				85,
				85,
				85,
				86,
				86,
				100,
				100,
				100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			actual := getMax(tt.args.operations)
			// Assert

			if len(tt.expected) != len(actual) {
				t.Fatalf("expected len: %d, actual: %d", len(tt.expected), len(actual))
			}

			if !reflect.DeepEqual(tt.expected, actual) {
				t.Fatal("arrays are not equal")
			}
		})
	}
}
