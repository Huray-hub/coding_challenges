package treepostordertraversal

import (
	"bytes"
	"testing"
)

func TestTreePostorderTraversal(t *testing.T) {
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
			expected: "4 3 6 5 2 1 ",
		},
		{
			name: "Input2",
			args: args{
				nodeLen:   15,
				nodeElems: []int{1, 14, 3, 7, 4, 5, 15, 6, 13, 10, 11, 2, 12, 8, 9},
			},
			expected: "2 6 5 4 9 8 12 11 10 13 7 3 15 14 1 ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := newTree(tt.args.nodeLen, tt.args.nodeElems)
			// Act
			buffer := &bytes.Buffer{}
			postOrder(root, buffer)
			// Assert
			actual := buffer.String()
			if tt.expected != actual {
				t.Fatalf("expected: %q, actual: %q", tt.expected, actual)
			}
		})
	}
}
