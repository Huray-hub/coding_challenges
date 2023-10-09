// Package inordertraversal contains solution to https://hackerrank.com/challenges/tree-inorder-traversal
package inordertraversal

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

func inOrder(root *node, w io.Writer) {
	if root == nil {
		return
	}
	inOrder(root.left, w)
	fmt.Fprint(w, root.data, " ")
	inOrder(root.right, w)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	scanner.Scan()
	firstLine := scanner.Text()
	nodeLen, err := strconv.Atoi(firstLine)
	checkError(err)

	scanner.Scan()
	secondLine := scanner.Text()

	fields := strings.Fields(secondLine)
	if len(fields) != nodeLen {
		panic("fields and number of nodes differ!")
	}

	var (
		root *node
		data int
	)
	for _, field := range fields {
		data, err = strconv.Atoi(field)
		checkError(err)
		root = insert(root, data)
	}

	inOrder(root, stdout)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
