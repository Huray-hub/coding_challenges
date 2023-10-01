package treeinordertraversal

import (
	"bytes"
	"testing"
)

func TestTreeInorderTraversal(t *testing.T) {
	// Arrange
	type args struct {
		nodeLen   int
		nodeElems []int
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "Input1",
			args: args{
				nodeLen:   6,
				nodeElems: []int{1, 2, 5, 3, 6, 4},
			},
			expected: "1 2 3 4 5 6 ",
		},
		{
			name: "Input2",
			args: args{
				nodeLen:   15,
				nodeElems: []int{1, 14, 3, 7, 4, 5, 15, 6, 13, 10, 11, 2, 12, 8, 9},
			},
			expected: "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := newTree(tt.args.nodeLen, tt.args.nodeElems)
			// Act
			buffer := &bytes.Buffer{}
			inOrder(root, buffer)
			// Assert
			actual := buffer.String()
			if tt.expected != actual {
				t.Fatalf("expected: %q, actual: %q", tt.expected, actual)
			}
		})
	}
}
