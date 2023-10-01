// Package treepostordertraversal contains solution to https://hackerrank.com/challenges/tree-postorder-traversal
package treepostordertraversal

import (
	"fmt"
	"io"
)

type node struct {
	data  int
	left  *node
	right *node
}

func newTree(len int, elements []int) (root *node) {
	for i := 0; i < len; i++ {
		root = insert(root, elements[i])
	}
	return root
}

func insert(root *node, data int) *node {
	if root == nil {
		root = &node{data: data}
		return root
	}

	var current *node
	if data <= root.data {
		current = insert(root.left, data)
		root.left = current
	} else {
		current = insert(root.right, data)
		root.right = current
	}
	return root
}

func postOrder(root *node, w io.Writer) {
	if root == nil {
		return
	}
	postOrder(root.left, w)
	postOrder(root.right, w)
	fmt.Fprint(w, root.data, " ")
}

// main() for hackerrank editor (reads input from stdin)
// func solution() {
// 	var (
// 		nodeLen int
// 		data    int
// 		root    *node
// 	)
// 	fmt.Scanf("%d", &nodeLen)
// 	for i := 0; i < nodeLen; i++ {
// 		fmt.Scanf("%d", &data)
// 		root = insert(root, data)
// 	}
// 	postOrder(root, os.Stdout)
// }
